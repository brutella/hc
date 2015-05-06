package service

import (
	"github.com/brutella/hc/model/characteristic"
)

// Switch is a service to represent a switch.
type Switch struct {
	*Service
	On   *characteristic.On
	Name *characteristic.Name
}

// NewSwitch returns a switch service.
func NewSwitch(name string, on bool) *Switch {
	onChar := characteristic.NewOn(on)
	nameChar := characteristic.NewName(name)

	svc := New()
	svc.Type = typeSwitch
	svc.addCharacteristic(onChar.Characteristic)
	svc.addCharacteristic(nameChar.Characteristic)

	return &Switch{svc, onChar, nameChar}
}
