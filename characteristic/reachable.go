// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeReachable = "00000063-0000-1000-8000-0026BB765291"

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
