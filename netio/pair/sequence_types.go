package pair

import "fmt"

// pairStepType defines the type of pairing steps.
type pairStepType byte

const (
	// PairStepWaiting is the step when waiting server waits for pairing request from a client.
	PairStepWaiting pairStepType = 0x00

	// PairStepStartRequest sent from the client to the accessory to start pairing.
	PairStepStartRequest pairStepType = 0x01

	// PairStepStartResponse sent from the accessory to the client alongside the server's salt and public key
	PairStepStartResponse pairStepType = 0x02

	// PairStepVerifyRequest sent from the client to the accessory alongside the client public key and proof.
	PairStepVerifyRequest pairStepType = 0x03

	// PairStepVerifyResponse sent from the accessory to the client alongside the server's proof.
	PairStepVerifyResponse pairStepType = 0x04

	// PairStepKeyExchangeRequest sent from the client to the accessory alongside the client encrypted username and public key
	PairStepKeyExchangeRequest pairStepType = 0x05

	// PairStepKeyExchangeResponse sent from the accessory to the client alongside the accessory encrypted username and public key
	PairStepKeyExchangeResponse pairStepType = 0x06
)

// Byte returns the raw byte value.
func (t pairStepType) Byte() byte {
	return byte(t)
}

func (t pairStepType) String() string {
	switch t {
	case PairStepWaiting:
		return "Waiting"
	case PairStepStartRequest:
		return "Pairing Start Request"
	case PairStepStartResponse:
		return "Pairing Start Response"
	case PairStepVerifyRequest:
		return "Pairing Verify Request"
	case PairStepVerifyResponse:
		return "Pair Verify Response"
	case PairStepKeyExchangeRequest:
		return "Pair Key Exchange Request"
	case PairStepKeyExchangeResponse:
		return "Pair Key Exchange Response"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}

// VerifyStepType defines the type of pairing verification steps.
type VerifyStepType byte

const (
	// VerifyStepWaiting is the step when waiting server waits for pair verification request from the client.
	VerifyStepWaiting VerifyStepType = 0x00
	// VerifyStepStartRequest sent from the client to the accessory to start pairing verification alongside the client public key.
	VerifyStepStartRequest VerifyStepType = 0x01
	// VerifyStepStartResponse sent from the accessory to the client alongside the accessory public key and signature (derived from the on the accessory public key, username and client public and private key)
	VerifyStepStartResponse VerifyStepType = 0x02
	// VerifyStepFinishRequest sent from the client to the accessory alongside the client public key and signature (derived from the on the client public key, username and accessory public and private key)
	VerifyStepFinishRequest VerifyStepType = 0x03
	// VerifyStepFinishResponse sent from the accessory to the client alongside an error code when verification failed.
	VerifyStepFinishResponse VerifyStepType = 0x04
)

// Byte returns the raw byte value.
func (t VerifyStepType) Byte() byte {
	return byte(t)
}

func (t VerifyStepType) String() string {
	switch t {
	case VerifyStepWaiting:
		return "Waiting"
	case VerifyStepStartRequest:
		return "Verify Start Request"
	case VerifyStepStartResponse:
		return "Verify Start Response"
	case VerifyStepFinishRequest:
		return "Verify Finish Request"
	case VerifyStepFinishResponse:
		return "Verify Finish Response"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}
