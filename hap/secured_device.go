package hap

import (
	"github.com/brutella/hc/db"
)

// SecuredDevice is a HomeKit device with a pin.
type SecuredDevice interface {
	Device
	Pin() string
}

type securedDevice struct {
	Device
	pin string
}

// NewSecuredDevice returns a device for a specific name either loaded from the database or newly created.
// Additionally other device can only pair with by providing the correct pin.
func NewSecuredDevice(name string, pin string, database db.Database) (SecuredDevice, error) {
	d, err := NewDevice(name, database)
	return &securedDevice{d, pin}, err
}

// Pin returns the device pin.
func (d *securedDevice) Pin() string {
	return d.pin
}
