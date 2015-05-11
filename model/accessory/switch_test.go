package accessory

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"testing"
)

var info = model.Info{
	Name:         "My Switch",
	SerialNumber: "001",
	Manufacturer: "Google",
	Model:        "Switchy",
}

func TestSwitch(t *testing.T) {
	var s model.Switch = NewSwitch(info)

	if is, want := s.IsOn(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	s.SetOn(true)

	if is, want := s.IsOn(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestSwitchOnChanged(t *testing.T) {
	s := NewSwitch(info)

	var newValue = false
	s.OnStateChanged(func(value bool) {
		newValue = value
	})

	s.switcher.On.SetValueFromConnection(true, characteristic.TestConn)

	if is, want := s.IsOn(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := newValue, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
