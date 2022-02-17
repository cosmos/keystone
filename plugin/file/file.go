package main

import (
	"crypto"
	"crypto/ecdsa"
	"crypto/x509"
	"crypto/elliptic"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"os"
	"crypto/rand"
	"encoding/pem"
	"log"

	pb "github.com/regen-network/keystone2/keystone"
	"github.com/frumioj/crypto11"
)

const Plugin_Type_File_Id = "urn:network.regen.keystone.plugins:file"

// PRIVATE functions

// encodeKeypair takes an ECDSA keypair as a private and a public key,
// and uses the x509 routines to return PEM encoded strings.
func encodeKeypair(privateKey *ecdsa.PrivateKey, publicKey *ecdsa.PublicKey) (string, string, error) {
	x509Encoded, err := x509.MarshalECPrivateKey(privateKey)
	
	if err != nil {
		return "", "", err
	}
	
	pemEncoded := pem.EncodeToMemory(&pem.Block{Type: "EC PRIVATE KEY", Bytes: x509Encoded})
	
	x509EncodedPub, err := x509.MarshalPKIXPublicKey(publicKey)

	if err != nil {
		return "", "", err
	}
	
	pemEncodedPub := pem.EncodeToMemory(&pem.Block{Type: "PUBLIC KEY", Bytes: x509EncodedPub})
	
	return string(pemEncoded), string(pemEncodedPub), nil
}

func decodeKeypair(pemEncoded string, pemEncodedPub string) (*ecdsa.PrivateKey, crypto.PublicKey, error) {
	block, _ := pem.Decode([]byte(pemEncoded))

	x509Encoded := block.Bytes
	privateKey, err := x509.ParseECPrivateKey(x509Encoded)

	if err != nil {
		return nil, nil, err
	}
		
	publicKey, err := decodePubkeyPem( pemEncodedPub )

	if err != nil {
		return nil, nil, err
	}
	
	return privateKey, publicKey, nil
}

func decodePubkeyPem( pemEncodedPub string ) (crypto.PublicKey, error ) {
	blockPub, _ := pem.Decode([]byte(pemEncodedPub))

	x509EncodedPub := blockPub.Bytes
	publicKey, err := x509.ParsePKIXPublicKey(x509EncodedPub)

	if err != nil {
		return nil, err
	}

	//publicKey := genericPublicKey.(*ecdsa.PublicKey)

	return publicKey, nil 
}
	
func writeKeys( kr string, filePrefix string, pemEncodedKey string, pemEncodedPub string ) (error) {
	log.Printf("Creating key in %s", kr + "/" + filePrefix + ".pem")
	
	err := os.WriteFile(kr + "/" + filePrefix + ".pem", []byte(pemEncodedKey), 0600)

	if err != nil {
		return errors.New("Could not create key")
	}
	
	log.Printf("Creating key in %s", kr + "/" + filePrefix + ".pem")
	
	err = os.WriteFile(kr + "/" + filePrefix + "pub.pem", []byte(pemEncodedPub), 0600)

	if err != nil {
		return errors.New("Could not create key")
	}

	return nil
}

func readKeys( kr string, filePrefix string ) ( pemEncodedKey string, pemEncodedPub string, err error) {	
	privateBytes, err := os.ReadFile(kr + "/" + filePrefix + ".pem")

	if err != nil {
		return "", "", errors.New("Could not retrieve key")
	}

	publicBytes, err := readPubKey( kr, filePrefix )

	if err != nil {
		return "", "", errors.New("Could not retrieve key")
	}

	return string( privateBytes ), publicBytes, nil
}

func readPubKey( kr string, filePrefix string ) ( pemEncodedPub string, err error) {	
	publicBytes, err := os.ReadFile(kr + "/" + filePrefix + "pub.pem")

	if err != nil {
		return "", errors.New("Could not retrieve key")
	}

	return string( publicBytes ), nil
}

// This keyring just uses the filesystem to store keys

var kr *keyring = nil

type keyring struct {
	keystorePath string
}

// PUBLIC interface for a plugin

func TypeIdentifier() string {
	return Plugin_Type_File_Id
}
	
