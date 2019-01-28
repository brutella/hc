// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	InputSourceTypeOther          int = 0
	InputSourceTypeHomeScreen     int = 1
	InputSourceTypeApplication    int = 10
	InputSourceTypeTuner          int = 2
	InputSourceTypeHdmi           int = 3
	InputSourceTypeCompositeVideo int = 4
	InputSourceTypeSVideo         int = 5
	InputSourceTypeComponentVideo int = 6
	InputSourceTypeDvi            int = 7
	InputSourceTypeAirplay        int = 8
	InputSourceTypeUsb            int = 9
)

const TypeInputSourceType = "DB"

type InputSourceType struct {
	*Int
}

func NewInputSourceType() *InputSourceType {
	char := NewInt(TypeInputSourceType)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}
	char.SetMinValue(0)
	char.SetMaxValue(10)
	char.SetStepValue(1)
	char.SetValue(0)

	return &InputSourceType{char}
}
