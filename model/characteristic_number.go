package model

type NumberCharacteristic struct {
    *Characteristic
}

func NewNumberCharacteristic(value, min, max, step interface{}, format string) *NumberCharacteristic {    
    return &NumberCharacteristic{NewCharacteristic(value, format, CharTypeUnknown, nil)}
}