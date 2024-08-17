package hashing

import (
	"crypto/sha256"
	"encoding/hex"
)

// Function ValidFile implements sha256 file encryption system to ensure correctness, authenticity and intergrity of every banner file
func ValidFile(bannerFileData []byte) string {
	hasher := sha256.New()
	hasher.Write(bannerFileData)
	hashInBytes := hasher.Sum(nil)

	fileHash := hex.EncodeToString(hashInBytes)

	return fileHash
}
