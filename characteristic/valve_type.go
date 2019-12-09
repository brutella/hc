// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	ValveTypeGenericValve int = 0
	ValveTypeIrrigation   int = 1
	ValveTypeShowerHead   int = 2
	ValveTypeWaterFaucet  int = 3
)

const TypeValveType = "D5"

type ValveType struct {
	*Int
}

func NewValveType() *ValveType {
	char := NewInt(TypeValveType)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &ValveType{char}
}
