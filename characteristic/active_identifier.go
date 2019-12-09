// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeActiveIdentifier = "E7"

type ActiveIdentifier struct {
	*Int
}

func NewActiveIdentifier() *ActiveIdentifier {
	char := NewInt(TypeActiveIdentifier)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}
	char.SetMinValue(0)

	char.SetValue(0)

	return &ActiveIdentifier{char}
}
