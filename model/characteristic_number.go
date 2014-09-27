package model

type NumberCharacteristic struct {
    *Characteristic
}

func NewNumberCharacteristic(value, min, max, step interface{}, format string) *NumberCharacteristic {    
    c := NumberCharacteristic{NewCharacteristic(value, format, CharTypeUnknown, nil)}
    c.MinValue = min
    c.MaxValue = max
    c.MinStep = step
    
    return &c
}

type BoolCharacteristic struct {
    *NumberCharacteristic
}

func NewBoolCharacteristic(value bool) *BoolCharacteristic {
    number := NewNumberCharacteristic(value, nil, nil, nil, FormatBool, )
    return &BoolCharacteristic{number}
}

type FloatCharacteristic struct {
    *NumberCharacteristic
}

func NewFloatCharacteristic(value float64) *FloatCharacteristic {
    number := NewNumberCharacteristic(value, nil, nil, nil, FormatFloat)
    return &FloatCharacteristic{number}
}

func NewFloatCharacteristicMinMaxSteps(value, min, max, steps float64) *FloatCharacteristic {
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

type ByteCharacteristic struct {
    *NumberCharacteristic
}

func NewByteCharacteristic(value byte) *ByteCharacteristic {
    number := NewNumberCharacteristic(value, nil, nil, nil, FormatByte)
    return &ByteCharacteristic{number}
}

func NewByteCharacteristicMinMaxSteps(value, min, max, steps byte) *ByteCharacteristic {
    number := NewNumberCharacteristic(value, min, max, steps, FormatByte)
    return &ByteCharacteristic{number}
}