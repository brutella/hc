package service

import (
	"github.com/brutella/hc/model"
	"github.com/brutella/hc/model/characteristic"
	"reflect"
	"testing"
)

func TestContactSensor(t *testing.T) {
	sw := NewContactSensor("My Sensor")

	if is, want := sw.Type, typeContactSensor; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := sw.Name.GetValue(), "My Sensor"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if x := sw.Name.Permissions; reflect.DeepEqual(x, characteristic.PermsRead()) == false {
		t.Fatal(x)
	}
	if x := sw.ContactSensorState.Permissions; reflect.DeepEqual(x, characteristic.PermsRead()) == false {
		t.Fatal(x)
	}
	if is, want := sw.ContactSensorState.GetValue(), uint8(model.ContactNotDetected); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
