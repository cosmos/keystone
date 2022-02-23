package main

import (
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"

	keystonepb "github.com/regen-network/keystone2/keystone"
)

/*

   Prototyping this to see if it fits...

type keyring interface {

	// Returns an Item matching the key or ErrKeyNotFound
	Key(keyRef int32) (*KeyRef, error)

	// Returns the non-secret parts of an Item
	KeyMetadata(keyRef int32) (*KeyMetadata, error)

	// Stores an Item on the keyring
	NewKey(*KeySpec) error

	// Removes the item with matching key
	Remove(keyRef int32) error

	// Signs input bytes with the referenced key
	Sign(
	// Provides a slice of all keys stored on the keyring
	// by this entity
	Keys() ([]*keystonepb.KeyRef, error)

}

*/

func main() {

	fmt.Println("Keystone client ...")

	opts := grpc.WithInsecure()

	cc, err := grpc.Dial("localhost:8080", opts)

	if err != nil {
		log.Fatal(err)
	}

	defer cc.Close()

	client := keystonepb.NewKeyringClient(cc)

	request := &keystonepb.KeySpec{Label: "abcde12334",
		Algo: keystonepb.KeygenAlgorithm_KEYGEN_SECP256R1,}

	resp, _ := client.NewKey(context.Background(), request)

	fmt.Printf("Receive response => [%v]\n", *resp.Label)

	req2 := &keystonepb.KeySpec{Label: "abcde12334"}
	resp2, err := client.Pubkey(context.Background(), req2)
	fmt.Printf("Receive response => [%v]\n", *resp2)
}
