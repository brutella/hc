// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	LockTargetStateUnsecured int = 0
	LockTargetStateSecured   int = 1
)

const TypeLockTargetState = "0000001E-0000-1000-8000-0026BB765291"

type LockTargetState struct {
	*Int
}

func NewLockTargetState() *LockTargetState {
	char := NewInt(TypeLockTargetState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &LockTargetState{char}
}
