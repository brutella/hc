package characteristic

type Obstruction struct {
	*Bool
}

func NewObstruction() *Identify {
	b := NewBool(false, PermsReadOnly())
	b.Type = TypeObstruction

	return &Identify{b}
}
