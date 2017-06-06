// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonMonoxideLevel = "90"

type CarbonMonoxideLevel struct {
	*Float
}

func NewCarbonMonoxideLevel() *CarbonMonoxideLevel {
	char := NewFloat(TypeCarbonMonoxideLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)

	char.SetValue(0)

	return &CarbonMonoxideLevel{char}
}
