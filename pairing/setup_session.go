package gohap

import (
    "github.com/brutella/gohap"
    "crypto/sha512"
    
    // Do not use because of https://github.com/tadglines/go-pkgs/issues/2
    // "github.com/tadglines/go-pkgs/crypto/srp"
    
    // This is actually a fork of the SRP library above
    "github.com/theojulienne/go-srp/crypto/srp"
    
    "errors"
)

type PairSetupSession struct {
    srp *srp.SRP
    session *srp.ServerSession
    salt []byte // s
    publicKey []byte // B
    secretKey []byte // S
    encryptionKey [32]byte // K
}

func NewPairSetupSession(username string, password string) (*PairSetupSession, error) {
    var err error
    
    srp, err := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(username)))
    if err == nil {
        srp.SaltLength = 16
        salt, v, err := srp.ComputeVerifier([]byte(password))
        if err == nil {
            session := srp.NewServerSession([]byte(username), salt, v)
            pairing := PairSetupSession{
                        srp: srp, 
                        session: session, 
                        salt: salt,
                        publicKey: session.GetB(),
                    }
            return &pairing, nil
        }
    }
    
    return nil, err
}

// Validates `M1` from client
func (p *PairSetupSession) ProofFromClientProof(clientProof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(clientProof) { // Validates M1 based on S and A
		return nil, errors.New("Client proof is not valid")
	}
    
	return p.session.ComputeAuthenticator(clientProof), nil
}

// Calculates secret key `S` based on client public key `A`
func (p *PairSetupSession) SetupSecretKeyFromClientPublicKey(key []byte) (error) {
	key, err := p.session.ComputeKey(key) // S
    if err == nil {
        p.secretKey = key
    }
    
    return err
}

// Calculates encryption key `K` based on salt and info
//
// Only 32 bytes are used from HKDF-SHA512
func (p *PairSetupSession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := gohap.HKDF_SHA512(p.secretKey, salt, info)
    if err == nil {
        p.encryptionKey = key
    }
    
    return err
}

