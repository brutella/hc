package characteristic

type Brightness struct {
	*Int
}

func NewBrightness(value int) *Brightness {
	i := NewInt(value, 0, 100, 1, PermsAll())
	i.Unit = UnitPercentage
	i.Type = CharTypeBrightness

	return &Brightness{i}
}
