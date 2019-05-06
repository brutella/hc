// THIS FILE IS AUTO-GENERATED
package service

import (
	"github.com/brutella/hc/characteristic"
)

const TypeLightbulb = "43"

type Lightbulb struct {
	*Service

	On *characteristic.On

	Brightness *characteristic.Brightness
	Hue        *characteristic.Hue
	Saturation *characteristic.Saturation
	Name       *characteristic.Name
}

func NewLightbulb() *Lightbulb {
	svc := Lightbulb{}
	svc.Service = New(TypeLightbulb)

	svc.On = characteristic.NewOn()
	svc.AddCharacteristic(svc.On.Characteristic)

	return &svc
}

func (svc *Lightbulb) addOptionalCharaterics() {

	svc.Brightness = characteristic.NewBrightness()
	svc.AddCharacteristic(svc.Brightness.Characteristic)

	svc.Hue = characteristic.NewHue()
	svc.AddCharacteristic(svc.Hue.Characteristic)

	svc.Saturation = characteristic.NewSaturation()
	svc.AddCharacteristic(svc.Saturation.Characteristic)

	svc.Name = characteristic.NewName()
	svc.AddCharacteristic(svc.Name.Characteristic)

}
