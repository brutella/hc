package characteristic

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
