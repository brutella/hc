package accessory

import (
    "github.com/brutella/hap/model"
    
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestLightBulb(t *testing.T) {
    info := model.Info{
        Name: "My Light Bulb",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Switchy",
    }
    
    var bulb model.LightBulb = NewLightBulb(info)
    
    assert.Equal(t, bulb.GetId(), model.InvalidId)
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
        Name: "My Light Bulb",
        SerialNumber: "001",
        Manufacturer: "Google",
        Model: "Switchy",
    }
    
    light := NewLightBulb(info)
    
    var newBrightness int = 0
    var newSaturation float64 = 0
    var newHue float64 = 0
    light.OnBrightnessChanged(func(value int) {
        newBrightness = value
    })
    
    light.OnHueChanged(func(value float64) {
        newHue = value
    })
    
    light.OnSaturationChanged(func(value float64) {
        newSaturation = value
    })
    
    light.bulb.Brightness.SetValueFromRemote(80)
    light.bulb.Hue.SetValueFromRemote(15.5)
    light.bulb.Saturation.SetValueFromRemote(22.4)
    assert.Equal(t, newBrightness, 80)
    assert.Equal(t, newHue, 15.5)
    assert.Equal(t, newSaturation, 22.4)
}