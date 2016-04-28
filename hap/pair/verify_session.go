package pair

import (
	"github.com/brutella/hc/crypto/curve25519"
	"github.com/brutella/hc/crypto/hkdf"
)

// VerifySession holds keys to encrypt a tcp connection.
type VerifySession struct {
	OtherPublicKey [32]byte
	PublicKey      [32]byte
	PrivateKey     [32]byte
	SharedKey      [32]byte
	EncryptionKey  [32]byte
}

// NewVerifySession creates a new session with random public and private key
func NewVerifySession() *VerifySession {
	privateKey := curve25519.GeneratePrivateKey()
	publicKey := curve25519.PublicKey(privateKey)

	return &VerifySession{
		PublicKey:  publicKey,
		PrivateKey: privateKey,
	}
}

// GenerateSharedKeyWithOtherPublicKey generates a Curve25519 shared key based on a public key.
// The other public key is also stored for further use in `otherPublicKey` property.
func (s *VerifySession) GenerateSharedKeyWithOtherPublicKey(otherPublicKey [32]byte) {
	sharedKey := curve25519.SharedSecret(s.PrivateKey, otherPublicKey)

	s.OtherPublicKey = otherPublicKey
	s.SharedKey = sharedKey
}

// SetupEncryptionKey generates an encryption key based on the shared key, salt and info.
func (s *VerifySession) SetupEncryptionKey(salt []byte, info []byte) error {
	hash, err := hkdf.Sha512(s.SharedKey[:], salt, info)
	if err == nil {
		s.EncryptionKey = hash
	}

	return err
}
