package characteristic

import (
	"github.com/brutella/hc/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHeatingCoolingMode(t *testing.T) {
	b := NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	assert.Equal(t, b.HeatingCoolingMode(), model.HeatCoolModeOff)
	b.SetHeatingCoolingMode(model.HeatCoolModeHeat)
	assert.Equal(t, b.HeatingCoolingMode(), model.HeatCoolModeHeat)
}

func TestCurrentHeatingCoolingMode(t *testing.T) {
	b := NewCurrentHeatingCoolingMode(model.HeatCoolModeOff)
	assert.Equal(t, b.Type, CharTypeHeatingCoolingModeCurrent)
}

func TestTargetHeatingCoolingMode(t *testing.T) {
	b := NewTargetHeatingCoolingMode(model.HeatCoolModeOff)
	assert.Equal(t, b.Type, CharTypeHeatingCoolingModeTarget)
}
