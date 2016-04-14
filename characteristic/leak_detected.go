// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	LeakDetectedLeakNotDetected int = 0
	LeakDetectedLeakDetected    int = 1
)

const TypeLeakDetected = "70"

type LeakDetected struct {
	*Int
}

func NewLeakDetected() *LeakDetected {
	char := NewInt(TypeLeakDetected)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &LeakDetected{char}
}
