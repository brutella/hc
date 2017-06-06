// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeVolume = "119"

type Volume struct {
	*Int
}

func NewVolume() *Volume {
	char := NewInt(TypeVolume)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &Volume{char}
}
