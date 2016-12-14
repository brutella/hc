// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeNitrogenDioxideDensity = "C4"

type NitrogenDioxideDensity struct {
	*Float
}

func NewNitrogenDioxideDensity() *NitrogenDioxideDensity {
	char := NewFloat(TypeNitrogenDioxideDensity)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &NitrogenDioxideDensity{char}
}
