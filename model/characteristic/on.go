package characteristic

type On struct {
	*Bool
}

func NewOn(value bool) *On {
	char := NewBool(value)
	char.Type = CharTypePowerState
	return &On{char}
}

func (b *On) SetOn(value bool) {
	b.SetBool(value)
}

func (b *On) On() bool {
	return b.BoolValue()
}
