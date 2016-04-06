// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeMotionDetected = "00000022-0000-1000-8000-0026BB765291"

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
