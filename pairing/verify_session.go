package gohap

import(
    "github.com/brutella/gohap"
)

type PairVerifySession struct {
    otherPublicKey [32]byte
    publicKey [32]byte
    secretKey [32]byte
    sharedKey [32]byte
    encryptionKey [32]byte
}

func NewPairVerifySession() (*PairVerifySession) {
    return &PairVerifySession{}
}

// Generate Curve25519 public, secret and shared key for a specified other public key
// The other public key is also stored for further use in `otherPublicKey` property
func (s *PairVerifySession) GenerateKeysWithOtherPublicKey(otherPublicKey [32]byte) {
    secretKey := gohap.Curve25519_GenerateSecretKey()
    publicKey := gohap.Curve25519_PublicKey(secretKey)
    sharedKey := gohap.Curve25519_SharedSecret(secretKey, otherPublicKey)
    
    s.otherPublicKey = otherPublicKey
    s.secretKey = secretKey
    s.publicKey = publicKey
    s.sharedKey = sharedKey
}

func (s *PairVerifySession) PublicKey() []byte {
    return s.publicKey[:]
}

func (s *PairVerifySession) EncryptionKey() []byte {
    return s.encryptionKey[:]
}