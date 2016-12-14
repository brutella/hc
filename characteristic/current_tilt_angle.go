// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentTiltAngle = "C1"

type CurrentTiltAngle struct {
	*Int
}

func NewCurrentTiltAngle() *CurrentTiltAngle {
	char := NewInt(TypeCurrentTiltAngle)
	char.Format = FormatInt32
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(-90)
	char.SetMaxValue(90)
	char.SetStepValue(1)
	char.SetValue(-90)
	char.Unit = UnitArcDegrees

	return &CurrentTiltAngle{char}
}
