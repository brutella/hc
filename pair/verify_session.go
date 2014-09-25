package pair

import(
    "github.com/brutella/hap"
)

type VerifySession struct {
    otherPublicKey [32]byte
    publicKey [32]byte
    secretKey [32]byte
    sharedKey [32]byte
    encryptionKey [32]byte
}

func NewVerifySession() (*VerifySession) {
    secretKey := hap.Curve25519_GenerateSecretKey()
    publicKey := hap.Curve25519_PublicKey(secretKey)
    
    return &VerifySession{
        publicKey: publicKey,
        secretKey: secretKey,
    }
}

// Generate Curve25519 shared key for a specified other public key
// The other public key is also stored for further use in `otherPublicKey` property
func (s *VerifySession) GenerateSharedKeyWithOtherPublicKey(otherPublicKey [32]byte) {    
    sharedKey := hap.Curve25519_SharedSecret(s.secretKey, otherPublicKey)
    
    s.otherPublicKey = otherPublicKey
    s.sharedKey = sharedKey
}

// Generates encryption based on shared key, salt and info
func (s *VerifySession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := hap.HKDF_SHA512(s.sharedKey[:], salt, info)
    if err == nil {
        s.encryptionKey = key
    }
    
    return err
}