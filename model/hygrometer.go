package model

// A Hygrometer measures humidity and let you set a target humidity.
// This type is currently not implemented.
type Hygrometer interface {
	Humidity() float64

	SetTargetHumidity(float64)
	TargetHumidity() float64
}
