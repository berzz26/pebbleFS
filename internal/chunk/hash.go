package chunk

import (
	"crypto/sha256"
	"encoding/hex"
)

// this module implements CAS (content addressable storage)
// the entire idea of this is that the content determines the address.

func Hash(data []byte) string {
	sum := sha256.Sum256(data)

	return hex.EncodeToString(sum[:])
}