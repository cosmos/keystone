package client

import (
	"log"
	"fmt"
	"context"
	"crypto/ed25519"
	
	"google.golang.org/grpc"
	
	"github.com/cosmos/keystone/utils"
	keystonepb "github.com/cosmos/keystone/keystone"
	keystoneadminpb "github.com/cosmos/keystone/keystone_admin"
)	

var ksClients = make(map[string]*keystore)

// initKeys initializes the connection to the keystone server
func initKeys(server string) (cc *grpc.ClientConn, err error) {

	opts := grpc.WithInsecure()
	
	cc, err = grpc.Dial(server, opts)
	
	if err != nil {
		return nil, err
	}

	return cc, nil
}

// Keystore is a proxy for a "keystone server" which may hold one or
// more keyrings.
type keystore struct {
	client *keystonepb.KeyringClient
	adminclient *keystoneadminpb.KeyringAdminClient
	url string
	clientKey ed25519.PrivateKey
	clientSalt []byte
}

// keyring is a set of keys, which may be used in a multi-signature or
// other crypto transaction using multiple keys, but is also "just" a
// set of keys which may even be different versions of the "same" key,
// maintained that way so older keys can still be used for decryption
type keyring struct {
	keyringRef *keystoneadminpb.KeyringRef
}

// key is a proxy reference to a Keystone key stored on the
// server. Keys may be hosted on one Keystone server, while the
// Keyring that refers to them may be on a different keystone server
type key struct {
	keyRef *keystonepb.KeyRef
	keyringRef *keyring
}

type publicKey struct {
	keystonepb.PublicKey
}

type Keyring interface {
	NewKey(algo keystonepb.KeygenAlgorithm) (*key, error)
	Label() string
	//Sign(msg []byte, profile keystonepb.SigningProfile) ([]byte, error)
}

type Key interface {
	Label() string
	Sign(msg [] byte, profile keystonepb.SigningProfile) ([]byte, error)
	PubKey() (*publicKey, error)
}

func Keystore(keystoneUrl string, clientKey ed25519.PrivateKey, salt []byte) (ks *keystore, err error){

	if ksClients[keystoneUrl] == nil {
		cc, err := initKeys(keystoneUrl)
		
		if err != nil {
			return nil, err
		}
		
		client := keystonepb.NewKeyringClient(cc)
		adminClient := keystoneadminpb.NewKeyringAdminClient(cc)
		
		keystore := keystore{
			client: &client,
			adminclient: &adminClient,
			url: keystoneUrl,
			clientKey: clientKey,
			clientSalt: salt,
		}

		ksClients[keystoneUrl] = &keystore
	}
	
	return ksClients[keystoneUrl], nil
}

func (ks *keystore) NewKey(algo keystonepb.KeygenAlgorithm, krId string) (k *key, err error) {

	//@@TODO: add to an existing specified keyring, or create an
	// empty one, and then refr this key to that keyring
	
	labelBytes, err := utils.RandomBytes(16)
	
	label := fmt.Sprintf("%x",labelBytes)
	
	if err != nil {
		return nil, err
	}
	
	request := &keystonepb.KeySpec{
		Label: label,
		Algo: algo,
	}

	// if request.KeyringId is not set, then server should create
	// a new keyring
	if len(krId) > 0 {
		request.KeyringId = krId
	}
	
	client := *ks.client
	keyref, err := client.NewKey( context.Background(), request )
	
	if err != nil {
		return nil, err
	}
	
	fmt.Printf("New key: %s\n", keyref.Label)

	return &key{ keyRef: keyref }, nil
}

func (ks *keystore) Key(label string) (k *key, err error) {

	request := &keystonepb.KeySpec{
		Label: label,
	}

	client := *ks.client
	keyref, err := client.Key( context.Background(), request )
	
	if err != nil {
		return nil, err
	}
	
	fmt.Printf("Key: %s\n", keyref.Label)

	return &key{ keyRef: keyref }, nil
}

func (ks *keystore) NewKeyring() (*keyring, error) {
	labelBytes, err := utils.RandomBytes(16)
	
	if err != nil {
		log.Fatalf("Error getting randomness")
	}
	
	id, err := utils.RandomUint64()	
	
	if err != nil {
		log.Fatalf("Error getting randomness")
	}
	
	label := fmt.Sprintf("%x",labelBytes)
	
	if err != nil {
		log.Fatalf("Error creating key: %v", err)
	}
	
	request := &keystoneadminpb.KeyringSpec{Id: id, Label: label,}

	ac := *ks.adminclient
	keyringRef, err := ac.NewKeyring( context.Background(), request )
	
	if err != nil {
		log.Fatalf("Error creating keyring: %v", err)
	}
	
	fmt.Printf("New keyring: %s\n", keyringRef.Label)
	
	//@@TODO: create the keyring on keystone backend
	return &keyring{
		keyringRef: keyringRef,
	}, nil

}

func (k *key) PubKey() (*publicKey, error) {

	// @@TODO This should be for the case where public key is extractable
	// from the private key bytes - is that all cases?
	
	request := &keystonepb.KeySpec{Label: k.Label()}

	// get the keystore reference from the issuer URL but if
	// that's not found, it will try to contact that URL to create
	// a new keystore connection NOTE: key and salt will be nil
	// because the user hasn't supplied either of them to this
	// method. Will perhaps need to think about that in context of
	// PubKey being accessed from a "dead" KeyRef (one whose
	// keystore is not currently connected)
	ks, err := Keystore(k.keyRef.IssuerUrl, nil, nil)

	if err != nil {
		return nil, err
	}
	
	client := *ks.client
	
	pubKey, err := client.PubKey( context.Background(), request )
	
	if err != nil {
		return nil, err
	}

	return &publicKey{*pubKey}, nil
}

func (kr *keyring) Label() string {
	return kr.keyringRef.Label
}

func (k *key) Label() string {
	return k.keyRef.Label
}
