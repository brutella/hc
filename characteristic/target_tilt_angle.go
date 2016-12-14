// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTargetTiltAngle = "C2"

type TargetTiltAngle struct {
	*Int
}

func NewTargetTiltAngle() *TargetTiltAngle {
	char := NewInt(TypeTargetTiltAngle)
	char.Format = FormatInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(-90)
	char.SetMaxValue(90)
	char.SetStepValue(1)
	char.SetValue(-90)
	char.Unit = UnitArcDegrees

	return &TargetTiltAngle{char}
}
