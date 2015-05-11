package characteristic

import (
	"github.com/brutella/hc/model"
	"testing"
)

func TestHeatingCoolingMode(t *testing.T) {
	b := NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)

	if is, want := b.HeatingCoolingMode(), model.HeatCoolModeOff; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b.SetHeatingCoolingMode(model.HeatCoolModeHeat)

	if is, want := b.HeatingCoolingMode(), model.HeatCoolModeHeat; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestCurrentHeatingCoolingMode(t *testing.T) {
	b := NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	if is, want := b.Type, CharTypeHeatingCoolingModeCurrent; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTargetHeatingCoolingMode(t *testing.T) {
	b := NewTargetHeatingCoolingMode(model.HeatCoolModeOff)
	if is, want := b.Type, CharTypeHeatingCoolingModeTarget; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
