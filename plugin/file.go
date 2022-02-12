package main

import (
	"fmt"
	"errors"
	pb "github.com/regen-network/keystone2/keystone"
)

const PLUGIN_TYPE_FILE_ID = "urn:network.regen.keystone.plugins:file"

var kr *keyring = nil

type keyring struct {
	keystore    *hsmkeys.Pkcs11Keyring
}

func TypeIdentifier() string {
	return PLUGIN_TYPE_FILE_ID
}
	
// Init initializes this keyring using the passed in file path
// which should implement the KeyringPlugin interface @@TODO
func Init(configPath string) error {

}

// NewKey creates a new private key(pair) on the existing keyring that
// is implemented by this plugin. For PKCS11, it will do that on a
// connected HSM
func NewKey(in *pb.KeySpec) (*pb.KeyRef, error) {
	
}

// PubKey returns the bytes of the public key associated with the
// KeyRef that is specified by the input KeySpec
// @@TODO should really have pub key objects created on the HSM so
// they don't need priv key to be first retrieved, put in memory and
// then get the pub key bytes
func PubKey(in *pb.KeySpec) (*pb.PublicKey, error) {

}

// Sign takes a protobuf message containing content (bytes or a
// reference), KeySpec and signs it according to a SigningProfile
func Sign(in *pb.Msg) (*pb.Signed, error) {

}
