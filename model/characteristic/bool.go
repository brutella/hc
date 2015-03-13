package characteristic

type Bool struct {
	*Number
}

func NewBool(value bool, permissions []string) *Bool {
	number := NewNumber(value, nil, nil, nil, FormatBool, permissions)
	return &Bool{number}
}

func (c *Bool) SetBool(value bool) {
	c.SetNumber(value)
}

func (c *Bool) BoolValue() bool {
	return c.GetValue().(bool)
}
