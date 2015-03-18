package pair

import (
	"github.com/brutella/hc/crypto"

	"crypto/sha512"
	"github.com/tadglines/go-pkgs/crypto/srp"

	"errors"
)

type SetupServerSession struct {
	session       *srp.ServerSession
	Salt          []byte   // s
	PublicKey     []byte   // B
	PrivateKey    []byte   // S
	EncryptionKey [32]byte // K
	Username      []byte
}

func NewSetupServerSession(username, password string) (*SetupServerSession, error) {
	var err error
	pair_name := []byte("Pair-Setup")
	srp, err := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(pair_name)))

	if err == nil {
		srp.SaltLength = 16
		salt, v, err := srp.ComputeVerifier([]byte(password))
		if err == nil {
			session := srp.NewServerSession([]byte(pair_name), salt, v)
			pairing := SetupServerSession{
				session:   session,
				Salt:      salt,
				PublicKey: session.GetB(),
				Username:  []byte(username),
			}
			return &pairing, nil
		}
	}

	return nil, err
}

// ProofFromClientProof validates client proof (`M1`) and returns authenticator or error if proof is not valid.
func (p *SetupServerSession) ProofFromClientProof(clientProof []byte) ([]byte, error) {
	if !p.session.VerifyClientAuthenticator(clientProof) { // Validates M1 based on S and A
		return nil, errors.New("Client proof is not valid")
	}

	return p.session.ComputeAuthenticator(clientProof), nil
}

// SetupPrivateKeyFromClientPublicKey calculates and internally sets secret key `S` based on client public key `A`
func (p *SetupServerSession) SetupPrivateKeyFromClientPublicKey(key []byte) error {
	key, err := p.session.ComputeKey(key) // S
	if err == nil {
		p.PrivateKey = key
	}

	return err
}

// SetupEncryptionKey calculates and internally sets encryption key `K` based on salt and info
//
// Only 32 bytes are used from HKDF-SHA512
func (p *SetupServerSession) SetupEncryptionKey(salt []byte, info []byte) error {
	key, err := crypto.HKDF_SHA512(p.PrivateKey, salt, info)
	if err == nil {
		p.EncryptionKey = key
	}

	return err
}
