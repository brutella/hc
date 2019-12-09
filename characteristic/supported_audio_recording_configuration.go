// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedAudioRecordingConfiguration = "207"

type SupportedAudioRecordingConfiguration struct {
	*Bytes
}

func NewSupportedAudioRecordingConfiguration() *SupportedAudioRecordingConfiguration {
	char := NewBytes(TypeSupportedAudioRecordingConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue([]byte{})

	return &SupportedAudioRecordingConfiguration{char}
}
