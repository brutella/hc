package pair

import(
    "github.com/brutella/hap"
)

type HAPPairVerifyClient struct {
    Name string
    Password string
    PublicKey []byte
    SecretKey []byte
    Session *VerifySession
}

func NewHAPPairVerifyClient(username string, password string) *HAPPairVerifyClient {
    LTPK, LTSK, _ := hap.ED25519GenerateKey(username)
    
    // Client session keys
    secretKey := hap.Curve25519_GenerateSecretKey()
    publicKey := hap.Curve25519_PublicKey(secretKey)
    session := NewVerifySession()
    session.publicKey = publicKey
    session.secretKey = secretKey
    
    hap := HAPPairVerifyClient{
                Name: username, 
                Password: password, 
                PublicKey: LTPK, 
                SecretKey: LTSK,
                Session: session,
            }
            
    return &hap
}

func (c *HAPPairVerifyClient) GenerateSharedSecret(otherPublicKey []byte) {
    var key [32]byte
    copy(key[:], otherPublicKey)
    c.Session.sharedKey = hap.Curve25519_SharedSecret(c.Session.secretKey, key)
    
    K, _ := hap.HKDF_SHA512(c.Session.sharedKey[:], []byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))
    c.Session.encryptionKey = K
}