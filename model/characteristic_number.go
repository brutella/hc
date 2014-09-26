package model

type NumberCharacteristic struct {
    *Characteristic
}

func NewNumberCharacteristic(value, min, max, step interface{}, format string) *NumberCharacteristic {    
    return &NumberCharacteristic{NewCharacteristic(value, format, CharTypeUnknown, nil)}
}

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

type FloatCharacteristic struct {
    *NumberCharacteristic
}

func NewFloatCharacteristic(value, min, max, steps float64) *FloatCharacteristic {
    number := NewNumberCharacteristic(value, min, max, steps, FormatFloat)
    return &FloatCharacteristic{number}
}

type IntCharacteristic struct {
    *NumberCharacteristic
}

func NewIntCharacteristic(value, min, max, steps int) *IntCharacteristic {
    number := NewNumberCharacteristic(value, min, max, steps, FormatInt)
    return &IntCharacteristic{number}
}