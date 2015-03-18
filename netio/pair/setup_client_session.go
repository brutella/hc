package pair

import (
	"github.com/brutella/hc/crypto"

	"crypto/sha512"
	"github.com/tadglines/go-pkgs/crypto/srp"
)

type SetupClientSession struct {
	session       *srp.ClientSession
	PublicKey     []byte   // A
	PrivateKey    []byte   // S
	Proof         []byte   // M1
	EncryptionKey [32]byte // K
}

func NewSetupClientSession(username string, password string) *SetupClientSession {
	rp, _ := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(username)))

	client := rp.NewClientSession([]byte(username), []byte(password))
	hap := SetupClientSession{
		session: client,
	}

	return &hap
}

func (s *SetupClientSession) GenerateKeys(salt []byte, otherPublicKey []byte) error {
	privateKey, err := s.session.ComputeKey(salt, otherPublicKey)
	if err == nil {
		s.PublicKey = s.session.GetA()
		s.PrivateKey = privateKey
		s.Proof = s.session.ComputeAuthenticator()
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
func (p *SetupClientSession) SetupEncryptionKey(salt []byte, info []byte) error {
	key, err := crypto.HKDF_SHA512(p.PrivateKey, salt, info)
	if err == nil {
		p.EncryptionKey = key
	}

	return err
}
