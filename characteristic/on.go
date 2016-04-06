// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeOn = "00000025-0000-1000-8000-0026BB765291"

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
