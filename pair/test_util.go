package pair

import(
    "github.com/brutella/hap"
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
    //"github.com/theojulienne/go-srp/crypto/srp"
)

type HAPPairSetupClient struct {
    Name string
    Password string
    PublicKey []byte
    SecretKey []byte
    srp *srp.SRP
    Session *srp.ClientSession
}

func NewHAPPairSetupClient(username string, password string) *HAPPairSetupClient {
    srp_username := []byte("Pair-Setup")
    rp, _ := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, srp_username))
    
    client := rp.NewClientSession(srp_username, []byte(password))
    LTPK, LTSK, _ := hap.ED25519GenerateKey(username)
    
    hap := HAPPairSetupClient{
                Name: username, 
                Password: password, 
                PublicKey: LTPK, 
                SecretKey: LTSK, 
                srp: rp, 
                Session: client,
            }
            
    return &hap
}

type HAPPairVerifyClient struct {
    Name string
    Password string
    PublicKey []byte
    SecretKey []byte
    Session *PairVerifySession
}

func NewHAPPairVerifyClient(username string, password string) *HAPPairVerifyClient {
    LTPK, LTSK, _ := hap.ED25519GenerateKey(username)
    
    // Client session keys
    secretKey := hap.Curve25519_GenerateSecretKey()
    publicKey := hap.Curve25519_PublicKey(secretKey)
    session := NewPairVerifySession()
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