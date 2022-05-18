package utils

import (
	"math"
	"math/big"
	"crypto/rand"
	"crypto/ed25519"
	"golang.org/x/crypto/scrypt"
	"golang.org/x/crypto/blake2b"
)

// randomBytes returns up to <size> crypto-random bytes
func RandomBytes(size int) (blk []byte, err error) {
    blk = make([]byte, size)
    _, err = rand.Read(blk)
    return
}

// RandomUint64 returns a crypto-random integer between 0 and the
// maximum possible unsigned 64 bit integer.
func RandomUint64() (rnd uint64, err error) {
	bigInt, err := rand.Int( rand.Reader, new(big.Int).SetUint64(math.MaxUint64) )

	if err != nil {
		return 0, err
	}
	
	return bigInt.Uint64(), nil
}

// KeyFrom generates an ed25519 keypair from an input password by
// first hashing the input and then inputting the resulting hash into
// scrypt with recommended params. Finally, this scrypted key is used
// as the seed for an ed25519 key pair. Caller must store the salt in
// addition to the pwrod if it's needed
func KeyFrom(pwrod []byte, insalt []byte) (key ed25519.PrivateKey, salt []byte, err error) {
	stuff := blake2b.Sum256(pwrod)

	if insalt != nil && len(insalt) >= 8 {
		salt = insalt
	} else {
		salt, err = RandomBytes(8)

		if err != nil {
			return nil, nil, err
		}
	}

	// recommended values for scrypt as of 2017, 2^15 = 32768
	scryptkey, err := scrypt.Key( stuff[:], salt, 1<<15, 8, 1, 32 )

	if err != nil || len(scryptkey) < ed25519.SeedSize {
		return nil, nil, err
	}
	
	return ed25519.NewKeyFromSeed( scryptkey ), salt, nil
}
