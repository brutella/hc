// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSaturation = "2F"

type Saturation struct {
	*Float
}

func NewSaturation() *Saturation {
	char := NewFloat(TypeSaturation)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &Saturation{char}
}
