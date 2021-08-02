package hcl

import (
	"github.com/brutella/hc/characteristic"
)

const TypeSupportedTransitionConfiguration = "144"

type SupportedTransitionConfigurationCharacteristic struct {
	*characteristic.Bytes
}

func NewSupportedTransitionConfigurationCharacteristic() *SupportedTransitionConfigurationCharacteristic {
	char := characteristic.NewBytes(TypeSupportedTransitionConfiguration)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead}
	char.SetValue([]byte{})

	return &SupportedTransitionConfigurationCharacteristic{char}
}
