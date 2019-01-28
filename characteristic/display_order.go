// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDisplayOrder = "136"

type DisplayOrder struct {
	*Bytes
}

func NewDisplayOrder() *DisplayOrder {
	char := NewBytes(TypeDisplayOrder)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue([]byte{})

	return &DisplayOrder{char}
}
