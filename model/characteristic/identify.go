package characteristic

type Identify struct {
	*Bool
}

func NewIdentify(identify bool) *Identify {
	b := NewBool(identify)
	b.Type = CharTypeIdentify
	b.Permissions = PermsWriteOnly()

	return &Identify{b}
}

func (c *Identify) SetIdentify(identify bool) {
	c.SetBool(identify)
}

func (c *Identify) Identify() bool {
	return c.BoolValue()
}
