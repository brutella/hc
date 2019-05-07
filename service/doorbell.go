// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeDoorbell = "121"

type Doorbell struct {
	*Service

	ProgrammableSwitchEvent *characteristic.ProgrammableSwitchEvent

	Brightness *characteristic.Brightness
	Volume     *characteristic.Volume
	Name       *characteristic.Name
}

func NewDoorbell() *Doorbell {
	svc := Doorbell{}
	svc.Service = New(TypeDoorbell)

	svc.ProgrammableSwitchEvent = characteristic.NewProgrammableSwitchEvent()
	svc.AddCharacteristic(svc.ProgrammableSwitchEvent.Characteristic)

	return &svc
}

func (svc *Doorbell) AddOptionalCharacteristics() {

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
