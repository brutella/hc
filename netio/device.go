package netio

import (
	"github.com/brutella/hc/db"
)

type Device interface {
	// PairUsername returns the username used for pairing
	PairUsername() string

	// PairPrivateKey returns the client private key used for pairing
	PairPrivateKey() []byte

	// PairPublicKey returns the client public key used for pairing
	PairPublicKey() []byte
}

type device struct {
	entity db.Entity
}

// NewDevice returns a client for a specific name either loaded from the database
// or newly created.
func NewDevice(name string, database db.Database) (Device, error) {
	var err error
	entity := database.EntityWithName(name)
	if entity == nil {
		entity, err = db.NewRandomEntityWithName(name)
		if err == nil {
			err = database.SaveEntity(entity)
		}
	}

	return &device{entity}, err
}

func (d *device) PairUsername() string {
	return d.entity.Name()
}

// PairPrivateKey returns the client private key used for pairing
func (d *device) PairPrivateKey() []byte {
	return d.entity.PrivateKey()
}

// PairPublicKey returns the client public key used for pairing
func (d *device) PairPublicKey() []byte {
	return d.entity.PublicKey()
}
