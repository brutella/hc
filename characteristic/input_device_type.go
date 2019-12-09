// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	InputDeviceTypeOther       int = 0
	InputDeviceTypeTv          int = 1
	InputDeviceTypeRecording   int = 2
	InputDeviceTypeTuner       int = 3
	InputDeviceTypePlayback    int = 4
	InputDeviceTypeAudioSystem int = 5
)

const TypeInputDeviceType = "DC"

type InputDeviceType struct {
	*Int
}

func NewInputDeviceType() *InputDeviceType {
	char := NewInt(TypeInputDeviceType)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(5)
	char.SetStepValue(1)
	char.SetValue(0)

	return &InputDeviceType{char}
}
