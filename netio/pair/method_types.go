package pair

import "fmt"

type PairMethodType byte

const (
	PairingMethodDefault = 0x00
	PairingMethodMFi     = 0x01
	PairingMethodAdd     = 0x03
	PairingMethodDelete  = 0x04
)

func (m PairMethodType) String() string {
	switch m {
	case PairingMethodDefault:
		return "Default"
	case PairingMethodMFi:
		return "MFi"
	case PairingMethodAdd:
		return "Add"
	case PairingMethodDelete:
		return "Delete"
	}
	return fmt.Sprintf("%v Unknown", byte(m))
}
