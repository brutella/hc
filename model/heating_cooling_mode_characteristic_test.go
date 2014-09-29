package hk

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestHeatingCoolingModeCharacteristic(t *testing.T) {
    b := NewCurrentHeatingCoolingModeCharacteristic(ModeOff)
    assert.Equal(t, b.HeatingCoolingMode(), ModeOff)
    b.SetHeatingCoolingMode(ModeHeating)
    assert.Equal(t, b.HeatingCoolingMode(), ModeHeating)
}

func TestCurrentHeatingCoolingModeCharacteristic(t *testing.T) {
    b := NewCurrentHeatingCoolingModeCharacteristic(ModeOff)
    assert.Equal(t, b.Type, CharTypeHeatingCoolingModeCurrent)
}

func TestTargetHeatingCoolingModeCharacteristic(t *testing.T) {
    b := NewTargetHeatingCoolingModeCharacteristic(ModeOff)
    assert.Equal(t, b.Type, CharTypeHeatingCoolingModeTarget)
}