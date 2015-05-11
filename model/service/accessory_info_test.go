package service

import (
	"github.com/brutella/hc/model"
	"testing"
)

func TestAccessoryInfo(t *testing.T) {
	info := model.Info{
		Name:         "Test Accessory",
		SerialNumber: "001",
		Manufacturer: "Matthias",
		Model:        "Version 123",
	}

	i := NewInfo(info)

	if is, want := i.Type, typeAccessoryInfo; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := i.Identify.GetValue(); x != nil {
		t.Fatal(x)
	}
	if is, want := i.Serial.GetValue(), "001"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := i.Model.GetValue(), "Version 123"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := i.Manufacturer.GetValue(), "Matthias"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := i.Name.GetValue(), "Test Accessory"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := i.Firmware; x != nil {
		t.Fatal(x)
	}
	if x := i.Hardware; x != nil {
		t.Fatal(x)
	}
	if x := i.Software; x != nil {
		t.Fatal(x)
	}
}
