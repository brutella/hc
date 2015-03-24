package netio

import (
	"github.com/brutella/hc/db"
)

// Bridge contains basic information (like name, ...) and cryptho keys to secure
// the communication.
type Bridge struct {
	info   BridgeInfo
	entity db.Entity
}

// NewBridge returns a bridge from a BridgeInfo object
//
// The long-term public and secret key are based on the serial
// number which should be unique for every bridge.
func NewBridge(info BridgeInfo, database db.Database) (*Bridge, error) {
	var err error
	entity := database.EntityWithName(info.ID)
	if entity == nil {
		entity, err = db.NewRandomEntityWithName(info.ID)
		if err == nil {
			err = database.SaveEntity(entity)
		}
	}

	return &Bridge{info, entity}, err
}

// Name returns the bridge name
func (b *Bridge) Name() string {
	return b.info.Name
}

// ID returns the bridge id which is used as username for pairing.
func (b *Bridge) ID() string {
	return b.info.ID
}

// Password returns the bridge password
func (b *Bridge) Password() string {
	return b.info.Password
}

// PairUsername returns the username used for pairing, which is actually the return value of ID().
func (b *Bridge) PairUsername() string {
	return b.ID()
}

// PairPrivateKey returns the private key used for pairing.
func (b *Bridge) PairPrivateKey() []byte {
	return b.entity.PrivateKey()
}

// PairPublicKey returns the bridge public key used for pairing.
func (b *Bridge) PairPublicKey() []byte {
	return b.entity.PublicKey()
}
