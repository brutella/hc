package accessory

import (
	"github.com/brutella/hc/service"
)

//LightSensor struct
type LightSensor struct {
	*Accessory

	LightSensor *service.LightSensor
}

// NewLightSensor returns a Thermometer which implements model.Thermometer.
func NewLightSensor(info Info) *LightSensor {
	acc := LightSensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.LightSensor = service.NewLightSensor()

	acc.AddService(acc.LightSensor.Service)

	return &acc
}
