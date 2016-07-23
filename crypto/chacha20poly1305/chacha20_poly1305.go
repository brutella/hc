package chacha20poly1305

import (
	"errors"

	"github.com/aead/chacha20"
)

// DecryptAndVerify returns the chacha20 decrypted messages.
// An error is returned when the poly1305 message authenticator (seal) could not be verified.
// Nonce should be 8 byte.
func DecryptAndVerify(key, nonce, message []byte, mac [16]byte, add []byte) ([]byte, error) {
	if len(key) != 32 {
		return nil, errors.New("invalid key size")
	}
	if len(nonce) != 8 {
		return nil, errors.New("invalid key size")
	}

	var (
		Key         [32]byte
		Nonce       [12]byte
		chacha20Out = make([]byte, len(message))
	)
	copy(Key[:], key)
	copy(Nonce[4:], nonce)

	aead := chacha20.NewChaCha20Poly1305(&Key)
	return aead.Open(chacha20Out[:0], Nonce[:], append(message, mac[:]...), add)
}

// EncryptAndSeal returns the chacha20 encrypted message and poly1305 message authentictor (also refered as seals)
// Nonce should be 8 byte
func EncryptAndSeal(key, nonce, message []byte, add []byte) ([]byte /*encrypted*/, [16]byte /*mac*/, error) {
	var mac [chacha20.TagSize]byte
	if len(key) != 32 {
		return nil, mac, errors.New("invalid key size")
	}
	if len(nonce) != 8 {
		return nil, mac, errors.New("invalid key size")
	}

	var (
		Key         [32]byte
		Nonce       [12]byte
		chacha20Out = make([]byte, len(message)+chacha20.TagSize)
	)
	copy(Key[:], key)
	copy(Nonce[4:], nonce)

	aead := chacha20.NewChaCha20Poly1305(&Key)
	chacha20Out = aead.Seal(chacha20Out[:0], Nonce[:], message, add)

	copy(mac[:], chacha20Out[len(message):])
	return chacha20Out[:len(message)], mac, nil
}
