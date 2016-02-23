package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type LightBulb struct {
	*Accessory
	LightBulb *service.LightBulb
}

// NewLightBulb returns a light bulb which implements model.LightBulb.
func NewLightBulb(info model.Info) *LightBulb {
	accessory := New(info)
	bulb := service.NewLightBulb(info.Name, false) // off

	accessory.AddService(bulb.Service)

	return &LightBulb{accessory, bulb}
}

func (l *LightBulb) SetOn(on bool) {
	l.LightBulb.On.SetOn(on)
}

func (l *LightBulb) IsOn() bool {
	return l.LightBulb.On.On()
}

func (l *LightBulb) GetBrightness() int {
	return l.LightBulb.Brightness.IntValue()
}

func (l *LightBulb) SetBrightness(value int) {
	l.LightBulb.Brightness.SetInt(value)
}

func (l *LightBulb) GetHue() float64 {
	return l.LightBulb.Hue.FloatValue()
}

func (l *LightBulb) SetHue(value float64) {
	l.LightBulb.Hue.SetFloat(value)
}

func (l *LightBulb) GetSaturation() float64 {
	return l.LightBulb.Saturation.FloatValue()
}

func (l *LightBulb) SetSaturation(value float64) {
	l.LightBulb.Saturation.SetFloat(value)
}

func (l *LightBulb) OnStateChanged(fn func(bool)) {
	l.LightBulb.On.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}

func (l *LightBulb) OnBrightnessChanged(fn func(int)) {
	l.LightBulb.Brightness.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(int))
	})
}

func (l *LightBulb) OnHueChanged(fn func(float64)) {
	l.LightBulb.Hue.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}

func (l *LightBulb) OnSaturationChanged(fn func(float64)) {
	l.LightBulb.Saturation.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}
