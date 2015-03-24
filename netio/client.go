package netio

import (
	"github.com/brutella/hc/db"
)

// Client is a HomeKit client.
type Client struct {
	entity db.Entity
}

// NewClient returns a client for a specific name either loaded from the database
// or newly created.
func NewClient(name string, database db.Database) (*Client, error) {
	var err error
	entity := database.EntityWithName(name)
	if entity == nil {
		entity, err = db.NewRandomEntityWithName(name)
		if err == nil {
			err = database.SaveEntity(entity)
		}
	}

	return &Client{entity}, err
}

// PairUsername returns the client username used for pairing.
// The returned string is a MAC 48 address which is required by HomeKit.
func (c *Client) PairUsername() string {
	return MAC48Address(c.entity.Name())
}

// PairPrivateKey returns the client private key used for pairing
func (c *Client) PairPrivateKey() []byte {
	return c.entity.PrivateKey()
}

// PairPublicKey returns the client public key used for pairing
func (c *Client) PairPublicKey() []byte {
	return c.entity.PublicKey()
}
