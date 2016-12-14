// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSlat = "B9"

type Slat struct {
	*Service

	SlatType         *characteristic.SlatType
	CurrentSlatState *characteristic.CurrentSlatState
}

func NewSlat() *Slat {
	svc := Slat{}
	svc.Service = New(TypeSlat)

	svc.SlatType = characteristic.NewSlatType()
	svc.AddCharacteristic(svc.SlatType.Characteristic)

	svc.CurrentSlatState = characteristic.NewCurrentSlatState()
	svc.AddCharacteristic(svc.CurrentSlatState.Characteristic)

	return &svc
}
