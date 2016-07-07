// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedVideoStreamConfiguration = "114"

type SupportedVideoStreamConfiguration struct {
	*Bytes
}

func NewSupportedVideoStreamConfiguration() *SupportedVideoStreamConfiguration {
	char := NewBytes(TypeSupportedVideoStreamConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead}

	char.SetValue([]byte{})

	return &SupportedVideoStreamConfiguration{char}
}
