// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeVOCDensity = "C8"

type VOCDensity struct {
	*Float
}

func NewVOCDensity() *VOCDensity {
	char := NewFloat(TypeVOCDensity)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &VOCDensity{char}
}
