// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeLockManagementAutoSecurityTimeout = "1A"

type LockManagementAutoSecurityTimeout struct {
	*Int
}

func NewLockManagementAutoSecurityTimeout() *LockManagementAutoSecurityTimeout {
	char := NewInt(TypeLockManagementAutoSecurityTimeout)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)
	char.Unit = UnitSeconds

	return &LockManagementAutoSecurityTimeout{char}
}
