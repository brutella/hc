package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
)

type lightBulb struct {
	*Accessory
	bulb *service.LightBulb

	onChanged         func(bool)
	brightnessChanged func(int)
	saturationChanged func(float64)
	hueChanged        func(float64)
}

// NewLightBulb returns a light bulb which implements model.LightBulb.
func NewLightBulb(info model.Info) *lightBulb {
	accessory := New(info)
	bulb := service.NewLightBulb(info.Name, false) // off

	accessory.AddService(bulb.Service)

	lightBulb := lightBulb{accessory, bulb, nil, nil, nil, nil}

	bulb.On.OnRemoteChange(func(c *characteristic.Characteristic, new, old interface{}) {
		if lightBulb.onChanged != nil {
			lightBulb.onChanged(bulb.On.On())
		}
	})

	bulb.Brightness.OnRemoteChange(func(c *characteristic.Characteristic, new, old interface{}) {
		if lightBulb.brightnessChanged != nil {
			lightBulb.brightnessChanged(bulb.Brightness.IntValue())
		}
	})

	bulb.Hue.OnRemoteChange(func(c *characteristic.Characteristic, new, old interface{}) {
		if lightBulb.hueChanged != nil {
			lightBulb.hueChanged(bulb.Hue.FloatValue())
		}
	})

	bulb.Saturation.OnRemoteChange(func(c *characteristic.Characteristic, new, old interface{}) {
		if lightBulb.saturationChanged != nil {
			lightBulb.saturationChanged(bulb.Saturation.FloatValue())
		}
	})

	return &lightBulb
}

func (l *lightBulb) SetOn(on bool) {
	l.bulb.On.SetOn(on)
}

func (l *lightBulb) IsOn() bool {
	return l.bulb.On.On()
}

func (l *lightBulb) GetBrightness() int {
	return l.bulb.Brightness.IntValue()
}

func (l *lightBulb) SetBrightness(value int) {
	l.bulb.Brightness.SetInt(value)
}

func (l *lightBulb) GetHue() float64 {
	return l.bulb.Hue.FloatValue()
}

func (l *lightBulb) SetHue(value float64) {
	l.bulb.Hue.SetFloat(value)
}

func (l *lightBulb) GetSaturation() float64 {
	return l.bulb.Saturation.FloatValue()
}

func (l *lightBulb) SetSaturation(value float64) {
	l.bulb.Saturation.SetFloat(value)
}

func (l *lightBulb) OnStateChanged(fn func(bool)) {
	l.onChanged = fn
}

func (l *lightBulb) OnBrightnessChanged(fn func(int)) {
	l.brightnessChanged = fn
}

func (l *lightBulb) OnHueChanged(fn func(float64)) {
	l.hueChanged = fn
}

func (l *lightBulb) OnSaturationChanged(fn func(float64)) {
	l.saturationChanged = fn
}
