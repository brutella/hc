package characteristic

// WifiCapabilities
// this can't be right -- '0000021E-0000-1000-8000-0000022D'
// const TypeWifiCapabilities = "21E"
const TypeWifiCapabilities = "22D"

// I don't know what the various ints mean yet
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
