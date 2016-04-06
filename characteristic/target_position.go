// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTargetPosition = "0000007C-0000-1000-8000-0026BB765291"

type TargetPosition struct {
	*Int
}

func NewTargetPosition() *TargetPosition {
	char := NewInt(TypeTargetPosition)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &TargetPosition{char}
}
