// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeName = "00000023-0000-1000-8000-0026BB765291"

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
