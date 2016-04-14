// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeName = "23"

type Name struct {
	*String
}

func NewName() *Name {
	char := NewString(TypeName)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &Name{char}
}
