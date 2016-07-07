// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedAudioStreamConfiguration = "115"

type SupportedAudioStreamConfiguration struct {
	*Bytes
}

func NewSupportedAudioStreamConfiguration() *SupportedAudioStreamConfiguration {
	char := NewBytes(TypeSupportedAudioStreamConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead}

	char.SetValue([]byte{})

	return &SupportedAudioStreamConfiguration{char}
}
