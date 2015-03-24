package crypto

import (
	"bytes"
	"github.com/agl/ed25519"

	"fmt"
)

// ValidateED25519Signature return true when the ED25519 signature is a valid signature of the data based on the key, otherwise false.
func ValidateED25519Signature(key, data, signature []byte) bool {
	if len(key) != ed25519.PublicKeySize || len(signature) != ed25519.SignatureSize {
		return false
	}

	var k [ed25519.PublicKeySize]byte
	var s [ed25519.SignatureSize]byte
	copy(k[:], key)
	copy(s[:], signature)

	return ed25519.Verify(&k, data, &s)
}

// ED25519Signature returns the ED25519 signature of data using the key.
func ED25519Signature(key, data []byte) ([]byte, error) {
	if len(key) != ed25519.PrivateKeySize {
		return nil, fmt.Errorf("Invalid size of key (%v)", len(key))
	}

	var k [ed25519.PrivateKeySize]byte
	copy(k[:], key)
	signature := ed25519.Sign(&k, data)

	return signature[:], nil
}

// ED25519GenerateKey return a public and private ED25519 key pair from a string.
func ED25519GenerateKey(str string) ([]byte /* public */, []byte /* private */, error) {
	b := bytes.NewBuffer([]byte(str))
	if len(str) < 32 {
		zeros := make([]byte, 32-len(str))
		b.Write(zeros)
	}

	public, private, err := ed25519.GenerateKey(bytes.NewReader(b.Bytes()))

	return public[:], private[:], err
}
