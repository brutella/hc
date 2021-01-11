package characteristic

// WifiConfigurationControl
// this can't be right -- '0000021E-0000-1000-8000-0000022D'
// const TypeWifiCapabilities = "21E"
const TypeWifiConfigurationControl = "22D"

type WifiConfigurationControl = struct {
	*Bytes
}

func NewWifiConfigurationControl() *WifiConfigurationControl {
	char := NewBytes(TypeWifiCapabilities)
	char.Format = FormatTLV8
	char.Perms = []string{PermRead, PermWrite, PermEvents}

	return &WifiConfigurationControl{char}
}
