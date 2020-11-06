package chacha20poly1305

import (
	"errors"

	"golang.org/x/crypto/chacha20poly1305"
	"golang.org/x/crypto/poly1305"
)

// DecryptAndVerify returns the chacha20 decrypted messages.
// An error is returned when the poly1305 message authenticator (seal) could not be verified.
// Nonce should be 8 byte.
func DecryptAndVerify(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key size")
	}
	if len(nonce) != 8 {
		return nil, errors.New("invalid nonce size")
	}

	var (
		Nonce   [12]byte
		aeadOut = make([]byte, len(message))
	)
	copy(Nonce[4:], nonce)

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, err
	}

	return aead.Open(aeadOut[:0], Nonce[:], append(message, mac[:]...), add)
}

// EncryptAndSeal returns the chacha20 encrypted message and poly1305 message authentictor (also referred as seals)
// Nonce should be 8 byte
func EncryptAndSeal(key, nonce, message []byte, add []byte) ([]byte /*encrypted*/, [16]byte /*mac*/, error) {
	var mac [poly1305.TagSize]byte
	if len(key) != 32 {
		return nil, mac, errors.New("invalid key size")
	}
	if len(nonce) != 8 {
		return nil, mac, errors.New("invalid nonce size")
	}

	var (
		Nonce   [12]byte
		aeadOut = make([]byte, len(message)+poly1305.TagSize)
	)
	copy(Nonce[4:], nonce)

	aead, err := chacha20poly1305.New(key)
	if err != nil {
		return nil, mac, err
	}

	aeadOut = aead.Seal(aeadOut[:0], Nonce[:], message, add)
	copy(mac[:], aeadOut[len(message):])
	return aeadOut[:len(message)], mac, nil
}
