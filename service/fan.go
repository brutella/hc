// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeFan = "00000040-0000-1000-8000-0026BB765291"

type Fan struct {
	*Service

	On *characteristic.On
}

func NewFan() *Fan {
	svc := Fan{}
	svc.Service = New(TypeFan)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}
