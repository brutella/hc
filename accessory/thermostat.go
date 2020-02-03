package accessory

import (
	"github.com/brutella/hc/service"
)

type Thermostat struct {
	*Accessory

	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat which implements model.Thermostat.
func NewThermostat(info Info, temp, minTemp, maxTemp, stepsTemp float64, setTHCS, minTHCS, maxTHCS, stepsTHCS int) *Thermostat {
	acc := Thermostat{}
	acc.Accessory = New(info, TypeThermostat)
	acc.Thermostat = service.NewThermostat()

	acc.Thermostat.TargetTemperature.SetValue(temp)
	acc.Thermostat.TargetTemperature.SetMinValue(minTemp)
	acc.Thermostat.TargetTemperature.SetMaxValue(maxTemp)
	acc.Thermostat.TargetTemperature.SetStepValue(stepsTemp)

	acc.Thermostat.TargetHeatingCoolingState.SetValue(setTHCS)
	acc.Thermostat.TargetHeatingCoolingState.SetMinValue(minTHCS)
	acc.Thermostat.TargetHeatingCoolingState.SetMaxValue(maxTHCS)
	acc.Thermostat.TargetHeatingCoolingState.SetStepValue(stepsTHCS)

	acc.AddService(acc.Thermostat.Service)

	return &acc
}
