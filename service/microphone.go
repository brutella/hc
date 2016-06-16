// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeMicrophone = "112"

type Microphone struct {
	*Service

	Mute *characteristic.Mute
}

func NewMicrophone() *Microphone {
	svc := Microphone{}
	svc.Service = New(TypeMicrophone)

	svc.Mute = characteristic.NewMute()
	svc.AddCharacteristic(svc.Mute.Characteristic)

	return &svc
}
