package characteristic

const (
	MinBrightness  int = 0
	MaxBrightness  int = 100
	StepBrightness int = 1
)

type Brightness struct {
	*Int
}

func NewBrightness(value int) *Brightness {
	i := NewInt(value, MinBrightness, MaxBrightness, StepBrightness, PermsAll())
	i.Unit = UnitPercentage
	i.Type = TypeBrightness

	return &Brightness{i}
}
