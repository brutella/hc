package accessory

import (
	"github.com/brutella/hc/service"
)

type Humidity struct {
	*Accessory

	HumiditySensor *service.HumiditySensor
}

// NewHumiditySensor returns a Humidity which implements model.Humidity.
func NewHumiditySensor(info Info, hum, min, max, steps float64) *Humidity {
	acc := Humidity{}
	acc.Accessory = New(info, TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()
	acc.HumiditySensor.CurrentRelativeHumidity.SetValue(hum)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMinValue(min)
	acc.HumiditySensor.CurrentRelativeHumidity.SetMaxValue(max)
	acc.HumiditySensor.CurrentRelativeHumidity.SetStepValue(steps)

	acc.AddService(acc.HumiditySensor.Service)

	return &acc
}
