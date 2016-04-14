// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeOutletInUse = "26"

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
