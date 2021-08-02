package hcl

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSupportedDataStreamTransportConfiguration = "130"

type SupportedDataStreamTransportConfigurationCharacteristic struct {
	*characteristic.Bytes
}

func NewSupportedDataStreamTransportConfigurationCharacteristic() *SupportedDataStreamTransportConfigurationCharacteristic {
	char := characteristic.NewBytes(TypeSupportedDataStreamTransportConfiguration)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead}
	char.SetValue([]byte{})

	return &SupportedDataStreamTransportConfigurationCharacteristic{char}
}
