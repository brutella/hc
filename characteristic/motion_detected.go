// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeMotionDetected = "22"

type MotionDetected struct {
	*Bool
}

func NewMotionDetected() *MotionDetected {
	char := NewBool(TypeMotionDetected)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &MotionDetected{char}
}
