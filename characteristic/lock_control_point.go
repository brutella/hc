// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeLockControlPoint = "00000019-0000-1000-8000-0026BB765291"

type LockControlPoint struct {
	*Bytes
}

func NewLockControlPoint() *LockControlPoint {
	char := NewBytes(TypeLockControlPoint)
	char.Format = FormatTLV8
	char.Perms = []string{PermWrite}

	return &LockControlPoint{char}
}
