package service

import (
	"github.com/brutella/hc/characteristic"
)

type ColoredLightbulb struct {
	*Service

	On         *characteristic.On
	Brightness *characteristic.Brightness
	Saturation *characteristic.Saturation
	Hue        *characteristic.Hue
}

func NewColoredLightbulb() *ColoredLightbulb {
	svc := ColoredLightbulb{}
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
