package pair

import (
	"github.com/brutella/hc/crypto/hkdf"

	"crypto/sha512"
	"github.com/tadglines/go-pkgs/crypto/srp"
)

// SetupClientSession holds the keys to pair with an accessory.
type SetupClientSession struct {
	session       *srp.ClientSession
	PublicKey     []byte   // A
	PrivateKey    []byte   // S
	Proof         []byte   // M1
	EncryptionKey [32]byte // K
}

// NewSetupClientSession returns a new setup client session
func NewSetupClientSession(username string, password string) *SetupClientSession {
	rp, _ := srp.NewSRP(SRPGroup, sha512.New, KeyDerivativeFuncRFC2945(sha512.New, []byte(username)))

	client := rp.NewClientSession([]byte(username), []byte(password))
	hap := SetupClientSession{
		session: client,
	}

	return &hap
}

// GenerateKeys generates public and private keys based on server's salt and public key.
func (s *SetupClientSession) GenerateKeys(salt []byte, otherPublicKey []byte) error {
	privateKey, err := s.session.ComputeKey(salt, otherPublicKey)
	if err == nil {
		s.PublicKey = s.session.GetA()
		s.PrivateKey = privateKey
		s.Proof = s.session.ComputeAuthenticator()
	}

	return err
}

// IsServerProofValid returns true when the server proof `M2` is valid.
func (s *SetupClientSession) IsServerProofValid(proof []byte) bool {
	return s.session.VerifyServerAuthenticator(proof)
}

// SetupEncryptionKey calculates encryption key `K` based on salt and info.
func (s *SetupClientSession) SetupEncryptionKey(salt []byte, info []byte) error {
	hash, err := hkdf.Sha512(s.PrivateKey, salt, info)
	if err == nil {
		s.EncryptionKey = hash
	}

	return err
}
