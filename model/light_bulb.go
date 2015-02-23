package model

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
