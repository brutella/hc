// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeModel = "00000021-0000-1000-8000-0026BB765291"

type Model struct {
	*String
}

func NewModel() *Model {
	char := NewString(TypeModel)
	char.Format = FormatString
	char.Perms = []string{PermRead}

	char.SetValue("")

	return &Model{char}
}
