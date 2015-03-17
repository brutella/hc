package pair

import (
	"github.com/brutella/hc/crypto"

	"crypto/sha512"
	"github.com/tadglines/go-pkgs/crypto/srp"
)

type PairSetupClientSession struct {
	srp           *srp.SRP
	session       *srp.ClientSession
	PublicKey     []byte   // A
	SecretKey     []byte   // S
	Proof         []byte   // M1
	EncryptionKey [32]byte // K
}

func NewPairSetupClientSession(username string, password string) *PairSetupClientSession {
	rp, _ := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(username)))

	client := rp.NewClientSession([]byte(username), []byte(password))
	hap := PairSetupClientSession{
		srp:     rp,
		session: client,
	}

	return &hap
}

func (s *PairSetupClientSession) GenerateKeys(salt []byte, otherPublicKey []byte) error {
	secretKey, err := s.session.ComputeKey(salt, otherPublicKey)
	if err == nil {
		s.PublicKey = s.session.GetA()
		s.SecretKey = secretKey
		s.Proof = s.session.ComputeAuthenticator()
	}

	return err
}

// Validates `M2` from server
func (s *PairSetupClientSession) IsServerProofValid(proof []byte) bool {
	return s.session.VerifyServerAuthenticator(proof)
}

// Calculates encryption key `K` based on salt and info
//
// Only 32 bytes are used from HKDF-SHA512
func (p *PairSetupClientSession) SetupEncryptionKey(salt []byte, info []byte) error {
	key, err := crypto.HKDF_SHA512(p.SecretKey, salt, info)
	if err == nil {
		p.EncryptionKey = key
	}

	return err
}
