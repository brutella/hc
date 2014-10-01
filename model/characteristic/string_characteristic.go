package characteristic

type StringCharacteristic struct {
    *Characteristic
}

func NewStringCharacteristic(value string) *StringCharacteristic {
    return &StringCharacteristic{NewCharacteristic(value, FormatString, CharTypeUnknown, nil)}
}

func (c *StringCharacteristic) SetString(str string) {
    c.SetValue(str)
}

func (c *StringCharacteristic) String() string {
    return c.Value.(string)
}
