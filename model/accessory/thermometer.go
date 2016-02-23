package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/service"
)

type Thermometer struct {
	*Accessory

	TempSensor *service.TemperatureSensor
}

// NewTemperatureSensor returns a Thermometer which implements model.Thermometer.
func NewTemperatureSensor(info model.Info, temp, min, max, steps float64) *Thermometer {
	accessory := New(info)
	t := service.NewTemperatureSensor(info.Name, temp, min, max, steps)

	accessory.AddService(t.Service)

	return &Thermometer{accessory, t}
}

func (t *Thermometer) Temperature() float64 {
	return t.TempSensor.Temp.Temperature()
}

func (t *Thermometer) SetTemperature(value float64) {
	t.TempSensor.Temp.SetTemperature(value)
}

func (t *Thermometer) Unit() model.TempUnit {
	return t.TempSensor.Unit.Unit()
}
