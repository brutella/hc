package characteristic

type IntCharacteristic struct {
    *Number
}

func NewIntCharacteristic(value, min, max, steps int) *IntCharacteristic {
    number := NewNumber(value, min, max, steps, FormatInt)
    return &IntCharacteristic{number}
}