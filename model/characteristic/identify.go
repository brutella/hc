package characteristic

type Identify struct {
	*Bool
}

func NewIdentify() *Identify {
	b := NewBool(false, PermsWriteOnly())
	b.Type = TypeIdentify

	return &Identify{b}
}
