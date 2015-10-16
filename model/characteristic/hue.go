package characteristic

const (
	MinHue  float64 = 0
	MaxHue  float64 = 360
	StepHue float64 = 1
)

type Hue struct {
	*Float
}

func NewHue(value float64) *Hue {
	float := NewFloatMinMaxSteps(value, MinHue, MaxHue, StepHue, PermsAll())
	float.Unit = UnitArcDegrees
	float.Type = TypeHue

	return &Hue{float}
}
