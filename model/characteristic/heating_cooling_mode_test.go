package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "github.com/brutella/hap/model"
)

func TestHeatingCoolingMode(t *testing.T) {
    b := NewCurrentHeatingCoolingMode(model.ModeOff)
    assert.Equal(t, b.HeatingCoolingMode(), model.ModeOff)
    b.SetHeatingCoolingMode(model.ModeHeating)
    assert.Equal(t, b.HeatingCoolingMode(), model.ModeHeating)
}

func TestCurrentHeatingCoolingMode(t *testing.T) {
    b := NewCurrentHeatingCoolingMode(model.ModeOff)
    assert.Equal(t, b.Type, CharTypeHeatingCoolingModeCurrent)
}

func TestTargetHeatingCoolingMode(t *testing.T) {
    b := NewTargetHeatingCoolingMode(model.ModeOff)
    assert.Equal(t, b.Type, CharTypeHeatingCoolingModeTarget)
}