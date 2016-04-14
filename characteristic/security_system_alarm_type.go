// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSecuritySystemAlarmType = "8E"

type SecuritySystemAlarmType struct {
	*Int
}

func NewSecuritySystemAlarmType() *SecuritySystemAlarmType {
	char := NewInt(TypeSecuritySystemAlarmType)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(1)
	char.SetStepValue(1)
	char.SetValue(0)

	return &SecuritySystemAlarmType{char}
}
