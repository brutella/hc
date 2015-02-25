package netio

import (
	"github.com/brutella/hc/crypto"
)

// Bridge contains basic information (like name, ...) and cryptho keys to secure
// the communication.
type Bridge struct {
	info BridgeInfo

	PublicKey []byte
	SecretKey []byte
}

// NewBridge returns a bridge from a BridgeInfo object
//
// The long-term public and secret key are based on the serial
// number which should be unique for every bridge.
func NewBridge(info BridgeInfo) (*Bridge, error) {
	b := Bridge{info: info}
	public, secret, err := crypto.ED25519GenerateKey(b.info.SerialNumber)
	b.PublicKey = public
	b.SecretKey = secret

	return &b, err
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
