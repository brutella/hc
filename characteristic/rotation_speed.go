// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeRotationSpeed = "00000029-0000-1000-8000-0026BB765291"

type RotationSpeed struct {
	*Float
}

func NewRotationSpeed() *RotationSpeed {
	char := NewFloat(TypeRotationSpeed)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &RotationSpeed{char}
}
