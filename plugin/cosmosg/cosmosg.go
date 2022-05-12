package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"

	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	acc "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/group"

	pb "github.com/cosmos/keystone/keystone"
	krplugin "github.com/cosmos/keystone/plugin"
)

var nc *nodeclient = nil

type nodeclient struct {
	ChainId       string `json:"chainId"`
	ServerAddress string `json:"serverAddress"`
	KeyringDir    string `json:"keyringDir"`
	RpcUri        string `json:"rpcUri"`
}

// private methods

func loadConfiguration(pathtoconfig string) (*nodeclient, error) {
	var cf *nodeclient
	configFile, err := os.Open(pathtoconfig)
	defer configFile.Close()

	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&cf)

	log.Printf("Server Address loaded: %s", cf.ServerAddress)
	log.Printf("Chain loaded: %s", cf.ChainId)
	log.Printf("Keyring dir loaded: %s", cf.KeyringDir)

	return cf, nil
}

// how to retrieve node context beyond this one transaction?
func getLocalContext(nc *nodeclient) (*client.Context, error) {

	addr, err := sdk.AccAddressFromBech32(nc.ServerAddress)
	encodingConfig := simapp.MakeTestEncodingConfig()

	if err != nil {
		return nil, err
	}

	rpcclient, err := client.NewClientFromNode(nc.RpcUri)

	if err != nil {
		log.Printf("could not get rpcclient: %v", err)
		return nil, err
	}

	//@@TODO configure keyring.BackendTest using the server-global context, not hardcode
	kb, err := keyring.New("keystone", keyring.BackendTest, nc.KeyringDir, nil, encodingConfig.Codec)

	if err != nil {
		log.Printf("error opening keyring: %v", err)
		return nil, err
	} else {
		log.Printf("keyring: %v", kb)
	}

	l, err := kb.List()

	if err != nil || len(l) < 1 {
		log.Println("error retrieving keys")
		return nil, err
	}

	log.Printf("Keys?: %v\n", l)

	c := client.Context{FromAddress: addr, ChainID: nc.ChainId}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithBroadcastMode(flags.BroadcastSync).
		WithNodeURI(nc.RpcUri).
		WithAccountRetriever(acc.AccountRetriever{}).
		WithClient(rpcclient).
		WithKeyringDir(nc.KeyringDir).
		WithKeyring(kb)

	return &c, nil
}

func createAdminGroup(creatorAddress []byte, memberList []group.Member, metadata string, localContext *client.Context) ([]byte, error) {
	return createGroup(creatorAddress, memberList, metadata, localContext)
}

