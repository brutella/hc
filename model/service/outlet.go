package service

import (
	"github.com/brutella/hap/model/characteristic"
)

type Outlet struct {
	*Switch
	InUse *characteristic.InUse
}

// NewOutlet returns a outlet service.
func NewOutlet(name string, on, inUse bool) *Outlet {
	in_use := characteristic.NewInUse(on)

	sw := NewSwitch(name, on)
	sw.Type = TypeOutlet
	sw.AddCharacteristic(in_use.Characteristic)

	return &Outlet{sw, in_use}
}
