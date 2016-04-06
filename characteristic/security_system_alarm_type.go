// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeSecuritySystemAlarmType = "0000008E-0000-1000-8000-0026BB765291"

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
