package characteristic

import (
	"net"
)

type Float struct {
	*Characteristic
}

func NewFloat(typ string) *Float {
	number := NewCharacteristic(typ)
	return &Float{number}
}

// SetValue sets a value
func (c *Float) SetValue(value float64) {
	c.UpdateValue(value)
}

func (c *Float) SetMinValue(value float64) {
	c.MinValue = value
}

func (c *Float) SetMaxValue(value float64) {
	c.MaxValue = value
}

func (c *Float) SetStepValue(value float64) {
	c.StepValue = value
}

// GetValue returns the value as float
func (c *Float) GetValue() float64 {
	return c.Value.(float64)
}

func (c *Float) GetMinValue() float64 {
	return c.MinValue.(float64)
}

func (c *Float) GetMaxValue() float64 {
	return c.MaxValue.(float64)
}

func (c *Float) GetStepValue() float64 {
	return c.StepValue.(float64)
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *Float) OnValueRemoteUpdate(fn func(float64)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(float64))
	})
}
