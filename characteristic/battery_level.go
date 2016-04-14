// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeBatteryLevel = "68"

type BatteryLevel struct {
	*Int
}

func NewBatteryLevel() *BatteryLevel {
	char := NewInt(TypeBatteryLevel)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(1)
	char.SetValue(0)
	char.Unit = UnitPercentage

	return &BatteryLevel{char}
}
