// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentHorizontalTiltAngle = "0000006C-0000-1000-8000-0026BB765291"

type CurrentHorizontalTiltAngle struct {
	*Int
}

func NewCurrentHorizontalTiltAngle() *CurrentHorizontalTiltAngle {
	char := NewInt(TypeCurrentHorizontalTiltAngle)
	char.Format = FormatInt32
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(-90)
	char.SetMaxValue(90)
	char.SetStepValue(1)
	char.SetValue(-90)
	char.Unit = UnitArcDegrees

	return &CurrentHorizontalTiltAngle{char}
}
