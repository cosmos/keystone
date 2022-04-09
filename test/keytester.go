package main

import (
	"flag"
	"bufio"
	"fmt"
	"os"
	"log"

	"crypto/rand"
	"google.golang.org/grpc"

	keystonepb "github.com/regen-network/keystone2/keystone"
)

func randomBytes(size int) (blk []byte, err error) {
    blk = make([]byte, size)
    _, err = rand.Read(blk)
    return
}

func initKeys(server string) *keystonepb.KeyringClient {
	fmt.Println("Keystone client ...")
	
	opts := grpc.WithInsecure()
	
	cc, err := grpc.Dial(server, opts)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer cc.Close()

	client := keystonepb.NewKeyringClient(cc)

	return &client
}

func main() {
	
	var createKey bool
	var sign bool
	var algo string
	var profile string
	
	flag.BoolVar(&createKey, "create", false, "create a new key")
	flag.BoolVar(&sign, "sign", false, "sign something with a key")
	flag.StringVar(&algo, "algo", "KEYGEN_SECP256K1", "KEYGEN_SECP256K1 | KEYGEN_SECP256R1")
	flag.StringVar(&profile, "profile", "PROFILE_BC_ECDSA256", "PROFILE_BC_ECDSA256 | PROFILE_ECDSA256")
	flag.Parse()

	client := initKeys("localhost:8080")
	
	if createKey == true {
		fmt.Printf("client: %v", client)
		//client.NewKey( context.Background(), key(
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
