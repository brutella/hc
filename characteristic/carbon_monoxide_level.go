// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCarbonMonoxideLevel = "00000090-0000-1000-8000-0026BB765291"

type CarbonMonoxideLevel struct {
	*Float
}

func NewCarbonMonoxideLevel() *CarbonMonoxideLevel {
	char := NewFloat(TypeCarbonMonoxideLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)
	char.SetStepValue(0.1)
	char.SetValue(0)

	return &CarbonMonoxideLevel{char}
}
