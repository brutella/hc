// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetAirQualityExcellent int = 0
	TargetAirQualityGood      int = 1
	TargetAirQualityFair      int = 2
)

const TypeTargetAirQuality = "AE"

type TargetAirQuality struct {
	*Int
}

func NewTargetAirQuality() *TargetAirQuality {
	char := NewInt(TypeTargetAirQuality)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetAirQuality{char}
}
