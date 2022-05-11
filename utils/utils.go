package utils

import (
	"math"
	"math/big"
	"crypto/rand"
)

// randomBytes returns up to <size> crypto-random bytes
func randomBytes(size int) (blk []byte, err error) {
    blk = make([]byte, size)
    _, err = rand.Read(blk)
    return
}

// randomUint64 returns a crypto-random integer between 0 and the
// maximum possible unsigned 64 bit integer.
func randomUint64() (rnd uint64, err error) {
	bigInt, err := rand.Int( rand.Reader, new(big.Int).SetUint64(math.MaxUint64) )

	if err != nil {
		return 0, err
	}
	
	return bigInt.Uint64(), nil
}
