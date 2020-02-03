package accessory

import (
	"github.com/brutella/hc/service"
)

//HumiditySensor struct
type HumiditySensor struct {
	*Accessory

	HumiditySensor *service.HumiditySensor
}

// NewHumiditySensor returns a Thermometer which implements model.Thermometer.
func NewHumiditySensor(info Info) *HumiditySensor {
	acc := HumiditySensor{}
	acc.Accessory = New(info, TypeThermostat)
	acc.HumiditySensor = service.NewHumiditySensor()

	acc.AddService(acc.HumiditySensor.Service)

	return &acc
}
