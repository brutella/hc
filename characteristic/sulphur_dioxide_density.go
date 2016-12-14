// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSulphurDioxideDensity = "C5"

type SulphurDioxideDensity struct {
	*Float
}

func NewSulphurDioxideDensity() *SulphurDioxideDensity {
	char := NewFloat(TypeSulphurDioxideDensity)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &SulphurDioxideDensity{char}
}
