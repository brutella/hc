// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeObstructionDetected = "00000024-0000-1000-8000-0026BB765291"

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
