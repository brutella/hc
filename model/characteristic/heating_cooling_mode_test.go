package characteristic

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestHeatingCoolingMode(t *testing.T) {
    b := NewCurrentHeatingCoolingMode(ModeOff)
    assert.Equal(t, b.HeatingCoolingMode(), ModeOff)
    b.SetHeatingCoolingMode(ModeHeating)
    assert.Equal(t, b.HeatingCoolingMode(), ModeHeating)
}

func TestCurrentHeatingCoolingMode(t *testing.T) {
    b := NewCurrentHeatingCoolingMode(ModeOff)
    assert.Equal(t, b.Type, CharTypeHeatingCoolingModeCurrent)
}

func TestTargetHeatingCoolingMode(t *testing.T) {
    b := NewTargetHeatingCoolingMode(ModeOff)
    assert.Equal(t, b.Type, CharTypeHeatingCoolingModeTarget)
}