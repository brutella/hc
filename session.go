package gohap

import (
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
    "errors"
)

type PairingSession struct {
    srp *srp.SRP
    session *srp.ServerSession
    salt []byte
    publicKey []byte
    secretKey []byte
}

func NewPairingSession(username string, password string) (*PairingSession, error) {
    var err error
    
	srp, err := srp.NewSRP("openssl.3072", sha512.New, nil)
	if err == nil {
        srp.SaltLength = 16
    	salt, v, err := srp.ComputeVerifier([]byte(password))
    	if err == nil {
        	session := srp.NewServerSession([]byte(username), salt, v)
            pairing := PairingSession{
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

func (p *PairingSession) PublicKey() []byte {
    return p.publicKey
}

func (p *PairingSession) Salt() []byte {
    return p.salt
}

func (p *PairingSession) ProofFromClientProof(proof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(proof) { // Validates M1 based on S and A
		return nil, errors.New("Client proof is not valid")
	}
    
	return p.session.ComputeAuthenticator(proof), nil
}

func (p *PairingSession) SecretKeyFromClientPublicKey(key []byte) ([]byte, error) {
	skey, err := p.session.ComputeKey(key) // S
    p.secretKey = skey
    
    return skey, err
}

