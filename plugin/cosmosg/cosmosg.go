package main

import (
	"fmt"
	"os"
	"context"
	"strconv"
	"encoding/json"

	"google.golang.org/grpc"
	
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/group"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	clienttx "github.com/cosmos/cosmos-sdk/client/tx"
	authclient "github.com/cosmos/cosmos-sdk/x/auth/client"
	"github.com/cosmos/cosmos-sdk/types/tx"
	"github.com/cosmos/cosmos-sdk/simapp"
	acc "github.com/cosmos/cosmos-sdk/x/auth/types"
	
	krplugin "github.com/regen-network/keystone2/plugin"
	pb "github.com/regen-network/keystone2/keystone"
)

var nc *nodeclient = nil

type nodeclient struct {
	chainId         string `json:"chainId"`
	serverAddress   string `json:"serverAddress"`
	keyringDir      string `json:"keyringDir"`	
	rpcUri          string `json:"rpcUri"`
}

// private methods

func loadConfiguration(pathtoconfig string) (*nodeclient, error) {
	var config *nodeclient
	configFile, err := os.Open(pathtoconfig)
	defer configFile.Close()
	
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}
	
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, nil
}

// how to retrieve node context beyond this one transaction?
func getLocalContext(nc *nodeclient) (*client.Context, error) {

	addr, err := sdk.AccAddressFromBech32(nc.serverAddress)
	encodingConfig := simapp.MakeTestEncodingConfig()

	if err != nil {
		return nil, err
	}

	rpcclient, err := client.NewClientFromNode(nc.rpcUri)

	if err != nil {
		return nil, err
	}

	//@@TODO configure keyring.BackendTest using the server-global context, not hardcode
	k, err := keyring.New(sdk.KeyringServiceName(), keyring.BackendTest, nc.keyringDir, nil, encodingConfig.Codec)

	// l, err := k.List()

	// fmt.Printf("%v", l)

	if err != nil {
		fmt.Printf("error opening keyring: ", err)
		return nil, err
	}

	c := client.Context{FromAddress: addr, ChainID: nc.chainId}.
		WithCodec(encodingConfig.Codec).
		WithInterfaceRegistry(encodingConfig.InterfaceRegistry).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithBroadcastMode(flags.BroadcastSync).
		WithNodeURI(nc.rpcUri).
		WithAccountRetriever(acc.AccountRetriever{}).
		WithClient(rpcclient).
		WithKeyringDir(nc.keyringDir).
		WithKeyring(k)

	return &c, nil
}

func createAdminGroup(creatorAddress []byte, memberList []group.Member, metadata string, localContext *client.Context) ([]byte, error) {
	return createGroup( creatorAddress, memberList, metadata, localContext )
}

// CreateGroup creates a Cosmos Group using the MsgCreateGroup, filling the message with the input fields
func createGroup(creatorAddress []byte, memberList []group.Member, metadata string, localContext *client.Context) ([]byte, error) {

	encCfg := simapp.MakeTestEncodingConfig()
	txBuilder := encCfg.TxConfig.NewTxBuilder()

	// @@todo, how to get the private key from the keyring
	// associated with this address?
	adminAddr, err := sdk.AccAddressFromBech32(string(creatorAddress))

	if err != nil {
		fmt.Printf("Error converting address string: ", err)
		return nil, err
	}

	err = localContext.AccountRetriever.EnsureExists(*localContext, adminAddr)

	if err != nil {
		fmt.Println("Account does not exist because: ", err)
		return nil, err
	}

	num, seq, err := localContext.AccountRetriever.GetAccountNumberSequence(*localContext, adminAddr)

	if err != nil {
		fmt.Printf("Error retrieving account number/sequence: ", err)
		return nil, err
	} else {
		fmt.Printf("Account retrieved: %v with seq: %v", num, seq)
	}

	//txBuilder := localContext.TxConfig.NewTxBuilder()
	txBuilder.SetMsgs(&group.MsgCreateGroup{
		Admin:    adminAddr.String(),
		Members:  memberList,
		Metadata: nil,
	})

	txBuilder.SetFeeAmount(sdk.Coins{sdk.NewInt64Coin("uregen", 5000)})
	txBuilder.SetGasLimit(50000)

	txFactory := clienttx.Factory{}
	txFactory = txFactory.
		WithChainID(localContext.ChainID).
		WithKeybase(localContext.Keyring).
		WithTxConfig(encCfg.TxConfig)

	// Only needed for "offline" accounts?
	//.WithAccountNumber(num).WithSequence(seq)

	info, err := txFactory.Keybase().Key("delegator")

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
func adminMembers( addr1 string, addr2 string ) []group.Member{
	
	member1 := group.Member{
		Address:  addr1,
		Weight:   strconv.Itoa(1),
	}

	member2 := group.Member{
		Address:  addr2,
		Weight:   strconv.Itoa(1),
	}
	
	return []group.Member{member1, member2}
}

func Init(configPath string) (krplugin.Plugin, error) {

	nc, err := loadConfiguration( configPath )

	if err != nil {
		return nil, err
	} else {
		return nc, nil
	}
}

func (*nodeclient) NewKey(in *pb.KeySpec) (*pb.KeyRef, error) {
	// add a group member to the named group
	return nil, nil
}

func (*nodeclient) PubKey(in *pb.KeySpec) (*pb.PublicKey, error) {
	return nil, nil
}

func (*nodeclient) Sign(in *pb.Msg) (*pb.Signed, error) {
	return nil, nil
}
