package characteristic

import (
	"github.com/brutella/hc/model"
	"testing"
)

func TestContactSensorState(t *testing.T) {
	b := NewCurrentContactSensorState(model.ContactDetected)

	if is, want := b.ContactSensorState(), model.ContactDetected; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b.SetContactSensorState(model.ContactNotDetected)

	if is, want := b.ContactSensorState(), model.ContactNotDetected; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCurrentContactSensorState(t *testing.T) {
	b := NewCurrentContactSensorState(model.ContactDetected)
	if is, want := b.Type, TypeContactSensorState; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}