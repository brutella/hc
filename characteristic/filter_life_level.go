// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeFilterLifeLevel = "AB"

type FilterLifeLevel struct {
	*Float
}

func NewFilterLifeLevel() *FilterLifeLevel {
	char := NewFloat(TypeFilterLifeLevel)
	char.Format = FormatFloat
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(100)

	char.SetValue(0)

	return &FilterLifeLevel{char}
}
