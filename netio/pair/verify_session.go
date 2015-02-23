package pair

import (
	"github.com/brutella/hap/crypto"
)

type PairVerifySession struct {
	OtherPublicKey [32]byte
	PublicKey      [32]byte
	SecretKey      [32]byte
	SharedKey      [32]byte
	EncryptionKey  [32]byte
}

func NewPairVerifySession() *PairVerifySession {
	secretKey := crypto.Curve25519_GenerateSecretKey()
	publicKey := crypto.Curve25519_PublicKey(secretKey)

	return &PairVerifySession{
		PublicKey: publicKey,
		SecretKey: secretKey,
	}
}

// Generate Curve25519 shared key for a specified other public key
// The other public key is also stored for further use in `otherPublicKey` property
func (s *PairVerifySession) GenerateSharedKeyWithOtherPublicKey(otherPublicKey [32]byte) {
	sharedKey := crypto.Curve25519_SharedSecret(s.SecretKey, otherPublicKey)

	s.OtherPublicKey = otherPublicKey
	s.SharedKey = sharedKey
}

// Generates encryption based on shared key, salt and info
func (s *PairVerifySession) SetupEncryptionKey(salt []byte, info []byte) error {
	key, err := crypto.HKDF_SHA512(s.SharedKey[:], salt, info)
	if err == nil {
		s.EncryptionKey = key
	}

	return err
}
