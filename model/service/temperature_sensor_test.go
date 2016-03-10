package service

import (
	"testing"
)

func TestThermometer(t *testing.T) {
	thermometer := NewTemperatureSensor("Thermometer", 10.5, -10, 100, 1)

	if is, want := thermometer.Type, TypeTemperatureSensor; is != want {
		t.Fatalf("type: is=%v want=%v", is, want)
	}
	if is, want := thermometer.Name.GetValue(), "Thermometer"; is != want {
		t.Fatalf("name: is=%v want=%v", is, want)
	}
	if is, want := thermometer.Temp.GetValue(), 10.5; is != want {
		t.Fatalf("temp: is=%v want=%v", is, want)
	}
}
