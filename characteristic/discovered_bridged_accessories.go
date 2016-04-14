// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDiscoveredBridgedAccessories = "9F"

type DiscoveredBridgedAccessories struct {
	*Int
}

func NewDiscoveredBridgedAccessories() *DiscoveredBridgedAccessories {
	char := NewInt(TypeDiscoveredBridgedAccessories)
	char.Format = FormatUInt16
	char.Perms = []string{PermRead, PermEvents}

	char.SetValue(0)

	return &DiscoveredBridgedAccessories{char}
}
