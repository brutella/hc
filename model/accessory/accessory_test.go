package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"testing"
)

func TestAccessoryIdentifyChanged(t *testing.T) {
	a := New(info, TypeOther)

	var identifyCalled = 0
	a.OnIdentify(func() {
		identifyCalled++
	})

	a.Info.Identify.SetValueFromConnection(true, characteristic.TestConn)

	// Identify is set to false immediately
	if x := a.Info.Identify.GetValue(); x != nil {
		t.Fatal(x)
	}

	if is, want := identifyCalled, 1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestAccessoryInfo(t *testing.T) {
	var accessoryInfo = model.Info{
		Name:         "My Accessory",
		SerialNumber: "0009",
		Manufacturer: "Matthias",
		Model:        "1A",
		Firmware:     "0.1",
		Hardware:     "1.0",
		Software:     "2.1",
	}

	a := New(accessoryInfo, TypeOther)

	if is, want := a.Name(), "My Accessory"; is != want {
		t.Fatalf("name is=%v want=%v", is, want)
	}
	if is, want := a.SerialNumber(), "0009"; is != want {
		t.Fatalf("serialnumber is=%v want=%v", is, want)
	}
	if is, want := a.Manufacturer(), "Matthias"; is != want {
		t.Fatalf("manufacturer is=%v want=%v", is, want)
	}
	if is, want := a.Model(), "1A"; is != want {
		t.Fatalf("model is=%v want=%v", is, want)
	}
	if is, want := a.Firmware(), "0.1"; is != want {
		t.Fatalf("firmware is=%v want=%v", is, want)
	}
	if is, want := a.Hardware(), "1.0"; is != want {
		t.Fatalf("hardware is=%v want=%v", is, want)
	}
	if is, want := a.Software(), "2.1"; is != want {
		t.Fatalf("software is=%v want=%v", is, want)
	}
}
