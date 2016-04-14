// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	RotationDirectionClockwise        int = 0
	RotationDirectionCounterclockwise int = 1
)

const TypeRotationDirection = "28"

type RotationDirection struct {
	*Int
}

func NewRotationDirection() *RotationDirection {
	char := NewInt(TypeRotationDirection)
	char.Format = FormatInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &RotationDirection{char}
}
