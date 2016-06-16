// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeVolume = "119"

type Volume struct {
	*Float
}

func NewVolume() *Volume {
	char := NewFloat(TypeVolume)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &Volume{char}
}
