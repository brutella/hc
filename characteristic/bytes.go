package characteristic

import (
	"bytes"
	"encoding/base64"
	"github.com/brutella/hc/util"
)

type Bytes struct {
	*Characteristic
}

func NewBytes(typ string) *Bytes {
	s := NewCharacteristic(typ)

	return &Bytes{s}
}

func (bs *Bytes) SetValue(b []byte) {
	bs.UpdateValue(tlv8Base64FromBytes(b))
}

func (bs *Bytes) GetValue() []byte {
	if str, ok := bs.Value.(string); ok == true {
		b, _ := bytesFromTLV8Base64(str)
		return b
	}

	return []byte{}
}

func tlv8FromBytes(b []byte) []byte {
	c := util.NewTLV8Container()
	c.SetBytes(0x00, b)
	return c.BytesBuffer().Bytes()
}

func bytesFromTLV8(b []byte) ([]byte, error) {
	buf := bytes.NewBuffer(b)
	c, err := util.NewTLV8ContainerFromReader(buf)
	if err != nil {
		return nil, err
	}

	return c.GetBytes(0x00), nil
}

func tlv8Base64FromBytes(b []byte) string {
	return base64.StdEncoding.EncodeToString(tlv8FromBytes(b))
}

func bytesFromTLV8Base64(b64 string) ([]byte, error) {
	b, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return nil, err
	}

	return bytesFromTLV8(b)
}
