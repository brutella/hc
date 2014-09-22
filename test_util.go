package gohap

import(
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
    "log"
)

type HAPClient struct {
    name string
    password string
    publicKey []byte
    secretKey []byte
    srp *srp.SRP
    session *srp.ClientSession
}

func NewHAPClient(username string, password string) *HAPClient {
    rp, err := srp.NewSRP("openssl.3072", sha512.New, nil)
    client := rp.NewClientSession([]byte("Pair-Setup"), []byte(password))
    _, _, err = rp.ComputeVerifier([]byte(password))
    if err != nil {
        log.Fatal(err)
    }
    
    LTPK, LTSK, err := ED25519GenerateKey(username)
    
    hap := HAPClient{
                name: username, 
                password: password, 
                publicKey: LTPK, 
                secretKey: LTSK, 
                srp: rp, 
                session: client,
            }
            
    return &hap
}