// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSupportedCameraRecordingConfiguration = "205"

type SupportedCameraRecordingConfiguration struct {
	*Bytes
}

func NewSupportedCameraRecordingConfiguration() *SupportedCameraRecordingConfiguration {
	char := NewBytes(TypeSupportedCameraRecordingConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue([]byte{})

	return &SupportedCameraRecordingConfiguration{char}
}
