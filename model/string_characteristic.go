package model

type StringCharacteristic struct {
    Characteristic
}

func NewStringCharacteristic(value string) StringCharacteristic {
    return StringCharacteristic{NewCharacteristic(value, FormatString, CharTypeUnknown, nil)}
}
