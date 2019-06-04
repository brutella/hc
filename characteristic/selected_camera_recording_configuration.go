// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSelectedCameraRecordingConfiguration = "209"

type SelectedCameraRecordingConfiguration struct {
	*Bytes
}

func NewSelectedCameraRecordingConfiguration() *SelectedCameraRecordingConfiguration {
	char := NewBytes(TypeSelectedCameraRecordingConfiguration)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue([]byte{})

	return &SelectedCameraRecordingConfiguration{char}
}
