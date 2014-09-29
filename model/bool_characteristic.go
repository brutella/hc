package hk

type BoolCharacteristic struct {
    *NumberCharacteristic
}

func NewBoolCharacteristic(value bool) *BoolCharacteristic {
    number := NewNumberCharacteristic(value, nil, nil, nil, FormatBool, )
    return &BoolCharacteristic{number}
}

func (c *BoolCharacteristic) SetBool(value bool) {
    c.SetValue(value)
}

func (c *BoolCharacteristic) Bool() bool {
    return c.GetValue().(bool)
}