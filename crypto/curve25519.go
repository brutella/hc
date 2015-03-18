package crypto

import (
	"crypto/rand"
	"github.com/tonnerre/golang-go.crypto/curve25519"
)

const (
	KeySize = 32
)

// Curve25519_GeneratePrivateKey returns random bytes.
func Curve25519_GeneratePrivateKey() [KeySize]byte {
	var b [KeySize]byte
	rand.Read(b[:])

	return b
}

// Curve25519_PublicKey returns a Curve25519 public key derived from privateKey.
func Curve25519_PublicKey(privateKey [KeySize]byte) [KeySize]byte {
	var k [KeySize]byte
	curve25519.ScalarBaseMult(&k, &privateKey)

	return k
}

// Curve25519_SharedSecret returns a Curve25519 shared secret derived from privateKey and otherPublicKey.
func Curve25519_SharedSecret(privateKey, otherPublicKey [KeySize]byte) [KeySize]byte {
	var k [KeySize]byte
	curve25519.ScalarMult(&k, &privateKey, &otherPublicKey)

	return k
}
