package characteristic

type InUse struct {
	*Bool
}

func NewInUse(value bool) *InUse {
	char := NewBool(value, PermsRead())
	char.Type = TypeInUse
	return &InUse{char}
}

func (b *InUse) SetInUse(value bool) {
	b.SetBool(value)
}

func (b *InUse) InUse() bool {
	return b.BoolValue()
}
