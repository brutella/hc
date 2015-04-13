package netio

import (
	"github.com/brutella/hc/db"
)

// Device is a HomeKit device with a name, private and public key.
type Device interface {
	// Name returns the username used for pairing
	Name() string

	// PrivateKey returns the client private key used for pairing
	PrivateKey() []byte

	// PublicKey returns the client public key used for pairing
	PublicKey() []byte
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

func (d *device) Name() string {
	return d.entity.Name()
}

// PairPrivateKey returns the client private key used for pairing
func (d *device) PrivateKey() []byte {
	return d.entity.PrivateKey()
}

// PairPublicKey returns the client public key used for pairing
func (d *device) PublicKey() []byte {
	return d.entity.PublicKey()
}
