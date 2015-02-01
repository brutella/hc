package model

type TempUnit string

const (
	TempUnitCelsius = TempUnit("celsius")
)

// A thermometer is an acessory which measures temperature
// HomeKit
type Thermometer interface {
	Accessory

	SetTemperature(float64)
	Temperature() float64
	Unit() TempUnit
}
