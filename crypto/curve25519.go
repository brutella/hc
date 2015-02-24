package crypto

import (
	"crypto/rand"
	"github.com/tonnerre/golang-go.crypto/curve25519"
)

const (
	KeySize = 32
)

// Curve25519_GenerateSecretKey returns random bytes.
func Curve25519_GenerateSecretKey() [KeySize]byte {
	var b [KeySize]byte
	rand.Read(b[:])

	return b
}

// Curve25519_PublicKey returns a Curve25519 public key derived from secretKey.
func Curve25519_PublicKey(secretKey [KeySize]byte) [KeySize]byte {
	var k [KeySize]byte
	curve25519.ScalarBaseMult(&k, &secretKey)

	return k
}

// Curve25519_SharedSecret returns a Curve25519 shared secret derived from secretKey and otherPublicKey.
func Curve25519_SharedSecret(secretKey, otherPublicKey [KeySize]byte) [KeySize]byte {
	var k [KeySize]byte
	curve25519.ScalarMult(&k, &secretKey, &otherPublicKey)

	return k
}
