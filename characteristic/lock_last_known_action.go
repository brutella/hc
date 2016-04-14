// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	LockLastKnownActionSecuredPhysicallyInterior   int = 0
	LockLastKnownActionUnsecuredPhysicallyInterior int = 1
	LockLastKnownActionSecuredPhysicallyExterior   int = 2
	LockLastKnownActionUnsecuredPhysicallyExterior int = 3
	LockLastKnownActionSecuredByKeypad             int = 4
	LockLastKnownActionUnsecuredByKeypad           int = 5
	LockLastKnownActionSecuredRemotely             int = 6
	LockLastKnownActionUnsecuredRemotely           int = 7
	LockLastKnownActionSecuredByAutoSecureTimeout  int = 8
)

const TypeLockLastKnownAction = "1C"

type LockLastKnownAction struct {
	*Int
}

func NewLockLastKnownAction() *LockLastKnownAction {
	char := NewInt(TypeLockLastKnownAction)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &LockLastKnownAction{char}
}
