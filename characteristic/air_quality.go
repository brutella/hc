// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	AirQualityUnknown   int = 0
	AirQualityExcellent int = 1
	AirQualityGood      int = 2
	AirQualityFair      int = 3
	AirQualityInferior  int = 4
	AirQualityPoor      int = 5
)

const TypeAirQuality = "95"

type AirQuality struct {
	*Int
}

func NewAirQuality() *AirQuality {
	char := NewInt(TypeAirQuality)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &AirQuality{char}
}
