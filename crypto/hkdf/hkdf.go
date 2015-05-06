package hkdf

import (
	"crypto/sha512"
	"golang.org/x/crypto/hkdf"
	"io"
)

// Sha512 returns a 256-bit key
func Sha512(master, salt, info []byte) ([32]byte, error) {
	hash := sha512.New
	hkdf := hkdf.New(hash, master, salt, info)

	key := make([]byte, 32) // 256 bit
	_, err := io.ReadFull(hkdf, key)

	var result [32]byte
	copy(result[:], key)

	return result, err
}
