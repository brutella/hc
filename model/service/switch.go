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

	service := New()
	service.Type = typeSwitch
	service.addCharacteristic(onChar.Characteristic)
	service.addCharacteristic(nameChar.Characteristic)

	return &Switch{service, onChar, nameChar}
}
