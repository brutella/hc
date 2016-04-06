// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCurrentPosition = "0000006D-0000-1000-8000-0026BB765291"

type CurrentPosition struct {
	*Int
}

func NewCurrentPosition() *CurrentPosition {
	char := NewInt(TypeCurrentPosition)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &CurrentPosition{char}
}
