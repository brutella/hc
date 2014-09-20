package gohap

import (
    "crypto/sha512"
    "github.com/tadglines/go-pkgs/crypto/srp"
    "errors"
)

type PairingSession struct {
    srp *srp.SRP
    session *srp.ServerSession
    salt []byte // s
    publicKey []byte // B
    secretKey []byte // S
    encryptionKey []byte // K
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

func (p *PairingSession) ProofFromClientProof(clientProof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(clientProof) { // Validates M1 based on S and A
		return nil, errors.New("Client proof is not valid")
	}
    
	return p.session.ComputeAuthenticator(clientProof), nil
}

func (p *PairingSession) SetupSecretKeyFromClientPublicKey(key []byte) (error) {
	key, err := p.session.ComputeKey(key) // S
    if err == nil {
        p.secretKey = key
    }
    
    return err
}

func (p *PairingSession) SetupEncryptionKey(salt []byte, info []byte) (error) {
    key, err := HKDF_SHA512_256(p.secretKey, salt, info)
    if err == nil {
        p.encryptionKey = key
    }
    
    return err
}

