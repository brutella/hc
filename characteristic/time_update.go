// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTimeUpdate = "0000009A-0000-1000-8000-0026BB765291"

type TimeUpdate struct {
	*Bool
}

func NewTimeUpdate() *TimeUpdate {
	char := NewBool(TypeTimeUpdate)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &TimeUpdate{char}
}
