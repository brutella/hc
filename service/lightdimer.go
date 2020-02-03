// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLightDimer = "43"

type LightDimer struct {
	*Service

	On         *characteristic.On
	Brightness *characteristic.Brightness
}

func NewLightDimer() *LightDimer {
	svc := LightDimer{}
	svc.Service = New(TypeLightDimer)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	return &svc
}
