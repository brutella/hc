package accessory

import (
	"github.com/brutella/hc/service"
)

type Thermometer struct {
	*Accessory

	TempSensor *service.TemperatureSensor
}

// NewTemperatureSensor returns a Thermometer which implements model.Thermometer.
func NewTemperatureSensor(info Info, temp, min, max, steps float64) *Thermometer {
	acc := Thermometer{}
	acc.Accessory = New(info, TypeThermostat)
	acc.TempSensor = service.NewTemperatureSensor()
	acc.TempSensor.CurrentTemperature.SetValue(temp)
	acc.TempSensor.CurrentTemperature.SetMinValue(min)
	acc.TempSensor.CurrentTemperature.SetMaxValue(max)
	acc.TempSensor.CurrentTemperature.SetStepValue(steps)

	acc.AddService(acc.TempSensor.Service)

	return &acc
}
