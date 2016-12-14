// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePM10Density = "C7"

type PM10Density struct {
	*Float
}

func NewPM10Density() *PM10Density {
	char := NewFloat(TypePM10Density)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &PM10Density{char}
}
