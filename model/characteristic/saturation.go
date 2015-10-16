package characteristic

const (
	MinSaturation  float64 = 0
	MaxSaturation  float64 = 100
	StepSaturation float64 = 1
)

type Saturation struct {
	*Float
}

func NewSaturation(value float64) *Saturation {
	float := NewFloatMinMaxSteps(value, MinSaturation, MaxSaturation, StepSaturation, PermsAll())
	float.Unit = UnitPercentage
	float.Type = TypeSaturation

	return &Saturation{float}
}
