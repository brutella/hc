// THIS FILE IS AUTO-GENERATED
package characteristic

const (
	DiscoverBridgedAccessoriesStartDiscovery int = 0
	DiscoverBridgedAccessoriesStopDiscovery  int = 1
)

const TypeDiscoverBridgedAccessories = "9E"

type DiscoverBridgedAccessories struct {
	*Int
}

func NewDiscoverBridgedAccessories() *DiscoverBridgedAccessories {
	char := NewInt(TypeDiscoverBridgedAccessories)
	char.Format = FormatUInt8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	char.SetValue(0)

	return &DiscoverBridgedAccessories{char}
}
