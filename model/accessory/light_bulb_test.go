package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"

	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLightBulb(t *testing.T) {
	info := model.Info{
		Name:         "My Light Bulb",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Switchy",
	}

	var bulb model.LightBulb = NewLightBulb(info)

	assert.Equal(t, bulb.Name(), "My Light Bulb")
	assert.Equal(t, bulb.SerialNumber(), "001")
	assert.Equal(t, bulb.Manufacturer(), "Google")
	assert.Equal(t, bulb.Model(), "Switchy")
	assert.Equal(t, bulb.Model(), "Switchy")
	assert.Equal(t, bulb.GetBrightness(), 100)
	bulb.SetBrightness(90)
	assert.Equal(t, bulb.GetBrightness(), 90)
}

func TestLightBulbCallbacks(t *testing.T) {
	info := model.Info{
		Name:         "My Light Bulb",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Switchy",
	}

	light := NewLightBulb(info)

	var newBrightness int
	var newSaturation float64
	var newHue float64
	light.OnBrightnessChanged(func(value int) {
		newBrightness = value
	})

	light.OnHueChanged(func(value float64) {
		newHue = value
	})

	light.OnSaturationChanged(func(value float64) {
		newSaturation = value
	})

	light.bulb.Brightness.SetValueFromConnection(80, characteristic.TestConn)
	light.bulb.Hue.SetValueFromConnection(15.5, characteristic.TestConn)
	light.bulb.Saturation.SetValueFromConnection(22.4, characteristic.TestConn)
	assert.Equal(t, newBrightness, 80)
	assert.Equal(t, newHue, 15.5)
	assert.Equal(t, newSaturation, 22.4)
}
