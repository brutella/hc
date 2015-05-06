package characteristic

import (
	"bytes"
	"encoding/base64"
	"github.com/brutella/hc/util"
)

type Bytes struct {
	*String
}

func NewBytes(b []byte) *Bytes {
	s := NewString(tlv8Base64FromBytes(b))
	s.Type = CharTypeUnknown
	s.Format = FormatTLV8

	return &Bytes{s}
}

func (bs *Bytes) SetBytes(b []byte) {
	bs.SetString(tlv8Base64FromBytes(b))
}

func (bs *Bytes) Bytes() []byte {
	b, _ := bytesFromTLV8Base64(bs.StringValue())
	return b
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
