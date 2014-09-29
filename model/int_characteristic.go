package hk

type IntCharacteristic struct {
    *NumberCharacteristic
}

func NewIntCharacteristic(value, min, max, steps int) *IntCharacteristic {
    number := NewNumberCharacteristic(value, min, max, steps, FormatInt)
    return &IntCharacteristic{number}
}