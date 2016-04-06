package characteristic

import (
	"testing"
)

func TestHeatingCoolingMode(t *testing.T) {
	b := NewCurrentHeatingCoolingState()
	b.Value = CurrentHeatingCoolingStateOff

	if is, want := b.GetValue(), CurrentHeatingCoolingStateOff; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b.SetValue(CurrentHeatingCoolingStateHeat)

	if is, want := b.GetValue(), CurrentHeatingCoolingStateHeat; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
