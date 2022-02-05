package main

import (
	"context"
	"log"
	"net"
	"flag"

	"google.golang.org/grpc"

	pb "github.com/regen-network/keystone2/keystone"
	hsmkeys "github.com/regen-network/keystone/keys"
)

type Server = pb.KeyringServer

type server struct{
	pb.UnimplementedKeyringServer
	ServerAddress    string
	ChainID          string
	KeyringType      string
	KeyringDir       string
	RpcURI           string
	Keystore         *hsmkeys.Pkcs11Keyring
}

func New() (Server, error) {
	return &server{}, nil
}

// NewKey implements the method given in the protobuf definition for
// the Keystone service (proto/keystone.proto)
func (s *server) NewKey(ctx context.Context, in *pb.KeySpec) (*pb.KeyRef, error) {
	log.Printf("Receive message body from client: %v", in)

	newLabel := "urn:network.regen.keystone:keystore123:abcde123"
	return &pb.KeyRef{Label: &newLabel}, nil
}

func (s *server) Key(ctx context.Context, in *pb.KeySpec) (*pb.KeyRef, error) {
	log.Printf("Receive message body from client: %v", in)

	newLabel := "abcde123"
	return &pb.KeyRef{Label: &newLabel}, nil
}

func main() {

	// Retrieve the command line parameters passed in to configure the server
	// Most have likely-reasonable defaults.
	keystoneAddress := flag.String("key-addr", "", "the address associated with the key used to sign transactions on behalf of Keystone")
	blockchain := flag.String("chain-id", "test-chain", "the blockchain that Keystone should connect to")
	keyringType := flag.String("keyring-type", "test", "the keyring backend type where keys should be read from")
	keyringDir := flag.String("keyring-dir", "~/.regen/", "the directory where the keys are")
	chainRpcURI := flag.String("chain-rpc", "tcp://localhost:26657", "the address of the RPC endpoint to communicate with the blockchain")
	grpcListenPort := flag.String("listen-port", "8080", "the port where the server will listen for connections")
	pkcs11KeyringConfig := flag.String("pkcsll-cfg", "./pkcs11-config", "configuration file for PKCS11 HSM connection")

	flag.Parse()

	if len(*keystoneAddress) <= 0 {
		log.Fatalln("Keystone server blockchain address may not be left empty")
		return
	}

	kr, err := hsmkeys.NewPkcs11FromConfig(*pkcs11KeyringConfig)

	if err != nil {
		log.Fatalln("Failed to initialize keystore")
		return
	}

	lis, err := net.Listen("tcp", ":" + *grpcListenPort)

	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create new server context, used for passing server-global state
	ss := server{
		ServerAddress: *keystoneAddress,
		ChainID: *blockchain,
		KeyringType: *keyringType,
		KeyringDir: *keyringDir,
		RpcURI: *chainRpcURI,
		Keystore: kr,
	}
	
	s := grpc.NewServer()
	pb.RegisterKeyringServer(s, &ss)

	s.Serve(lis)
	return

}
