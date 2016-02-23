package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
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

	if is, want := bulb.GetBrightness(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	bulb.SetBrightness(90)

	if is, want := bulb.GetBrightness(), 90; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
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

	light.LightBulb.Brightness.SetValueFromConnection(80, characteristic.TestConn)
	light.LightBulb.Hue.SetValueFromConnection(15.5, characteristic.TestConn)
	light.LightBulb.Saturation.SetValueFromConnection(22.4, characteristic.TestConn)

	if is, want := newBrightness, 80; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := newHue, 15.5; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := newSaturation, 22.4; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