// CreateGroup creates a Cosmos Group using the MsgCreateGroup, filling the message with the input fields
func createGroup(creatorAddress []byte, memberList []group.Member, metadata string, localContext *client.Context) ([]byte, error) {

	//encCfg := simapp.MakeTestEncodingConfig()
	txBuilder := localContext.TxConfig.NewTxBuilder()

	// @@todo, how to get the private key from the keyring
	// associated with this address?
	adminAddr, err := sdk.AccAddressFromBech32(string(creatorAddress))

	if err != nil {
		log.Printf("Error converting address string: %v", err)
		return nil, err
	}

	// err = localContext.AccountRetriever.EnsureExists(*localContext, adminAddr)

	// if err != nil {
	//  	log.Println("Account does not exist because:", err)
	//  	return nil, err
	// }

	// num, seq, err := localContext.AccountRetriever.GetAccountNumberSequence(*localContext, adminAddr)

	// if err != nil {
	// 	fmt.Printf("Error retrieving account number/sequence: %v", err)
	// 	return nil, err
	// } else {
	// 	fmt.Printf("Account retrieved: %v with seq: %v", num, seq)
	// }

	txBuilder.SetMsgs(&group.MsgCreateGroup{
		Admin:   adminAddr.String(),
		Members: memberList,
		//Metadata: "",
	})

	// @@TODO: abstract the stake fee - will depend on chain
	// config for example
	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("stake", 20)})
	txBuilder.SetGasLimit(50000)

	txFactory := clienttx.Factory{}
	txFactory = txFactory.
		WithChainID(localContext.ChainID).
		WithKeybase(localContext.Keyring).
		WithTxConfig(localContext.TxConfig)

	// Only needed for "offline" accounts?
	//.WithAccountNumber(num).WithSequence(seq)

	info, err := txFactory.Keybase().Key("my_validator")

	if err != nil {
		return nil, err
	}

	fmt.Printf("%v", info)

	// NOT NEEDED IF USING SignTx from the x/auth/client, which
	// does all these things
	// signerData := xauthsigning.SignerData{
	// 	ChainID:       MY_CHAIN,
	// 	AccountNumber: num,
	// 	Sequence:      seq,
	// }

	// signBytes, err := localContext.TxConfig.SignModeHandler().GetSignBytes(signing.SignMode_SIGN_MODE_DIRECT, signerData, txBuilder.GetTx())

	// if err != nil{
	// 	fmt.Println("Error getting signed bytes: ", err)
	// 	return nil, err
	// }

	txJSON, err := localContext.TxConfig.TxJSONEncoder()(txBuilder.GetTx())

	if err != nil {
		fmt.Println("Error getting JSON: ", err)
		return nil, err
	}

	fmt.Printf("Unsigned TX %s\n", txJSON)

	err = authclient.SignTx(txFactory, *localContext, "validator", txBuilder, false, true)

	if err != nil {
		fmt.Println("Error signing: ", err)
		return nil, err
	}

	txBytes, err := localContext.TxConfig.TxEncoder()(txBuilder.GetTx())

	if err != nil {
		fmt.Println("Error encoding transaction: ", err)
		return nil, err
	}

	txJSON, err = localContext.TxConfig.TxJSONEncoder()(txBuilder.GetTx())

	if err != nil {
		fmt.Println("Error getting JSON: ", err)
		return nil, err
	}

	fmt.Printf("Signed TX %s\n", txJSON)

	//res, err := localContext.BroadcastTx(txBytes)

	// @@TODO: use secure connection?
	opts := grpc.WithInsecure()

	// @@TODO: configure the dial location from server context
	grpcConn, err := grpc.Dial("127.0.0.1:9090", opts)

	if err != nil {
		fmt.Println("Err doing grpc dial: ", err)
		return nil, err
	}

	defer grpcConn.Close()

	// @@TODO: configure broadcast mode from server-global context?

	txClient := tx.NewServiceClient(grpcConn)

	res, err := txClient.BroadcastTx(
		context.Background(),
		&tx.BroadcastTxRequest{
			Mode:    tx.BroadcastMode_BROADCAST_MODE_SYNC,
			TxBytes: txBytes,
		},
	)

	if err != nil {
		return nil, err
	}

	fmt.Println(res.TxResponse.Code) // Should be `0` if the tx is successful

	if err != nil {
		fmt.Println("Error broadcasting ", err)
	}

	fmt.Printf("Result: %v", res)

	return []byte{}, nil
}

//adminMembers returns a []group.Member with two members
func adminMembers(addr1 string, addr2 string) []group.Member {

	member1 := group.Member{
		Address: addr1,
		Weight:  strconv.Itoa(1),
	}

	member2 := group.Member{
		Address: addr2,
		Weight:  strconv.Itoa(1),
	}

	return []group.Member{member1, member2}
}

func TypeIdentifier() string {
	return krplugin.Plugin_Type_CosmosG_Id
}

func Init(configPath string) (krplugin.Plugin, error) {

	nc, err := loadConfiguration(configPath)

	if err != nil {
		return nil, err
	} else {
		return nc, nil
	}
}

func (nc *nodeclient) NewKey(in *pb.KeySpec) (*pb.KeyRef, error) {

	member1 := group.Member{
		Address: "cosmos16puqhqxr0zj3r374t64qmj2g8ke2sgrwjwq9s6",
		Weight:  strconv.Itoa(1),
	}

	members := []group.Member{member1}

	log.Printf("Server address: %s", nc.ServerAddress)

	bcContext, err := getLocalContext(nc)

	if err != nil {
		log.Printf("generate local blockchain context failed %v", err)
		return nil, err
	}

	log.Printf("now trying to create the group")

	groupAddr, err := createGroup([]byte("cosmos16puqhqxr0zj3r374t64qmj2g8ke2sgrwjwq9s6"), members, "", bcContext)

	if err != nil {
		log.Printf("group add failed: %v", err)
		return nil, err
	}

	strAddress := string(groupAddr)

	ref := pb.KeyRef{
		Label: strAddress,
	}

	return &ref, nil
}

func (nc *nodeclient) PubKey(in *pb.KeySpec) (*pb.PublicKey, error) {
	return nil, nil
}

func (nc *nodeclient) Sign(in *pb.Msg) (*pb.Signed, error) {
	return nil, nil
}
