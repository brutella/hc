package accessory

import (
	"github.com/brutella/hc/model"
	"testing"
)

var contact_sensor_info = model.Info{
	Name:         "My Sensor",
	SerialNumber: "001",
	Manufacturer: "My Co",
	Model:        "Sensor1",
}

func TestContactSensor(t *testing.T) {
	var s model.ContactSensor = NewContactSensor(contact_sensor_info)

	if is, want := s.State(), model.ContactNotDetected; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	s.SetState(model.ContactDetected)

	if is, want := s.State(), model.ContactDetected; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
