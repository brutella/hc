// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTargetVerticalTiltAngle = "0000007D-0000-1000-8000-0026BB765291"

type TargetVerticalTiltAngle struct {
	*Int
}

func NewTargetVerticalTiltAngle() *TargetVerticalTiltAngle {
	char := NewInt(TypeTargetVerticalTiltAngle)
	char.Format = FormatInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(-90)
	char.SetMaxValue(90)
	char.SetStepValue(1)
	char.SetValue(-90)
	char.Unit = UnitArcDegrees

	return &TargetVerticalTiltAngle{char}
}
