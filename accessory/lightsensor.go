package accessory

import (
	"github.com/brutella/hc/service"
)

type AmbientLightLevel struct {
	*Accessory
	LightSensor *service.LightSensor
}

// NewAccAmbientLightLevel
func NewAmbientLightLevel(info Info) *AmbientLightLevel {
	acc := AmbientLightLevel{}
	acc.Accessory = New(info, TypeSensor)
	acc.LightSensor = service.NewLightSensor()
	acc.LightSensor.StatusActive.SetValue(true)
	
	acc.AddService(acc.LightSensor.Service)

	return &acc
}
