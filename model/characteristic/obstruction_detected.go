package characteristic

type ObstructionDetected struct {
	*Bool
}

func NewObstructionDetected() *Identify {
	b := NewBool(false, PermsReadOnly())
	b.Type = TypeObstructionDetected

	return &Identify{b}
}
