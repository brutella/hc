package hcl

import (
	"github.com/brutella/hc/characteristic"
)

const TypeTransitionControl = "143"

type TransitionControlCharacteristic struct {
	*characteristic.Bytes
}

func NewTransitionControlCharacteristic() *TransitionControlCharacteristic {
	char := characteristic.NewBytes(TypeTransitionControl)
	char.Format = characteristic.FormatTLV8
	char.Perms = []string{characteristic.PermRead, characteristic.PermWrite, characteristic.PermWriteResponse}
	char.SetValue([]byte{})

	return &TransitionControlCharacteristic{char}
}
