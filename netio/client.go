package netio

import (
	"github.com/brutella/hc/db"
)

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

// PairUsername returns a special formatted string (similar to a MAC address) based on the name.
// HomeKit requires this format.
func (c *Client) PairUsername() string {
	return MAC48Address(c.entity.Name())
}

func (c *Client) PairPrivateKey() []byte {
	return c.entity.PrivateKey()
}

func (c *Client) PairPublicKey() []byte {
	return c.entity.PublicKey()
}
