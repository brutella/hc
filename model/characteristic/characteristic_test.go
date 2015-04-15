package characteristic

import (
	"github.com/stretchr/testify/assert"
	"net"
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

	c.OnChange(func(c *Characteristic, new, old interface{}) {
		newValue = new
		oldValue = old
	})

	c.SetValue(10)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
	c.SetValueFromConnection(20, TestConn)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
}

func TestCharacteristicRemoteDelegate(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	var oldValue interface{}
	var newValue interface{}
	c.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		assert.Equal(t, conn, TestConn)
		newValue = new
		oldValue = old
	})

	c.SetValueFromConnection(10, TestConn)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
	c.SetValue(20)
	assert.Equal(t, oldValue, 5)
	assert.Equal(t, newValue, 10)
}

func TestNoValueChange(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)

	changed := false
	c.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		assert.Equal(t, conn, TestConn)
		changed = true
	})

	c.OnChange(func(c *Characteristic, new, old interface{}) {
		changed = true
	})

	c.SetValue(5)
	c.SetValueFromConnection(5, TestConn)
	assert.False(t, changed)
}

func TestReadOnlyValue(t *testing.T) {
	c := NewCharacteristic(5, FormatInt, CharTypePowerState, PermsRead())

	remoteChanged := false
	localChanged := false
	c.OnConnChange(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		assert.Equal(t, conn, TestConn)
		remoteChanged = true
	})

	c.OnChange(func(c *Characteristic, new, old interface{}) {
		localChanged = true
	})

	c.SetValue(10)
	c.SetValueFromConnection(11, TestConn)

	assert.Equal(t, c.GetValue(), 10)
	assert.False(t, remoteChanged)
	assert.True(t, localChanged)
}

func TestEqual(t *testing.T) {
	c1 := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)
	c2 := NewCharacteristic(5, FormatInt, CharTypePowerState, nil)
	assert.True(t, c1.Equal(c2))
}
