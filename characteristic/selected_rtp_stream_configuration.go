// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSelectedRTPStreamConfiguration = "117"

type SelectedRTPStreamConfiguration struct {
	*Bytes
}

func NewSelectedRTPStreamConfiguration() *SelectedRTPStreamConfiguration {
	char := NewBytes(TypeSelectedRTPStreamConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite}

	char.SetValue([]byte{})

	return &SelectedRTPStreamConfiguration{char}
}
