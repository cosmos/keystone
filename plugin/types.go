package types

import (
	"errors"
	"math/big"
	"encoding/asn1"
	"crypto/elliptic"

	pb "github.com/regen-network/keystone2/keystone"
)

const Plugin_Type_Pkcs11_Id = "urn:network.regen.keystone.plugins:pkcs11"
const Plugin_Type_File_Id = "urn:network.regen.keystone.plugins:file"
const Plugin_Type_CosmosG_Id = "urn:network.regen.keystone.plugins:cosmosg"

// Options struct contains anything that is needed for configuring a
// plugin with one element initially, a path to a configuration, which
// can be interpreted differently by an individual plugin
type Options struct {
	ConfigPath string
}

// Plugin interface specifies the methods required for implementation
// by a plugin
type Plugin interface {
	NewKey(in *pb.KeySpec) (*pb.KeyRef, error)
	PubKey(in *pb.KeySpec) (*pb.PublicKey, error)
	Sign(in *pb.Msg) (*pb.Signed, error)
}

// dsaSignature contains the two integers needed for
// an ECDSA signature value. They must be put in a struct
// to allow the asn1 unmarshalling which uses an interface{}
// type to return the values, instead of just returning the
// two integers.
type DsaSignature struct {
	R, S *big.Int
}

// SHARED functions for use by more than one plugin

// unmarshalDER takes a DER-encoded byte array, and dumps
// it into a (hopefully-appropriate) struct. If the struct
// given, is not appropriate for the data, then unmarshalling
// will fail.
func UnmarshalDER(sigDER []byte) (*DsaSignature, error) {
	var sig DsaSignature

	if rest, err := asn1.Unmarshal(sigDER, &sig); err != nil {
		return nil, err
	} else if len(rest) > 0 {
		return nil, errors.New("unexpected data found after DSA signature")
	}

	return &sig, nil
}

// isSNormalized returns true for the integer sigS if sigS falls in
// lower half of the curve order
// It is expected that the caller passes the curve order as a big Int along
// with the s portion of the signature.
func IsSNormalized(sigS *big.Int, order *big.Int) bool {
	// return the result of comparing the given s signature
	// component with half the value of the curve order. If the s
	// component is less than or equal to half the curve order,
	// then returns true (!= 1), if > than, will return false
	// (==1)
	return sigS.Cmp(new(big.Int).Rsh(order, 1)) != 1
}

// NormalizeS will invert the s value if not already in the lower half
// of curve order value by subtracting it from the curve order (N)
func NormalizeS(sigS *big.Int, curve elliptic.Curve) *big.Int {
	if IsSNormalized(sigS, curve.Params().N) {
		return sigS
	} else {
		order := curve.Params().N
		return new(big.Int).Sub(order, sigS)
	}
}

// signatureRaw takes two big integers and returns a byte value that
// is the result of concatenating the byte values of each of the given
// integers. The byte values are left-padded with zeroes
func SignatureRaw(r *big.Int, s *big.Int) []byte {

	rBytes := r.Bytes()
	sBytes := s.Bytes()
	sigBytes := make([]byte, 64)
	// 0 pad the byte arrays from the left if they aren't big enough.
	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)
	return sigBytes
}
