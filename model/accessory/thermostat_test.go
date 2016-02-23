package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"testing"
)

var thermostat_info = model.Info{
	Name:         "My Thermostat",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Thermostaty",
}

func TestThermostat(t *testing.T) {

	var thermo model.Thermostat = NewThermostat(thermostat_info, 10, 0, 100, 1)

	if is, want := thermo.Temperature(), 10.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := thermo.TargetTemperature(), 10.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := thermo.TargetMode(), model.HeatCoolModeOff; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := thermo.Mode(), model.HeatCoolModeOff; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	thermo.SetTemperature(11)
	thermo.SetTargetTemperature(12)

	if is, want := thermo.Temperature(), 11.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := thermo.TargetTemperature(), 12.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTargetTempCallback(t *testing.T) {
	ts := NewThermostat(thermostat_info, 10, 0, 100, 1)

	var newValue float64
	ts.OnTargetTempChange(func(value float64) {
		newValue = value
	})

	ts.Thermostat.TargetTemp.SetValueFromConnection(25.2, characteristic.TestConn)

	if is, want := ts.TargetTemperature(), 25.2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTargetModeCallback(t *testing.T) {
	ts := NewThermostat(thermostat_info, 10, 0, 100, 1)
	ts.SetTargetMode(model.HeatCoolModeHeat)

	var newValue model.HeatCoolModeType
	ts.OnTargetModeChange(func(value model.HeatCoolModeType) {
		newValue = value
	})

	ts.Thermostat.TargetMode.SetValueFromConnection(model.HeatCoolModeCool, characteristic.TestConn)

	if is, want := ts.TargetMode(), model.HeatCoolModeCool; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
