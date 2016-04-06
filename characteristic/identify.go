// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeIdentify = "00000014-0000-1000-8000-0026BB765291"

type Identify struct {
	*Bool
}

func NewIdentify() *Identify {
	char := NewBool(TypeIdentify)
	char.Format = FormatBool
	char.Perms = []string{PermWrite}

	return &Identify{char}
}
