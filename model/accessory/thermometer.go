package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/service"
)

type thermometer struct {
	*Accessory

	temperatureSensor *service.TemperatureSensor
}

// NewTemperatureSensor returns a thermometer  which implements model.Thermometer.
func NewTemperatureSensor(info model.Info, temp, min, max, steps float64) *thermometer {
	accessory := New(info)
	t := service.NewTemperatureSensor(info.Name, temp, min, max, steps)

	accessory.AddService(t.Service)

	return &thermometer{accessory, t}
}

func (t *thermometer) Temperature() float64 {
	return t.temperatureSensor.Temp.Temperature()
}

func (t *thermometer) SetTemperature(value float64) {
	t.temperatureSensor.Temp.SetTemperature(value)
}

func (t *thermometer) Unit() model.TempUnit {
	return t.temperatureSensor.Unit.Unit()
}
