package model

// A hygrometer measures humidity and let you set a target humidity.
type Hygrometer interface {
	Humidity() float64

	SetTargetHumidity(float64)
	TargetHumidity() float64
}
