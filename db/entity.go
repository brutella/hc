package db

import (
	"github.com/brutella/hc/util"
	"github.com/brutella/hc/crypto"
)

// Entity is a HomeKit entity (e.g. iOS device or HomeKit bridge).
type Entity interface {
	// Name returns the entity name
	Name() string

	// SetName sets the entity name
	SetName(name string)

	// PublicKey returns the entity (long-term) public key
	PublicKey() []byte

	// SetPublicKey sets the entity (long-term) public key
	SetPublicKey(publicKey []byte)

	// PrivateKey returns the entity (long-term) private key
	PrivateKey() []byte

	// SetPrivateKey sets the entity (long-term) private key
	SetPrivateKey(privateKey []byte)
}

type entity struct {
	name       string
	publicKey  []byte
	privateKey []byte
}

// NewRandomEntityWithName returns an entity with a random private and public keys
func NewRandomEntityWithName(name string) (Entity, error) {
	public, private, err := generateKeyPairs()
	if err == nil && len(public) > 0 && len(private) > 0 {
		return NewEntity(name, public, private), nil
	}

	return nil, err
}

// NewEntity returns a entity with a name and public key.
func NewEntity(name string, publicKey, privateKey []byte) Entity {
	return &entity{name: name, publicKey: publicKey, privateKey: privateKey}
}

func (c *entity) SetName(name string) {
	c.name = name
}

func (c *entity) Name() string {
	return c.name
}

func (c *entity) SetPublicKey(publicKey []byte) {
	c.publicKey = publicKey
}

func (c *entity) PublicKey() []byte {
	return c.publicKey
}

func (c *entity) SetPrivateKey(privateKey []byte) {
	c.privateKey = privateKey
}

func (c *entity) PrivateKey() []byte {
	return c.privateKey
}

// generateKeyPairs generates random public and private key pairs
func generateKeyPairs() ([]byte, []byte, error) {
	str := util.RandomHexString()
	public, private, err := crypto.ED25519GenerateKey(str)
	return public, private, err
}
