package model

type HeatCoolMode byte

const (
	ModeOff     = HeatCoolMode(0x00)
	ModeHeating = HeatCoolMode(0x01)
	ModeCooling = HeatCoolMode(0x02)
	ModeAuto    = HeatCoolMode(0x03)
)

// A thermostat measures and lets you change the  mode (heating, cooling, auto)
// to reach a certain target temperature
type Thermostat interface {
	Thermometer

	// SetTargetTemperature sets the target temperature
	SetTargetTemperature(float64)

	// TargetTemperature returns the target temperature
	TargetTemperature() float64

	// SetMode sets the current mode
	// ModeAuto is ignored because the current mode cannot be auto
	SetMode(HeatCoolMode)

	// Mode returns the current mode
	Mode() HeatCoolMode

	// SetTargetMode sets the target mode
	SetTargetMode(HeatCoolMode)

	// TargetMode returns the target mode
	TargetMode() HeatCoolMode
}
