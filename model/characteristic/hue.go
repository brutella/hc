package characteristic

type Hue struct {
	*Float
}

func NewHue(value float64) *Hue {
	float := NewFloatMinMaxSteps(value, 0, 360, 1, PermsAll())
	float.Unit = UnitArcDegrees
	float.Type = TypeHue

	return &Hue{float}
}
