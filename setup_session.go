package gohap

import (
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
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
    
    srp, err := srp.NewSRP("rfc5054.3072", sha512.New, nil)
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

func (p *PairSetupSession) PublicKey() []byte {
    return p.publicKey
}

func (p *PairSetupSession) Salt() []byte {
    return p.salt
}

func (p *PairSetupSession) ProofFromClientProof(clientProof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(clientProof) { // Validates M1 based on S and A
		return nil, errors.New("Client proof is not valid")
	}
    
	return p.session.ComputeAuthenticator(clientProof), nil
}

func (p *PairSetupSession) SetupSecretKeyFromClientPublicKey(key []byte) (error) {
	key, err := p.session.ComputeKey(key) // S
    if err == nil {
        p.secretKey = key
    }
    
    return err
}

func (p *PairSetupSession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := HKDF_SHA512(p.secretKey, salt, info)
    if err == nil {
        p.encryptionKey = key
    }
    
    return err
}

