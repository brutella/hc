package pair

import "fmt"

type PairStepType byte

const (
	PairStepWaiting             PairStepType = 0x00
	PairStepStartRequest        PairStepType = 0x01
	PairStepStartResponse       PairStepType = 0x02
	PairStepVerifyRequest       PairStepType = 0x03
	PairStepVerifyResponse      PairStepType = 0x04
	PairStepKeyExchangeRequest  PairStepType = 0x05
	PairStepKeyExchangeResponse PairStepType = 0x06
)

func (t PairStepType) Byte() byte {
	return byte(t)
}

func (t PairStepType) String() string {
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

type VerifyStepType byte

const (
	VerifyStepWaiting        VerifyStepType = 0x00
	VerifyStepStartRequest   VerifyStepType = 0x01
	VerifyStepStartResponse  VerifyStepType = 0x02
	VerifyStepFinishRequest  VerifyStepType = 0x03
	VerifyStepFinishResponse VerifyStepType = 0x04
)

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
