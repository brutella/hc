// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedRTPConfiguration = "116"

type SupportedRTPConfiguration struct {
	*Bytes
}

func NewSupportedRTPConfiguration() *SupportedRTPConfiguration {
	char := NewBytes(TypeSupportedRTPConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead}

	char.SetValue([]byte{})

	return &SupportedRTPConfiguration{char}
}
