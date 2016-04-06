// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTargetRelativeHumidity = "00000034-0000-1000-8000-0026BB765291"

type TargetRelativeHumidity struct {
	*Float
}

func NewTargetRelativeHumidity() *TargetRelativeHumidity {
	char := NewFloat(TypeTargetRelativeHumidity)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &TargetRelativeHumidity{char}
}
