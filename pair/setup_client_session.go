package pair

import (
    "github.com/brutella/hap"
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
)

type SetupClientSession struct {
    srp *srp.SRP
    session *srp.ClientSession
    publicKey []byte // A
    secretKey []byte // S
    proof []byte // M1
    Name []byte
    LTPK []byte
    LTSK []byte
    encryptionKey [32]byte // K
}

func NewSetupClientSession(username string, password string) (*SetupClientSession) {
    rp, _ := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(username)))
    
    client := rp.NewClientSession([]byte(username), []byte(password))
    LTPK, LTSK, _ := hap.ED25519GenerateKey(username)
    
    hap := SetupClientSession{
                Name: []byte(username),
                LTPK: LTPK,
                LTSK: LTSK,
                srp: rp,
                session: client,
            }
            
    return &hap
}

func (s *SetupClientSession) GenerateKeys(salt []byte, otherPublicKey []byte) error {
    secretKey, err := s.session.ComputeKey(salt, otherPublicKey)
    if err == nil {
        s.publicKey = s.session.GetA()
        s.secretKey = secretKey
        s.proof = s.session.ComputeAuthenticator()
    }
    
    return err
}

// Validates `M2` from server
func (s *SetupClientSession) IsServerProofValid(proof []byte) bool {
    return s.session.VerifyServerAuthenticator(proof)
}

// Calculates encryption key `K` based on salt and info
//
// Only 32 bytes are used from HKDF-SHA512
func (p *SetupClientSession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := hap.HKDF_SHA512(p.secretKey, salt, info)
    if err == nil {
        p.encryptionKey = key
    }
    
    return err
}

