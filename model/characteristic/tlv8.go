package characteristic

import ()

type TLV8 struct {
	*Characteristic
}

// TODO implement
func NewTLV8(tlv []byte) *TLV8 {
	b := &Characteristic{}
	b.Type = CharTypeUnknown

	return &TLV8{b}
}
