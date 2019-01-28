// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeIdentifier = "E6"

type Identifier struct {
	*Int
}

func NewIdentifier() *Identifier {
	char := NewInt(TypeIdentifier)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)

	char.SetStepValue(1)
	char.SetValue(0)

	return &Identifier{char}
}
