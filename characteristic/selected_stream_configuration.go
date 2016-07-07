// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSelectedStreamConfiguration = "117"

type SelectedStreamConfiguration struct {
	*Bytes
}

func NewSelectedStreamConfiguration() *SelectedStreamConfiguration {
	char := NewBytes(TypeSelectedStreamConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite}

	char.SetValue([]byte{})

	return &SelectedStreamConfiguration{char}
}
