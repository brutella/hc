package characteristic

import (
	"encoding/base64"
	"net"
)

type Bytes struct {
	*String
}

func NewBytes(typ string) *Bytes {
	s := NewString(typ)
	s.Format = FormatTLV8

	return &Bytes{s}
}

func (bs *Bytes) SetValue(b []byte) {
	bs.String.SetValue(base64FromBytes(b))
}

func (bs *Bytes) GetValue() []byte {
	str := bs.String.GetValue()
	if b, err := base64.StdEncoding.DecodeString(str); err != nil {
		return []byte{}
	} else {
		return b
	}
}

// OnValueRemoteUpdate calls fn when the value was updated by a client.
func (bs *Bytes) OnValueRemoteUpdate(fn func([]byte)) {
	bs.OnValueUpdateFromConn(func(conn net.Conn, c *Characteristic, new, old interface{}) {
		fn(new.([]byte))
	})
}

func base64FromBytes(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}
