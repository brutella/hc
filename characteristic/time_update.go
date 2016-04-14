// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeTimeUpdate = "9A"

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
