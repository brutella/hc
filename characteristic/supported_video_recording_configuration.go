// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedVideoRecordingConfiguration = "206"

type SupportedVideoRecordingConfiguration struct {
	*Bytes
}

func NewSupportedVideoRecordingConfiguration() *SupportedVideoRecordingConfiguration {
	char := NewBytes(TypeSupportedVideoRecordingConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue([]byte{})

	return &SupportedVideoRecordingConfiguration{char}
}
