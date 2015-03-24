package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTemperatureCharacteristic(t *testing.T) {
	temp := NewCurrentTemperatureCharacteristic(20.2, 0, 100, 1, "celsius")
	assert.Equal(t, temp.Temperature(), 20.2)
	assert.Equal(t, temp.MinTemperature(), 0)
	assert.Equal(t, temp.MaxTemperature(), 100)
	assert.Equal(t, temp.MinStepTemperature(), 1)

	temp.SetTemperature(10.1)
	assert.Equal(t, temp.Temperature(), 10.1)
}

func TestCurrentTemperatureCharacteristic(t *testing.T) {
	temp := NewCurrentTemperatureCharacteristic(20.2, 0, 100, 1, "celsius")
	assert.Equal(t, temp.Permissions, PermsRead())
	assert.Equal(t, temp.Type, CharTypeTemperatureCurrent)
}

func TestTargetTemperatureCharacteristic(t *testing.T) {
	temp := NewTargetTemperatureCharacteristic(20.2, 0, 100, 1, "celsius")
	assert.Equal(t, temp.Permissions, PermsAll())
	assert.Equal(t, temp.Type, CharTypeTemperatureTarget)
}
