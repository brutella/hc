// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeImageRotation = "11E"

type ImageRotation struct {
	*Float
}

func NewImageRotation() *ImageRotation {
	char := NewFloat(TypeImageRotation)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(270)
	char.SetStepValue(90)
	char.SetValue(0)
	char.Unit = UnitArcDegrees

	return &ImageRotation{char}
}
