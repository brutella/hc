package netio

import (
	"github.com/brutella/hc/db"
)

type SecuredDevice interface {
	Device
	Password() string
}

type securedDevice struct {
	Device
	password string
}

// NewDevice returns a client for a specific name either loaded from the database
// or newly created.
func NewSecuredDevice(name string, password string, database db.Database) (SecuredDevice, error) {
	d, err := NewDevice(name, database)
	return &securedDevice{d, password}, err
}

func (d *securedDevice) Password() string {
	return d.password
}
