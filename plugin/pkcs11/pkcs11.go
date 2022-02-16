package main

import (
	"fmt"
	"errors"
	pb "github.com/regen-network/keystone2/keystone"
	hsmkeys "github.com/regen-network/keystone/keys"
)

const Plugin_Type_Pkcs11_Id = "urn:network.regen.keystone.plugins:pkcs11"

var kr *keyring = nil

type keyring struct {
	keystore    *hsmkeys.Pkcs11Keyring
}

func TypeIdentifier() string {
	return Plugin_Type_Pkcs11_Id
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
		// @@TODO standardize on hsmkeys algo identifiers also in the PB
		
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

// PubKey returns the bytes of the public key associated with the
// KeyRef that is specified by the input KeySpec
// @@TODO should really have pub key objects created on the HSM so
// they don't need priv key to be first retrieved, put in memory and
// then get the pub key bytes
func PubKey(in *pb.KeySpec) (*pb.PublicKey, error) {
	key, err := kr.keystore.Key( string( in.Label ))

	if err != nil {
		return nil, err
	} else {
		public := pb.PublicKey{
			KeyBytes: key.PubKeyBytes(),
		}
		return &public, nil
	}
}

// Sign takes a protobuf message containing content (bytes or a
// reference), KeySpec and signs it according to a SigningProfile
func Sign(in *pb.Msg) (*pb.Signed, error) {

	keyLabel := in.KeySpec.Label
	key, err := kr.keystore.Key( string( keyLabel ))

	if err != nil {
		return nil, err
	}
	
	switch in.Content.Data.(type) {
	case *pb.Signable_SignableBytes:
		cleartext := in.Content.GetSignableBytes()

		signature, err := key.Sign( cleartext, nil )

		if err != nil {
			return nil, err
		}
		
		signedBytes := pb.Signed_SignedBytes{
			SignedBytes: signature,
		}

		signed := pb.Signed{
			Data: &signedBytes,
		}
		
		return &signed, nil
	default:
		return nil, errors.New("Cannot sign these data")
	}
}
	
