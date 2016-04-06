// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	AirParticulateSize2_5Μm int = 0
	AirParticulateSize10Μm  int = 1
)

const TypeAirParticulateSize = "00000065-0000-1000-8000-0026BB765291"

type AirParticulateSize struct {
	*Int
}

func NewAirParticulateSize() *AirParticulateSize {
	char := NewInt(TypeAirParticulateSize)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &AirParticulateSize{char}
}
