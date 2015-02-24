package service

import (
	"github.com/brutella/hap/model/characteristic"
)

type LightBulb struct {
	*Service
	On         *characteristic.On
	Name       *characteristic.Name
	Brightness *characteristic.Brightness
	Saturation *characteristic.Saturation
	Hue        *characteristic.Hue
}

// NewLightBulb returns a light bulb service.
func NewLightBulb(name string, on bool) *LightBulb {
	on_char := characteristic.NewOn(on)
	name_char := characteristic.NewName(name)
	brightness := characteristic.NewBrightness(100) // 100%
	saturation := characteristic.NewSaturation(0.0)
	hue := characteristic.NewHue(0.0)

	service := New()
	service.Type = TypeLightBulb
	service.AddCharacteristic(on_char.Characteristic)
	service.AddCharacteristic(name_char.Characteristic)
	service.AddCharacteristic(brightness.Characteristic)
	service.AddCharacteristic(saturation.Characteristic)
	service.AddCharacteristic(hue.Characteristic)

	return &LightBulb{service, on_char, name_char, brightness, saturation, hue}
}
