package model

// A LightBulb is a Switch and additionally has a brightness, saturation and hue value.
//
// TODO(brutella): The HAP protocol defines brightness, saturation and hue as optional. This
// is currently no reflected in the LightBulb interface yet.
type LightBulb interface {
	Switch

	// OnBrightnessChanged sets the brightness changed callback
	OnBrightnessChanged(func(int))

	// OnHueChanged sets the hue changed callback
	OnHueChanged(func(float64))

	// OnSaturationChanged sets the saturation changed callback
	OnSaturationChanged(func(float64))

	// GetBrightness returns the light bulb's brightness
	GetBrightness() int

	// SetBrightness sets the light bulb's brightness
	SetBrightness(int)

	// GetHue returns the light bulb's hue
	GetHue() float64

	// SetHue sets the light bulb's hue
	SetHue(float64)

	// GetSaturation returns the light bulb's saturation
	GetSaturation() float64

	// SetSaturation sets the light bulb's saturation
	SetSaturation(float64)
}
