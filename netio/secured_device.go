package netio

import (
	"github.com/brutella/hc/db"
)

// SecuredDevice is a HomeKit device with a password.
type SecuredDevice interface {
	Device
	Password() string
}

type securedDevice struct {
	Device
	password string
}

// NewSecuredDevice returns a device for a specific name either loaded from the database or newly created.
// Additionally other device can only pair with by providing the correct password.
func NewSecuredDevice(name string, password string, database db.Database) (SecuredDevice, error) {
	d, err := NewDevice(name, database)
	return &securedDevice{d, password}, err
}

// Password returns the device password.
func (d *securedDevice) Password() string {
	return d.password
}
