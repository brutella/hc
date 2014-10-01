package characteristic

type String struct {
    *Characteristic
}

func NewString(value string) *String {
    return &String{NewCharacteristic(value, FormatString, CharTypeUnknown, nil)}
}

func (c *String) SetString(str string) {
    c.SetValue(str)
}

func (c *String) StringValue() string {
    return c.Value.(string)
}
