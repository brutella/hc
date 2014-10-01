package pair

import(
    "github.com/brutella/hap/crypto"
)

type VerifySession struct {
    otherPublicKey [32]byte
    publicKey [32]byte
    secretKey [32]byte
    sharedKey [32]byte
    encryptionKey [32]byte
}

func NewVerifySession() (*VerifySession) {
    secretKey := crypto.Curve25519_GenerateSecretKey()
    publicKey := crypto.Curve25519_PublicKey(secretKey)
    
    return &VerifySession{
        publicKey: publicKey,
        secretKey: secretKey,
    }
}

// Generate Curve25519 shared key for a specified other public key
// The other public key is also stored for further use in `otherPublicKey` property
func (s *VerifySession) GenerateSharedKeyWithOtherPublicKey(otherPublicKey [32]byte) {    
    sharedKey := crypto.Curve25519_SharedSecret(s.secretKey, otherPublicKey)
    
    s.otherPublicKey = otherPublicKey
    s.sharedKey = sharedKey
}

// Generates encryption based on shared key, salt and info
func (s *VerifySession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := crypto.HKDF_SHA512(s.sharedKey[:], salt, info)
    if err == nil {
        s.encryptionKey = key
    }
    
    return err
}