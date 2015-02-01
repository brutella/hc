package characteristic

type InUse struct {
	*Bool
}

func NewInUse(value bool) *InUse {
	char := NewBool(value)
	char.Type = CharTypeInUse
	char.Permissions = PermsRead()
	return &InUse{char}
}

func (b *InUse) SetInUse(value bool) {
	b.SetBool(value)
}

func (b *InUse) InUse() bool {
	return b.BoolValue()
}
