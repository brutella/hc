package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type lightBulb struct {
	*Accessory
	bulb *service.LightBulb
}

// NewLightBulb returns a light bulb which implements model.LightBulb.
func NewLightBulb(info model.Info) *lightBulb {
	accessory := New(info)
	bulb := service.NewLightBulb(info.Name, false) // off

	accessory.AddService(bulb.Service)

	return &lightBulb{accessory, bulb}
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
	l.bulb.On.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}

func (l *lightBulb) OnBrightnessChanged(fn func(int)) {
	l.bulb.Brightness.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(int))
	})
}

func (l *lightBulb) OnHueChanged(fn func(float64)) {
	l.bulb.Hue.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}

func (l *lightBulb) OnSaturationChanged(fn func(float64)) {
	l.bulb.Saturation.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}
