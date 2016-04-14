// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeCategory = "A3"

type Category struct {
	*Int
}

func NewCategory() *Category {
	char := NewInt(TypeCategory)
	char.Format = FormatUInt16
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(1)
	char.SetMaxValue(16)
	char.SetStepValue(1)
	char.SetValue(1)

	return &Category{char}
}
