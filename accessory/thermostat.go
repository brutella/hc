package accessory

import (
	"github.com/brutella/hc/service"
)

type Thermostat struct {
	*Accessory

	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat which implements model.Thermostat.
func NewThermostat(info Info, temp, min, max, steps float64) *Thermostat {
	acc := Thermostat{}
	acc.Accessory = New(info, TypeThermostat)
	acc.Thermostat = service.NewThermostat()
	acc.Thermostat.CurrentTemperature.SetValue(temp)
	acc.Thermostat.CurrentTemperature.SetMinValue(min)
	acc.Thermostat.CurrentTemperature.SetMaxValue(max)
	acc.Thermostat.CurrentTemperature.SetStepValue(steps)

	acc.Thermostat.TargetTemperature.SetValue(temp)
	acc.Thermostat.TargetTemperature.SetMinValue(min)
	acc.Thermostat.TargetTemperature.SetMaxValue(max)
	acc.Thermostat.TargetTemperature.SetStepValue(steps)

	acc.AddService(acc.Thermostat.Service)

	return &acc
}
