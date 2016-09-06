package characteristic

import (
	"net"
)

type UInt16 struct {
	*Characteristic
}

func NewUInt16(typ string) *UInt16 {
	number := NewCharacteristic(typ)
	return &UInt16{number}
}

// SetValue sets a value
func (c *UInt16) SetValue(value uint16) {
	c.UpdateValue(value)
}

func (c *UInt16) SetMinValue(value uint16) {
	c.MinValue = value
}

func (c *UInt16) SetMaxValue(value uint16) {
	c.MaxValue = value
}

func (c *UInt16) SetStepValue(value uint16) {
	c.StepValue = value
}

// GetValue returns the value as int
func (c *UInt16) GetValue() uint16 {
	return c.Value.(uint16)
}

func (c *UInt16) GetMinValue() uint16 {
	return c.MinValue.(uint16)
}

func (c *UInt16) GetMaxValue() uint16 {
	return c.MaxValue.(uint16)
}

func (c *UInt16) GetStepValue() uint16 {
	return c.StepValue.(uint16)
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *UInt16) OnValueRemoteUpdate(fn func(uint16)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(uint16))
	})
}
