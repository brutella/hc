// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeStatusActive = "75"

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
