package characteristic

import (
	"net"
)

type String struct {
	*Characteristic
}

func NewString(typ string) *String {
	char := NewCharacteristic(typ)
	char.Format = FormatString

	return &String{char}
}

// SetValue sets a value
func (c *String) SetValue(str string) {
	c.UpdateValue(str)
}

// GetValue returns the value as string
func (c *String) GetValue() string {
	return c.Characteristic.GetValue().(string)
}

// OnValueRemoteGet calls fn when the value was read by a client.
func (c *String) OnValueRemoteGet(fn func() string) {
	c.OnValueGet(func() interface{} {
		return fn()
	})
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *String) OnValueRemoteUpdate(fn func(string)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(string))
	})
}
