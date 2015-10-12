package pair

import "fmt"

type PairMethodType byte

const (
	// PairingMethodDefault is the default pairing method.
	PairingMethodDefault PairMethodType = 0x00

	// PairingMethodMFi is used to pair with an MFi compliant accessory (not used).
	PairingMethodMFi PairMethodType = 0x01

	// PairingMethodAdd is used to pair a client by exchanging keys on a secured
	// connection and without going through the pairing process.
	PairingMethodAdd PairMethodType = 0x03

	// PairingMethodDelete is used to delete a pairing with a client.
	PairingMethodDelete PairMethodType = 0x04
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

func (m PairMethodType) Byte() byte {
	return byte(m)
}
