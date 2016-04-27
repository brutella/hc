package characteristic

import (
	"net"
)

type Bool struct {
	*Characteristic
}

func NewBool(typ string) *Bool {
	number := NewCharacteristic(typ)
	number.Format = FormatBool

	return &Bool{number}
}

// SetValue sets a value
func (c *Bool) SetValue(value bool) {
	c.UpdateValue(value)
}

// GetValue returns the value as bool
func (c *Bool) GetValue() bool {
	return c.Value.(bool)
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *Bool) OnValueRemoteUpdate(fn func(bool)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(bool))
	})
}
