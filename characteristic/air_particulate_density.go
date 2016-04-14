// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAirParticulateDensity = "64"

type AirParticulateDensity struct {
	*Float
}

func NewAirParticulateDensity() *AirParticulateDensity {
	char := NewFloat(TypeAirParticulateDensity)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &AirParticulateDensity{char}
}
