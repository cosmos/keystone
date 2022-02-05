package main

import (
	"log"
	"context"
	"fmt"
	
	"google.golang.org/grpc"
	
	keystonepb "github.com/regen-network/keystone2/keystone"
)

func main() {
	
	fmt.Println("Keystone client ...")

	opts := grpc.WithInsecure()
	
	cc, err := grpc.Dial("localhost:8080", opts)
	
	if err != nil {
		log.Fatal(err)
	}
	
	defer cc.Close()

	client := keystonepb.NewKeyringClient(cc)

	request := &keystonepb.KeySpec{Label: "regen1fyccfg8ylh79ey2qdtx677k568mn0q3pnkajfk"}

	resp, _ := client.NewKey(context.Background(), request)
	
	fmt.Printf("Receive response => [%v]", *resp.Label)
}
