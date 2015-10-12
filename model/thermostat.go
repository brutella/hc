package model

type HeatCoolModeType byte

const (
	HeatCoolModeOff  HeatCoolModeType = 0x00
	HeatCoolModeHeat HeatCoolModeType = 0x01
	HeatCoolModeCool HeatCoolModeType = 0x02
	HeatCoolModeAuto HeatCoolModeType = 0x03
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
	// HeatCoolModeAuto is ignored because the current mode cannot be auto
	SetMode(HeatCoolModeType)

	// Mode returns the current mode
	Mode() HeatCoolModeType

	// SetTargetMode sets the target mode
	SetTargetMode(HeatCoolModeType)

	// TargetMode returns the target mode
	TargetMode() HeatCoolModeType

	// OnTargetTempChange calls the argument function when the thermostat's
	// target temperature is changed
	OnTargetTempChange(func(float64))

	// OnTargetModeChange calls the argument function when the thermostat's
	// target mode is changed
	OnTargetModeChange(func(HeatCoolModeType))
}
