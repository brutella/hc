// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SecuritySystemTargetStateStayArm  int = 0
	SecuritySystemTargetStateAwayArm  int = 1
	SecuritySystemTargetStateNightArm int = 2
	SecuritySystemTargetStateDisarm   int = 3
)

const TypeSecuritySystemTargetState = "67"

type SecuritySystemTargetState struct {
	*Int
}

func NewSecuritySystemTargetState() *SecuritySystemTargetState {
	char := NewInt(TypeSecuritySystemTargetState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &SecuritySystemTargetState{char}
}
