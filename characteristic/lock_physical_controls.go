// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	LockPhysicalControlsControlLockDisabled int = 0
	LockPhysicalControlsControlLockEnabled  int = 1
)

const TypeLockPhysicalControls = "A7"

type LockPhysicalControls struct {
	*Int
}

func NewLockPhysicalControls() *LockPhysicalControls {
	char := NewInt(TypeLockPhysicalControls)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &LockPhysicalControls{char}
}
