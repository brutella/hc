package pair

import (
	"errors"
	"fmt"
)

var ErrInvalidClientKeyLength = errors.New("Invalid client public key size")

var ErrInvalidPairMethod = func(m PairMethodType) error {
	return errors.New(fmt.Sprintf("Invalid pairing method %v\n", m))
}

var ErrInvalidPairStep = func(t PairStepType) error {
	return errors.New(fmt.Sprintf("Invalid pairing step %v\n", t))
}

var ErrInvalidInternalPairStep = func(t PairStepType) error {
	return errors.New(fmt.Sprintf("Invalid internal pairing step %v\n", t))
}

var ErrInvalidVerifyStep = func(t VerifyStepType) error {
	return errors.New(fmt.Sprintf("Invalid pairing verify step %v\n", t))
}

var ErrInvalidInternalVerifyStep = func(t VerifyStepType) error {
	return errors.New(fmt.Sprintf("Invalid internal pairing verify step %v\n", t))
}
