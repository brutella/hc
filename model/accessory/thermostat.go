package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"github.com/brutella/hc/model/service"
	"net"
)

type Thermostat struct {
	*Accessory

	Thermostat *service.Thermostat
}

// NewThermostat returns a Thermostat which implements model.Thermostat.
func NewThermostat(info model.Info, temp, min, max, steps float64) *Thermostat {
	accessory := New(info)
	t := service.NewThermostat(info.Name, temp, min, max, steps)

	accessory.AddService(t.Service)

	return &Thermostat{accessory, t}
}

func (t *Thermostat) Temperature() float64 {
	return t.Thermostat.Temp.Temperature()
}

func (t *Thermostat) SetTemperature(value float64) {
	t.Thermostat.Temp.SetTemperature(value)
}

func (t *Thermostat) Unit() model.TempUnit {
	return t.Thermostat.Unit.Unit()
}

func (t *Thermostat) SetTargetTemperature(value float64) {
	t.Thermostat.TargetTemp.SetTemperature(value)
}

func (t *Thermostat) TargetTemperature() float64 {
	return t.Thermostat.TargetTemp.Temperature()
}

func (t *Thermostat) SetMode(value model.HeatCoolModeType) {
	if value != model.HeatCoolModeAuto {
		t.Thermostat.Mode.SetHeatingCoolingMode(value)
	}
}

func (t *Thermostat) Mode() model.HeatCoolModeType {
	return t.Thermostat.Mode.HeatingCoolingMode()
}

func (t *Thermostat) SetTargetMode(value model.HeatCoolModeType) {
	t.Thermostat.TargetMode.SetHeatingCoolingMode(value)
}

func (t *Thermostat) TargetMode() model.HeatCoolModeType {
	return t.Thermostat.TargetMode.HeatingCoolingMode()
}

func (t *Thermostat) OnTargetTempChange(fn func(float64)) {
	t.Thermostat.TargetTemp.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}

func (t *Thermostat) OnTargetModeChange(fn func(model.HeatCoolModeType)) {
	t.Thermostat.TargetMode.OnConnChange(func(conn net.Conn, c *characteristic.Characteristic, new, old interface{}) {
		fn(model.HeatCoolModeType(new.(byte)))
	})
}
