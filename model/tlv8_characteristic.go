package hk

import(
)

type TLV8Characteristic struct {
    *Characteristic
}

// TODO implement
func NewTLV8Characteristic(tlv []byte) *TLV8Characteristic {
    b := &Characteristic{}
    b.Type = CharTypeUnknown
    
    return &TLV8Characteristic{b}
}