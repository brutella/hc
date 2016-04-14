// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeModel = "21"

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
