package main

import (
	"fmt"
	"errors"
	pb "github.com/regen-network/keystone2/keystone"
	hsmkeys "github.com/regen-network/keystone/keys"
)

var kr *keyring = nil

type keyring struct {
	keystore    *hsmkeys.Pkcs11Keyring
}

// Init initializes this keyring using the passed in file path
// which should implement the KeyringPlugin interface @@TODO
func Init(configPath string) error {

	k11, err := hsmkeys.NewPkcs11FromConfig(configPath)

	if err != nil {
		return err
	} else {
		kr = &keyring{
			keystore: k11,
		}
		return nil
	}
}

// NewKey creates a new private key(pair) on the existing keyring that
// is implemented by this plugin. For PKCS11, it will do that on a
// connected HSM
func NewKey(in *pb.KeySpec) (*pb.KeyRef, error) {
	
	fmt.Println("PKCS11 plugin NewKey")
	
	if kr.keystore != nil {
		
		// @@TODO: support more than just this key type, based
		// on input KeySpec
		key, err := kr.keystore.NewKey(hsmkeys.KEYGEN_SECP256K1, string(in.Label))

		if err != nil {
			return nil, err
		}
		
		ref := pb.KeyRef{
			Label: &key.Label,
		}
	
		return &ref, nil
	} else {
		return nil, errors.New("Keystore not initialized!")
	}
}


