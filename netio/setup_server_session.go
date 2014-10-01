package netio

import (
    "github.com/brutella/hap/crypto"
    
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
    
    "errors"
)

type PairSetupServerSession struct {
    srp *srp.SRP
    session *srp.ServerSession
    Salt []byte // s
    PublicKey []byte // B
    SecretKey []byte // S
    EncryptionKey [32]byte // K
}

func NewPairSetupServerSession(username string, password string) (*PairSetupServerSession, error) {
    var err error
    
    srp, err := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(username)))
    if err == nil {
        srp.SaltLength = 16
        salt, v, err := srp.ComputeVerifier([]byte(password))
        if err == nil {
            session := srp.NewServerSession([]byte(username), salt, v)
            pairing := PairSetupServerSession{
                        srp: srp, 
                        session: session, 
                        Salt: salt,
                        PublicKey: session.GetB(),
                    }
            return &pairing, nil
        }
    }
    
    return nil, err
}

// Validates `M1` from client
func (p *PairSetupServerSession) ProofFromClientProof(clientProof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(clientProof) { // Validates M1 based on S and A
		return nil, errors.New("Client proof is not valid")
	}
    
	return p.session.ComputeAuthenticator(clientProof), nil
}

// Calculates secret key `S` based on client public key `A`
func (p *PairSetupServerSession) SetupSecretKeyFromClientPublicKey(key []byte) (error) {
	key, err := p.session.ComputeKey(key) // S
    if err == nil {
        p.SecretKey = key
    }
    
    return err
}

// Calculates encryption key `K` based on salt and info
//
// Only 32 bytes are used from HKDF-SHA512
func (p *PairSetupServerSession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := crypto.HKDF_SHA512(p.SecretKey, salt, info)
    if err == nil {
        p.EncryptionKey = key
    }
    
    return err
}

