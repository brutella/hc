// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLightbulb = "00000043-0000-1000-8000-0026BB765291"

type Lightbulb struct {
	*Service

	On         *characteristic.On
	Brightness *characteristic.Brightness
	Saturation *characteristic.Saturation
	Hue        *characteristic.Hue
}

func NewLightbulb() *Lightbulb {
	svc := Lightbulb{}
	svc.Service = New(TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	svc.Saturation = characteristic.NewSaturation()
	svc.AddCharacteristic(svc.Saturation.Characteristic)

	svc.Hue = characteristic.NewHue()
	svc.AddCharacteristic(svc.Hue.Characteristic)

	return &svc
}
