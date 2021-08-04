package characteristic

import (
	"net"
)

type Int struct {
	*Characteristic
}

func NewInt(typ string) *Int {
	number := NewCharacteristic(typ)
	return &Int{number}
}

// SetValue sets a value
func (c *Int) SetValue(value int) {
	c.UpdateValue(value)
}

func (c *Int) SetMinValue(value int) {
	c.MinValue = value
}

func (c *Int) SetMaxValue(value int) {
	c.MaxValue = value
}

func (c *Int) SetStepValue(value int) {
	c.StepValue = value
}

func (c *Int) SetValidValues(values []int) {
	c.ValidValues = values
}

func (c *Int) SetValidValuesRange(start int, end int) {
	c.ValidValuesRange = []int{start, end}
}

// GetValue returns the value as int
func (c *Int) GetValue() int {
	return c.Characteristic.GetValue().(int)
}

func (c *Int) GetMinValue() int {
	return c.MinValue.(int)
}

func (c *Int) GetMaxValue() int {
	return c.MaxValue.(int)
}

func (c *Int) GetStepValue() int {
	return c.StepValue.(int)
}

func (c *Int) GetValidValues() []int {
	return c.ValidValues.([]int)
}

func (c *Int) GetValidValuesRange() (int, int) {
	return c.ValidValues.([]int)[0], c.ValidValues.([]int)[1]
}

// OnValueRemoteGet calls fn when the value was read by a client.
func (c *Int) OnValueRemoteGet(fn func() int) {
	c.OnValueGet(func() interface{} {
		return fn()
	})
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *Int) OnValueRemoteUpdate(fn func(int)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(int))
	})
}
