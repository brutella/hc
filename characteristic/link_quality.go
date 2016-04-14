// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeLinkQuality = "9C"

type LinkQuality struct {
	*Int
}

func NewLinkQuality() *LinkQuality {
	char := NewInt(TypeLinkQuality)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(1)
	char.SetMaxValue(4)
	char.SetStepValue(1)
	char.SetValue(1)

	return &LinkQuality{char}
}
