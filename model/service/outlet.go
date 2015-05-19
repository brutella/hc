package service

import (
	"github.com/brutella/hc/model/characteristic"
)

// Outlet is a service representing an outlet.
type Outlet struct {
	*Switch
	InUse *characteristic.InUse
}

// NewOutlet returns a outlet service.
func NewOutlet(name string, on, inUse bool) *Outlet {
	inUseChar := characteristic.NewInUse(inUse)

	sw := NewSwitch(name, on)
	sw.Type = typeOutlet
	sw.addCharacteristic(inUseChar.Characteristic)

	return &Outlet{sw, inUseChar}
}
