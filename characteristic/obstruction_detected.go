// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeObstructionDetected = "24"

type ObstructionDetected struct {
	*Bool
}

func NewObstructionDetected() *ObstructionDetected {
	char := NewBool(TypeObstructionDetected)
	char.Format = FormatBool
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(false)

	return &ObstructionDetected{char}
}
