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

type VerifySequenceType byte

const (
	SequenceVerifyWaiting        VerifySequenceType = 0x00
	SequenceVerifyStartRequest   VerifySequenceType = 0x01
	SequenceVerifyStartResponse  VerifySequenceType = 0x02
	SequenceVerifyFinishRequest  VerifySequenceType = 0x03
	SequenceVerifyFinishResponse VerifySequenceType = 0x04
)

func (t VerifySequenceType) Byte() byte {
	return byte(t)
}

func (t VerifySequenceType) String() string {
	switch t {
	case SequenceVerifyWaiting:
		return "Waiting"
	case SequenceVerifyStartRequest:
		return "Verify Start Request"
	case SequenceVerifyStartResponse:
		return "Verify Start Response"
	case SequenceVerifyFinishRequest:
		return "Verify Finish Request"
	case SequenceVerifyFinishResponse:
		return "Verify Finish Response"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}
