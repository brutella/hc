package curve25519

import (
	"crypto/rand"
	"golang.org/x/crypto/curve25519"
)

const (
	keySize = 32
)

// GeneratePrivateKey returns random bytes.
func GeneratePrivateKey() [keySize]byte {
	var b [keySize]byte
	rand.Read(b[:])

	return b
}

// PublicKey returns a Curve25519 public key derived from privateKey.
func PublicKey(privateKey [keySize]byte) [keySize]byte {
	var k [keySize]byte
	curve25519.ScalarBaseMult(&k, &privateKey)

	return k
}

// SharedSecret returns a Curve25519 shared secret derived from privateKey and otherPublicKey.
func SharedSecret(privateKey, otherPublicKey [keySize]byte) [keySize]byte {
	var k [keySize]byte
	curve25519.ScalarMult(&k, &privateKey, &otherPublicKey)

	return k
}
