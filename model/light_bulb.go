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

	// GetBrightness returns the light bulb's brightness between 0 and 100
	GetBrightness() int

	// SetBrightness sets the light bulb's brightness
	// The argument should be between 0 and 100
	SetBrightness(int)

	// GetHue returns the light bulb's hue between 0.0 and 360.0
	GetHue() float64

	// SetHue sets the light bulb's hue
	// The argument should be between 0.0 and 360.0
	SetHue(float64)

	// GetSaturation returns the light bulb's saturation between 0.0 and 100.0
	GetSaturation() float64

	// SetSaturation sets the light bulb's saturation
	// The argument should be between 0 and 100
	SetSaturation(float64)
}
