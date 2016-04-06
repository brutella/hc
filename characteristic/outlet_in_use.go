// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeOutletInUse = "00000026-0000-1000-8000-0026BB765291"

type OutletInUse struct {
	*Bool
}

func NewOutletInUse() *OutletInUse {
	char := NewBool(TypeOutletInUse)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &OutletInUse{char}
}
