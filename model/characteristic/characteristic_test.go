package characteristic

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCharacteristicSetValuesOfWrongType(t *testing.T) {
	var value int = 5
	c := NewCharacteristic(value, FormatInt, CharTypePowerState, nil)

	c.SetValue(float64(20.5))
	assert.Equal(t, c.Value, 20)

	c.SetValue("91")
	assert.Equal(t, c.Value, 91)

	c.SetValue(true)
	assert.Equal(t, c.Value, 1)
}

func TestCharacteristicLocalDelegate(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	var oldValue interface{}
	var newValue interface{}

	c.OnLocalChange(func(c *Characteristic, old interface{}) {
		newValue = c.Value
		oldValue = old
	})

	c.SetValue(10)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
	c.SetValueFromRemote(20)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
}

func TestCharacteristicRemoteDelegate(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	var oldValue interface{}
	var newValue interface{}
	c.OnRemoteChange(func(c *Characteristic, old interface{}) {
		newValue = c.Value
		oldValue = old
	})

	c.SetValueFromRemote(10)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
	c.SetValue(20)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
}

func TestNoValueChange(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	changed := false
	c.OnRemoteChange(func(c *Characteristic, old interface{}) {
		changed = true
	})

	c.OnLocalChange(func(c *Characteristic, old interface{}) {
		changed = true
	})

	c.SetValue(5)
	c.SetValueFromRemote(5)
	assert.False(t, changed)
}

func TestEqual(t *testing.T) {
	c1 := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)
	c2 := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)
	assert.True(t, c1.Equal(c2))
}
