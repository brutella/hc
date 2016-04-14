// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTargetHorizontalTiltAngle = "7B"

type TargetHorizontalTiltAngle struct {
	*Int
}

func NewTargetHorizontalTiltAngle() *TargetHorizontalTiltAngle {
	char := NewInt(TypeTargetHorizontalTiltAngle)
	char.Format = FormatInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(-90)
	char.SetMaxValue(90)
	char.SetStepValue(1)
	char.SetValue(-90)
	char.Unit = UnitArcDegrees

	return &TargetHorizontalTiltAngle{char}
}
