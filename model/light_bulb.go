package model

type LightBulb interface {
	Switch

	// Sets the brightness changed callback
	OnBrightnessChanged(func(int))

	// Sets the hue changed callback
	OnHueChanged(func(float64))

	// Sets the saturation changed callback
	OnSaturationChanged(func(float64))

	// Returns the light bulb's brightness
	GetBrightness() int

	// Sets the light bulb's brightness
	SetBrightness(int)

	// Returns the light bulb's hue
	GetHue() float64

	// Sets the light bulb's hue
	SetHue(float64)

	// Returns the light bulb's saturation
	GetSaturation() float64

	// Sets the light bulb's saturation
	SetSaturation(float64)
}
