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
	entity := database.EntityWithName(info.Id)
	if entity == nil {
		entity, err = db.NewRandomEntityWithName(info.Id)
	}

	return &Bridge{info, entity}, err
}

// Name returns the bridge name
func (b *Bridge) Name() string {
	return b.info.Name
}

// Id returns the bridge id which is used as username for pairing.
func (b *Bridge) Id() string {
	return b.info.Id
}

// Password returns the bridge password
func (b *Bridge) Password() string {
	return b.info.Password
}

func (b *Bridge) PairUsername() string {
	return b.Id()
}

// PrivateKey returns the bridge private key
func (b *Bridge) PairPrivateKey() []byte {
	return b.entity.PrivateKey()
}

// PublicKey returns the bridge public key
func (b *Bridge) PairPublicKey() []byte {
	return b.entity.PublicKey()
}
