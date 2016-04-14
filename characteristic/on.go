// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeOn = "25"

type On struct {
	*Bool
}

func NewOn() *On {
	char := NewBool(TypeOn)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(false)

	return &On{char}
}
