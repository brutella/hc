// THIS FILE IS AUTO-GENERATED
package characteristic

const TypeDiscoveredBridgedAccessories = "0000009F-0000-1000-8000-0026BB765291"

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
