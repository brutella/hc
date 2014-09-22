package gohap

import(
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
    "log"
)

type HAPPairSetupClient struct {
    name string
    password string
    publicKey []byte
    secretKey []byte
    srp *srp.SRP
    session *srp.ClientSession
}

func NewHAPPairSetupClient(username string, password string) *HAPPairSetupClient {
    rp, err := srp.NewSRP("openssl.3072", sha512.New, nil)
    rp.SaltLength = 16
    client := rp.NewClientSession([]byte("Pair-Setup"), []byte(password))
    _, _, err = rp.ComputeVerifier([]byte(password))
    if err != nil {
        log.Fatal(err)
    }
    
    LTPK, LTSK, err := ED25519GenerateKey(username)
    
    hap := HAPPairSetupClient{
                name: username, 
                password: password, 
                publicKey: LTPK, 
                secretKey: LTSK, 
                srp: rp, 
                session: client,
            }
            
    return &hap
}

type HAPPairVerifyClient struct {
    name string
    password string
    publicKey []byte
    secretKey []byte
    session *PairVerifySession
}

func NewHAPPairVerifyClient(username string, password string) *HAPPairVerifyClient {
    LTPK, LTSK, _ := ED25519GenerateKey(username)
    
    // Client session keys
    secretKey := Curve25519_GenerateSecretKey()
    publicKey := Curve25519_PublicKey(secretKey)
    session := NewPairVerifySession()
    session.publicKey = publicKey
    session.secretKey = secretKey
    
    hap := HAPPairVerifyClient{
                name: username, 
                password: password, 
                publicKey: LTPK, 
                secretKey: LTSK,
                session: session,
            }
            
    return &hap
}

func (c *HAPPairVerifyClient) GenerateSharedSecret(otherPublicKey []byte) {
    var key [32]byte
    copy(key[:], otherPublicKey)
    c.session.sharedKey = Curve25519_SharedSecret(c.session.secretKey, key)
    
    K, _ := HKDF_SHA512(c.session.sharedKey[:], []byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))
    c.session.encryptionKey = K
}