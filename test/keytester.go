package main

import (
	"flag"
	"bufio"
	"fmt"
	"os"
	"log"
	"crypto/ed25519"
	
	"golang.org/x/crypto/ssh/terminal"
	
	"github.com/cosmos/keystone/utils"
	keystonepb "github.com/cosmos/keystone/keystone"
	//keystoneadminpb "github.com/cosmos/keystone/keystone_admin"
	"github.com/cosmos/keystone/client"
)

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
	
	flag.BoolVar(&createKey, "create", false, "create a new key")
	flag.BoolVar(&sign, "sign", false, "sign something with a key")
	flag.BoolVar(&pubkey, "key", false, "get the public key for a private key")
	flag.BoolVar(&createKeyring, "keyring", false, "create a new keyring")
	
	flag.StringVar(&algo, "algo", "KEYGEN_SECP256K1", "KEYGEN_SECP256K1 | KEYGEN_SECP256R1")
	flag.StringVar(&profile, "profile", "PROFILE_BC_ECDSA256", "PROFILE_BC_ECDSA256 | PROFILE_ECDSA256")

	flag.Parse()

	fmt.Println("Enter a password to encrypt data in transit: ")
	pwrod, err := terminal.ReadPassword(int(os.Stdout.Fd()))

	var clientKey ed25519.PrivateKey
	var clientSalt []byte
	
	if err == nil {
		clientKey, clientSalt, err = utils.KeyFrom(pwrod, nil)

		if err != nil {
			log.Fatalf("Could not create server key")
		}
	}

	ks, err := client.Keystore("localhost:8080", clientKey, clientSalt)

	if err != nil {
		log.Fatalf( "Could not connect to key server: %v", err )
	}
	
	if createKeyring == true {

		kr, err := ks.NewKeyring()

		if err != nil {
			log.Fatalf("Could not create keyring: %v", err)
		}
		
		fmt.Printf("New keyring: %s\n", kr.Label())
	}
	
	if createKey == true {
		//fmt.Printf("client: %v", client)

		key, err := ks.NewKey(keystonepb.KeygenAlgorithm_KEYGEN_SECP256R1, "0" )
		
		if err != nil {
			log.Fatalf("Error creating key: %v", err)
		}

		fmt.Printf("New key: %s\n", key.Label())
	}

	if pubkey == true {
		keyname := flag.Args()[0]

		
		if len(keyname) == 0 {
			fmt.Println("Usage: program_name [-key] keyname")
			flag.PrintDefaults()
			os.Exit(1)
		} else {
			privateKey, err := ks.Key( keyname )
			pubKey, err := privateKey.PubKey()

			if err != nil {
				log.Fatalf("error retrieving key: %v", err)
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
