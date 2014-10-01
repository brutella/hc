package characteristic

type On struct {
    *Bool
}

func NewOn(value bool) *On {
    char := NewBool(value)
    char.Type = CharTypeOn
    return &On{char}
}

func (c *On) SetOn(value bool) {
    c.SetBool(value)
}

func (c *On) On() bool {
    return c.BoolValue()
}