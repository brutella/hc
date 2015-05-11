package accessory

import (
	"github.com/brutella/hc/model"
	"testing"
)

func TestThermostat(t *testing.T) {
	info := model.Info{
		Name:         "My Thermostat",
		SerialNumber: "001",
		Manufacturer: "Google",
		Model:        "Thermostaty",
	}

	var thermo model.Thermostat = NewThermostat(info, 10, 0, 100, 1)

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