// Init initializes this keyring using the passed in file path
// which should implement the KeyringPlugin interface @@TODO
func Init(configPath string) error {
	err := os.MkdirAll(configPath, os.ModePerm)

	if err != nil && !os.IsExist(err) {
		return errors.New("Could not initialize keystore")
	}

	kr = &keyring{
		keystorePath: configPath,
	}

	return nil
}

// NewKey creates a new private key(pair) on the existing keyring that
// is implemented by this plugin. This creates a new keypair, and
// stores each key in a file on the filesystem.
func NewKey(in *pb.KeySpec) (*pb.KeyRef, error) {

	var keygenSpec elliptic.Curve

	filenamePrefix := string(in.Label)
	
	if in.Algo == pb.KeygenAlgorithm_KEYGEN_SECP256R1 {
		keygenSpec = elliptic.P256()
	} else {
		if in.Algo == pb.KeygenAlgorithm_KEYGEN_SECP256K1 {
			keygenSpec = crypto11.P256K1()
		} else {
			return nil, errors.New("Could not create key")
		}
	}
	
	privateKey, _ := ecdsa.GenerateKey(keygenSpec, rand.Reader)
	publicKey := &privateKey.PublicKey
	pemEncodedKey, pemEncodedPub, err := encodeKeypair( privateKey, publicKey )

	if err != nil {
		return nil, errors.New("Could not create key")
	}

	// even though kr is a global var, I pass it in
	// to make that state explicit
	err = writeKeys( kr.keystorePath, filenamePrefix, pemEncodedKey, pemEncodedPub )

	if err != nil {
		return nil, errors.New("Could not create key")
	}	

	// @@TODO: what should the label be really?
	ref := pb.KeyRef{
		Label: &in.Label,
	}

	return &ref, nil
}

// PubKey returns the bytes of the public key associated with the
// KeyRef that is specified by the input KeySpec
// @@TODO should really have pub key objects created on the HSM so
// they don't need priv key to be first retrieved, put in memory and
// then get the pub key bytes
func PubKey(in *pb.KeySpec) (*pb.PublicKey, error) {
	pemEncodedPub, err := readPubKey( kr.keystorePath, string(in.Label) )
	key, err := decodePubkeyPem( pemEncodedPub )

	if err != nil {
		return nil, err
	}
	
	switch k := key.(type) {
	case *ecdsa.PublicKey:
		// @@TODO: check the curve params for which curve it is first (although outcome is same for both k1 and r1)
		// is this OK for a *btcec* secp256k1 key?
		public := pb.PublicKey{
			KeyBytes: elliptic.MarshalCompressed(k.Curve, k.X, k.Y),
		}
		return &public, nil
	default:
		return nil, errors.New("Unsupported public key type!")
	}



}

// Sign takes a protobuf message containing content (bytes or a
// reference), KeySpec and signs it according to a SigningProfile
func Sign(in *pb.Msg) (*pb.Signed, error) {
	pemEncodedPriv, pemEncodedPub, err := readKeys( kr.keystorePath, string(in.KeySpec.Label) )
	
	if err != nil {
		return nil, err
	}

	// @@TODO - pub key can be returned - should the bytes go in the response message?
	// ignored for now
	key, _, err := decodeKeypair( pemEncodedPriv, pemEncodedPub )

	if err != nil {
		return nil, err
	}

	switch in.Content.Data.(type) {
	case *pb.Signable_SignableBytes:
		cleartext := in.Content.GetSignableBytes()

		var digested []byte
		
		if in.SigningProfile == pb.SigningProfile_PROFILE_ECDSA_SHA256 ||
			in.SigningProfile == pb.SigningProfile_PROFILE_BC_ECDSA_SHA256 {
			digest := sha256.Sum256( cleartext )
			digested = digest[:]
		} else {
			if in.SigningProfile == pb.SigningProfile_PROFILE_BC_ECDSA_SHA512 {
				digest := sha512.Sum512( cleartext )
				digested = digest[:]
			} else {
				digested = cleartext
			}
		}
		
		signature, err := key.Sign( rand.Reader, digested, nil )
		
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
