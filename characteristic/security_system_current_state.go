// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	SecuritySystemCurrentStateStayArm        int = 0
	SecuritySystemCurrentStateAwayArm        int = 1
	SecuritySystemCurrentStateNightArm       int = 2
	SecuritySystemCurrentStateDisarmed       int = 3
	SecuritySystemCurrentStateAlarmTriggered int = 4
)

const TypeSecuritySystemCurrentState = "66"

type SecuritySystemCurrentState struct {
	*Int
}

func NewSecuritySystemCurrentState() *SecuritySystemCurrentState {
	char := NewInt(TypeSecuritySystemCurrentState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &SecuritySystemCurrentState{char}
}
