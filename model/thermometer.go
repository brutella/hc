package model

type TempUnit string

const (
	TempUnitCelsius = TempUnit("celsius")
)

// A thermometer measures the temperature
type Thermometer interface {
	Accessory

	// SetTemperature sets the current temperature
	SetTemperature(float64)

	// Temperature returns the current temperature
	Temperature() float64

	// Unit returns the temperature unit
	Unit() TempUnit
}
