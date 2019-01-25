// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeWaterLevel = "B5"

type WaterLevel struct {
	*Float
}

func NewWaterLevel() *WaterLevel {
	char := NewFloat(TypeWaterLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)

	char.SetValue(0)
	char.Unit = UnitPercentage

	return &WaterLevel{char}
}
