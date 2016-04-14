// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeReachable = "63"

type Reachable struct {
	*Bool
}

func NewReachable() *Reachable {
	char := NewBool(TypeReachable)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &Reachable{char}
}
