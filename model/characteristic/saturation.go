package characteristic

type Saturation struct {
	*Float
}

func NewSaturation(value float64) *Saturation {
	float := NewFloatMinMaxSteps(value, 0, 100, 1, PermsAll())
	float.Unit = UnitPercentage
	float.Type = CharTypeSaturation

	return &Saturation{float}
}
