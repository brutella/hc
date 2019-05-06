// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeMicrophone = "112"

type Microphone struct {
	*Service

	Volume *characteristic.Volume
	Mute   *characteristic.Mute

	Name *characteristic.Name
}

func NewMicrophone() *Microphone {
	svc := Microphone{}
	svc.Service = New(TypeMicrophone)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	return &svc
}

func (svc *Microphone) AddOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
