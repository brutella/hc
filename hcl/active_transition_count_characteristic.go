package hcl

import (
	"github.com/brutella/hc/characteristic"
)

const TypeActiveTransitionCount = "24B"

type ActiveTransitionCountCharacteristic struct {
	*characteristic.Int
}

func NewActiveTransitionCountCharacteristic() *ActiveTransitionCountCharacteristic {
	char := characteristic.NewInt(TypeActiveTransitionCount)
	char.Format = characteristic.FormatUInt8
	char.Perms = []string{characteristic.PermRead, characteristic.PermEvents}
	char.SetValue(0)

	return &ActiveTransitionCountCharacteristic{char}
}
