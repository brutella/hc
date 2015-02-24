package model

type HeatCoolMode byte

const (
	ModeOff     = HeatCoolMode(0x00)
	ModeHeating = HeatCoolMode(0x01)
	ModeCooling = HeatCoolMode(0x02)
	ModeAuto    = HeatCoolMode(0x03)
)

// A Thermostat is a Thermometer but additionally lets you change the target temperature
// and  mode (heating, cooling, auto).
//
// TODO(brutella): The HAP protocol defines additional optional properties (heating- and
// cooling  threshold, current and target relative humidity), which are not implemented yet.
// The humidity values are currently defined in the Hygrometer interface, which is not
// implemented yet.
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
