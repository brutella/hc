package characteristic

type Bool struct {
    *Number
}

func NewBool(value bool) *Bool {
    number := NewNumber(value, nil, nil, nil, FormatBool, )
    return &Bool{number}
}

func (c *Bool) SetBool(value bool) {
    c.SetValue(value)
}

func (c *Bool) BoolValue() bool {
    return c.GetValue().(bool)
}