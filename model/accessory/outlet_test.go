package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"testing"
)

var outlet_info = model.Info{
	Name:         "My Outlet",
	SerialNumber: "001",
	Manufacturer: "brutella",
	Model:        "Outletty",
}

func TestOutlet(t *testing.T) {
	var o model.Outlet = NewOutlet(outlet_info)

	if is, want := o.IsOn(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := o.IsInUse(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	o.SetOn(true)
	o.SetInUse(false)

	if is, want := o.IsOn(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := o.IsInUse(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestOutletOnChanged(t *testing.T) {
	o := NewOutlet(outlet_info)

	var newValue = false
	o.OnStateChanged(func(value bool) {
		newValue = value
	})

	o.outlet.On.SetValueFromConnection(true, characteristic.TestConn)

	if is, want := o.IsOn(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := newValue, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
