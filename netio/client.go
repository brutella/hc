package netio

import (
	"github.com/brutella/hc/db"
)

type Client struct {
	entity db.Entity
}

func NewClient(name string, database db.Database) (*Client, error) {
	var err error
	entity := database.EntityWithName(name)
	if entity == nil {
		entity, err = db.NewRandomEntityWithName(name)
	}

	return &Client{entity}, err
}

// Id returns a special formatted string (similar to a MAC address) based on the name.
// HomeKit requires this format.
func (c *Client) Id() string {
	return MAC48Address(c.entity.Name())
}

func (c *Client) PrivateKey() []byte {
	return c.entity.PrivateKey()
}

func (c *Client) PublicKey() []byte {
	return c.entity.PublicKey()
}
