package hash

import (
	"crypto/sha256"
)

// GenerateChecksum generates the checksum of piece data
func GenerateChecksum(pieceData []byte) []byte {
	hash := sha256.New()
	hash.Write(pieceData)
	return hash.Sum(nil)
}

// GenerateIntegrityHash generate integrity hash of ec data
func GenerateIntegrityHash(checksumList [][]byte) []byte {
	hash := sha256.New()
	for _, j := range checksumList {
		hash.Write(j)
	}
	return hash.Sum(nil)
}
