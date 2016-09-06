package characteristic

import (
	"net"
)

type UInt8 struct {
	*Characteristic
}

func NewUInt8(typ string) *UInt8 {
	number := NewCharacteristic(typ)
	return &UInt8{number}
}

// SetValue sets a value
func (c *UInt8) SetValue(value uint8) {
	c.UpdateValue(value)
}

func (c *UInt8) SetMinValue(value uint8) {
	c.MinValue = value
}

func (c *UInt8) SetMaxValue(value uint8) {
	c.MaxValue = value
}

func (c *UInt8) SetStepValue(value uint8) {
	c.StepValue = value
}

// GetValue returns the value as int
func (c *UInt8) GetValue() uint8 {
	return c.Value.(uint8)
}

func (c *UInt8) GetMinValue() uint8 {
	return c.MinValue.(uint8)
}

func (c *UInt8) GetMaxValue() uint8 {
	return c.MaxValue.(uint8)
}

func (c *UInt8) GetStepValue() uint8 {
	return c.StepValue.(uint8)
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *UInt8) OnValueRemoteUpdate(fn func(uint8)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(uint8))
	})
}
