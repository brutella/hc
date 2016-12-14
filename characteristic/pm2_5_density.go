// THIS FILE IS AUTO-GENERATED
package characteristic

const TypePM2_5Density = "C6"

type PM2_5Density struct {
	*Float
}

func NewPM2_5Density() *PM2_5Density {
	char := NewFloat(TypePM2_5Density)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1000)
	char.SetStepValue(1)
	char.SetValue(0)

	return &PM2_5Density{char}
}
