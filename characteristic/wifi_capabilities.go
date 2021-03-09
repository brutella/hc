package characteristic

const TypeWifiCapabilities = "22C"

type WifiCapabilities = struct {
	*Int
}

func NewWifiCapabilities() *WifiCapabilities {
	char := NewInt(TypeWifiCapabilities)
	char.Format = FormatUInt32
	char.Perms = []string{PermRead}

	char.SetValue(1)
	return &WifiCapabilities{char}
}
