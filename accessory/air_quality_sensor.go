package accessory

import (
	"github.com/brutella/hc/service"
)

//AirQualitySensor structure
type AirQualitySensor struct {
	*Accessory
	AirQualitySensor *service.AirQualitySensor
}

// NewAirQualitySensor returns an outlet accessory containing one outlet service.
func NewAirQualitySensor(info Info) *AirQualitySensor {
	acc := AirQualitySensor{}
	acc.Accessory = New(info, TypeSensor)
	acc.AirQualitySensor = service.NewAirQualitySensor()

	acc.AddService(acc.AirQualitySensor.Service)

	return &acc
}
