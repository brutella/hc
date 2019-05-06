// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSpeaker = "113"

type Speaker struct {
	*Service

	Mute *characteristic.Mute

	Name   *characteristic.Name
	Volume *characteristic.Volume
}

func NewSpeaker() *Speaker {
	svc := Speaker{}
	svc.Service = New(TypeSpeaker)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	return &svc
}

func (svc *Speaker) addOptionalCharaterics() {

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

	svc.Volume = characteristic.NewVolume()
	svc.AddCharacteristic(svc.Volume.Characteristic)

}
