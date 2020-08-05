package accessory

import (
	"github.com/brutella/hc/service"
)

type EnviroPlus struct {
	*Accessory
	TemperatureSensor *service.TemperatureSensor
	HumiditySensor *service.HumiditySensor
	AirQualitySensor *service.AirQualitySensor
	LightSensor *service.LightSensor
}

func NewEnviroPlus(info Info) *EnviroPlus {
	acc := EnviroPlus{}
	acc.Accessory = New(info, TypeSensor)

	acc.TemperatureSensor = service.NewTemperatureSensor()
	acc.AddService(acc.TemperatureSensor.Service)

	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddService(acc.HumiditySensor.Service)

	acc.AirQualitySensor = service.NewAirQualitySensor()
	acc.AddService(acc.AirQualitySensor.Service)

	acc.LightSensor = service.NewLightSensor()
	acc.AddService(acc.LightSensor.Service)

	return &acc
}

