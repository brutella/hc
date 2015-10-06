package db

import (
	"github.com/brutella/hc/crypto"
	"github.com/brutella/hc/util"
)

type Entity struct {
	Name       string
	PublicKey  []byte
	PrivateKey []byte
}

// NewRandomEntityWithName returns an entity with a random private and public keys
func NewRandomEntityWithName(name string) (e Entity, err error) {
	var public []byte
	var private []byte

	public, private, err = generateKeyPairs()
	if err == nil && len(public) > 0 && len(private) > 0 {
		e = NewEntity(name, public, private)
	}

	return
}

// NewEntity returns a entity with a name, public and private key.
func NewEntity(name string, publicKey, privateKey []byte) Entity {
	return Entity{Name: name, PublicKey: publicKey, PrivateKey: privateKey}
}

// generateKeyPairs generates random public and private key pairs
func generateKeyPairs() ([]byte, []byte, error) {
	str := util.RandomHexString()
	public, private, err := crypto.ED25519GenerateKey(str)
	return public, private, err
}
