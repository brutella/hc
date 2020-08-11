package accessory

import (
	"github.com/brutella/hc/characteristic"
	"github.com/brutella/hc/service"
)

// EnviroPlus provides air quality, light and motion sensor readings
type EnviroPlus struct {
	*Accessory
	TemperatureSensor *service.TemperatureSensor
	HumiditySensor    *service.HumiditySensor
	AirQualitySensor  *service.AirQualitySensor
	LightSensor       *service.LightSensor
	MotionSensor      *service.MotionSensor
}

// NewEnviroPlus returns an EnviroPlus accessory
func NewEnviroPlus(info Info) *EnviroPlus {
	acc := EnviroPlus{}
	acc.Accessory = New(info, TypeSensor)

	acc.TemperatureSensor = service.NewTemperatureSensor()
	acc.AddService(acc.TemperatureSensor.Service)

	acc.HumiditySensor = service.NewHumiditySensor()
	acc.AddService(acc.HumiditySensor.Service)

	acc.AirQualitySensor = service.NewAirQualitySensor()
	pm25 := characteristic.NewPM2_5Density()
	pm10 := characteristic.NewPM10Density()
	carbonMonoxide := characteristic.NewCarbonMonoxideLevel()
	nitrogenDioxide := characteristic.NewNitrogenDioxideDensity()
	acc.AirQualitySensor.Service.AddCharacteristic(pm25.Characteristic)
	acc.AirQualitySensor.Service.AddCharacteristic(pm10.Characteristic)
	acc.AirQualitySensor.Service.AddCharacteristic(carbonMonoxide.Characteristic)
	acc.AirQualitySensor.Service.AddCharacteristic(nitrogenDioxide.Characteristic)
	acc.AddService(acc.AirQualitySensor.Service)

	acc.LightSensor = service.NewLightSensor()
	acc.AddService(acc.LightSensor.Service)

	acc.MotionSensor = service.NewMotionSensor()
	acc.AddService(acc.MotionSensor.Service)

	return &acc
}
