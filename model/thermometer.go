package model

type TempUnit string

const (
	TempUnitCelsius = TempUnit("celsius")
)

// A Thermometer measures the temperature.
//
// Discussion: Because there is no thermometer accessory type in HomeKit, we
// use a readonly thermostat as a thermometer. This means that the target
// temperature and current/target heating-cooling modes characteristics
// are defined readonly.
type Thermometer interface {
	Accessory

	// SetTemperature sets the current temperature
	SetTemperature(float64)

	// Temperature returns the current temperature
	Temperature() float64

	// Unit returns the temperature unit
	Unit() TempUnit
}
