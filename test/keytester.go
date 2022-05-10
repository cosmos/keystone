package main

import (
	"flag"
	"bufio"
	"fmt"
	"os"
	"log"
	"context"
	"math"
	"math/big"

	"crypto/rand"
	"google.golang.org/grpc"

	keystonepb "github.com/regen-network/keystone2/keystone"
	keystoneadminpb "github.com/regen-network/keystone2/keystone_admin"
)

// randomBytes returns up to <size> crypto-random bytes
func randomBytes(size int) (blk []byte, err error) {
    blk = make([]byte, size)
    _, err = rand.Read(blk)
    return
}

// randomUint64 returns a crypto-random integer between 0 and the
// maximum possible unsigned 64 bit integer.
func randomUint64() (rnd uint64, err error) {
	bigInt, err := rand.Int( rand.Reader, new(big.Int).SetUint64(math.MaxUint64) )

	if err != nil {
		return 0, err
	}
	
	return bigInt.Uint64(), nil
}

// initKeys initializes the connection to the keystone server
func initKeys(server string) *grpc.ClientConn {

	opts := grpc.WithInsecure()
	
	cc, err := grpc.Dial(server, opts)
	
	if err != nil {
		log.Fatal(err)
	}

	return cc
}

// main parses command line flags for the various options and then
// constructs requests to the keystone server, checks responses and
// presents a minimal cli ui to the user
func main() {
	
	var createKey bool
	var sign bool
	var pubkey bool
	var createKeyring bool
	
	var algo string
	var profile string
	var cc *grpc.ClientConn
	
	flag.BoolVar(&createKey, "create", false, "create a new key")
	flag.BoolVar(&sign, "sign", false, "sign something with a key")
	flag.BoolVar(&pubkey, "key", false, "get the public key for a private key")
	flag.BoolVar(&createKeyring, "keyring", false, "create a new keyring")
	
	flag.StringVar(&algo, "algo", "KEYGEN_SECP256K1", "KEYGEN_SECP256K1 | KEYGEN_SECP256R1")
	flag.StringVar(&profile, "profile", "PROFILE_BC_ECDSA256", "PROFILE_BC_ECDSA256 | PROFILE_ECDSA256")

	flag.Parse()
	
	cc = initKeys("localhost:8080")
	client := keystonepb.NewKeyringClient(cc)
	adminClient := keystoneadminpb.NewKeyringAdminClient(cc)
	
	defer cc.Close()

	if createKeyring == true {
		labelBytes, err := randomBytes(16)

		if err != nil {
			log.Fatalf("Error getting randomness")
		}
		
		id, err := randomUint64()	

		if err != nil {
			log.Fatalf("Error getting randomness")
		}
			
		label := fmt.Sprintf("%x",labelBytes)

		if err != nil {
			log.Fatalf("Error creating key: %v", err)
		}
		
		request := &keystoneadminpb.KeyringSpec{Id: id, Label: label,}

		keyringRef, err := adminClient.NewKeyring( context.Background(), request )
		
		if err != nil {
			log.Fatalf("Error creating keyring: %v", err)
		}

		fmt.Printf("New keyring: %s\n", keyringRef.Label)
	}
	
	if createKey == true {
		//fmt.Printf("client: %v", client)

		labelBytes, err := randomBytes(16)
		
		label := fmt.Sprintf("%x",labelBytes)

		if err != nil {
			log.Fatalf("Error creating key: %v", err)
		}
		
		request := &keystonepb.KeySpec{Label: label,
			Algo: keystonepb.KeygenAlgorithm_KEYGEN_SECP256R1,}
		
		keyref, err := client.NewKey( context.Background(), request )
		
		if err != nil {
			log.Fatalf("Error creating key: %v", err)
		}

		fmt.Printf("New key: %s\n", keyref.Label)
	}

	if pubkey == true {
		keyname := flag.Args()[0]

		if len(keyname) == 0 {
			fmt.Println("Usage: program_name [-key] keyname")
			flag.PrintDefaults()
			os.Exit(1)
		} else {
			request := &keystonepb.KeySpec{Label: keyname}
			
			pubKey, err := client.PubKey( context.Background(), request )
			
			if err != nil {
				log.Fatalf("Error getting key: %v", err)
			}

			fmt.Printf("%x\n", pubKey.KeyBytes)
		}
		
	}
	
	if sign == true {
		keyname := flag.Args()

		fmt.Printf("values: %v", keyname)
		
		if len(keyname) == 0 {
			fmt.Println("Usage: program_name [-sign] keyname")
			flag.PrintDefaults()
			os.Exit(1)
		}

		scanner := bufio.NewScanner(os.Stdin)
		
		for scanner.Scan() {
			fmt.Println(scanner.Text())
		}
		
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, "reading standard input:", err)
		}

	}



}
