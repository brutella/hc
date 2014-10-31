package model

type LightBulb interface {
    Switch
    
    OnBrightnessChanged(func(int))
    
    GetBrightness() int
    GetHue() float64
    GetSaturation() float64
}