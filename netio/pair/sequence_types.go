package pair

import "fmt"

type PairSequenceType byte

const (
	SequencePairWaiting             PairSequenceType = 0x00
	SequencePairStartRequest        PairSequenceType = 0x01
	SequencePairStartResponse       PairSequenceType = 0x02
	SequencePairVerifyRequest       PairSequenceType = 0x03
	SequencePairVerifyResponse      PairSequenceType = 0x04
	SequencePairKeyExchangeRequest  PairSequenceType = 0x05
	SequencePairKeyExchangeResponse PairSequenceType = 0x06
)

func (t PairSequenceType) Byte() byte {
	return byte(t)
}

func (t PairSequenceType) String() string {
	switch t {
	case SequencePairWaiting:
		return "Waiting"
	case SequencePairStartRequest:
		return "Pairing Start Request"
	case SequencePairStartResponse:
		return "Pairing Start Response"
	case SequencePairVerifyRequest:
		return "Pairing Verify Request"
	case SequencePairVerifyResponse:
		return "Pair Verify Response"
	case SequencePairKeyExchangeRequest:
		return "Pair Key Exchange Request"
	case SequencePairKeyExchangeResponse:
		return "Pair Key Exchange Response"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}

type VerifyStepType byte

const (
	StepVerifyWaiting        VerifyStepType = 0x00
	StepVerifyStartRequest   VerifyStepType = 0x01
	StepVerifyStartResponse  VerifyStepType = 0x02
	StepVerifyFinishRequest  VerifyStepType = 0x03
	StepVerifyFinishResponse VerifyStepType = 0x04
)

func (t VerifyStepType) Byte() byte {
	return byte(t)
}

func (t VerifyStepType) String() string {
	switch t {
	case StepVerifyWaiting:
		return "Waiting"
	case StepVerifyStartRequest:
		return "Verify Start Request"
	case StepVerifyStartResponse:
		return "Verify Start Response"
	case StepVerifyFinishRequest:
		return "Verify Finish Request"
	case StepVerifyFinishResponse:
		return "Verify Finish Response"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}
