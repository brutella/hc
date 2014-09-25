package model

type BoolCharacteristic struct {
    *NumberCharacteristic
}

func NewBoolCharacteristic(value bool) *BoolCharacteristic {
    integer := 0
    if value == true {
        integer = 1
    }
    number := NewNumberCharacteristic(integer, 0, 0, 0, FormatBool, )
    return &BoolCharacteristic{number}
}