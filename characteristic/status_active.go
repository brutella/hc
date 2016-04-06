// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeStatusActive = "00000075-0000-1000-8000-0026BB765291"

type StatusActive struct {
	*Bool
}

func NewStatusActive() *StatusActive {
	char := NewBool(TypeStatusActive)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &StatusActive{char}
}
