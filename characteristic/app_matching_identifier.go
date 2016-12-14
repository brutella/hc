// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeAppMatchingIdentifier = "A4"

type AppMatchingIdentifier struct {
	*Bytes
}

func NewAppMatchingIdentifier() *AppMatchingIdentifier {
	char := NewBytes(TypeAppMatchingIdentifier)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead}

	char.SetValue([]byte{})

	return &AppMatchingIdentifier{char}
}
