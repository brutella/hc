// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeOzoneDensity = "C3"

type OzoneDensity struct {
	*Float
}

func NewOzoneDensity() *OzoneDensity {
	char := NewFloat(TypeOzoneDensity)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &OzoneDensity{char}
}
