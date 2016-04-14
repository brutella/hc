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
	char.SetMinValue(0)
	char.SetMaxValue(86400)
	char.SetStepValue(1)
	char.SetValue(0)

	return &LockManagementAutoSecurityTimeout{char}
}
