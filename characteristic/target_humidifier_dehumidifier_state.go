// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	TargetHumidifierDehumidifierStateHumidifierOrDehumidifier int = 0
	TargetHumidifierDehumidifierStateHumidifier               int = 1
	TargetHumidifierDehumidifierStateDehumidifier             int = 2
)

const TypeTargetHumidifierDehumidifierState = "B4"

type TargetHumidifierDehumidifierState struct {
	*Int
}

func NewTargetHumidifierDehumidifierState() *TargetHumidifierDehumidifierState {
	char := NewInt(TypeTargetHumidifierDehumidifierState)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &TargetHumidifierDehumidifierState{char}
}
