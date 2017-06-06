// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonMonoxidePeakLevel = "91"

type CarbonMonoxidePeakLevel struct {
	*Float
}

func NewCarbonMonoxidePeakLevel() *CarbonMonoxidePeakLevel {
	char := NewFloat(TypeCarbonMonoxidePeakLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)

	char.SetValue(0)

	return &CarbonMonoxidePeakLevel{char}
}
