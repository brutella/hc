package characteristic

import (
	"net"
)

type String struct {
	*Characteristic
}

func NewString(typ string) *String {
	return &String{NewCharacteristic(typ)}
}

// SetValue sets a value
func (c *String) SetValue(str string) {
	c.UpdateValue(str)
}

// GetValue returns the value as string
func (c *String) GetValue() string {
	return c.Value.(string)
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (c *String) OnValueRemoteUpdate(fn func(string)) {
	c.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.(string))
	})
}
