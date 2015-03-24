package pair

import (
	"errors"
	"fmt"
)

var errInvalidClientKeyLength = errors.New("Invalid client public key size")

var errInvalidPairMethod = func(m pairMethodType) error {
	return fmt.Errorf("Invalid pairing method %v\n", m)
}

var errInvalidPairStep = func(t pairStepType) error {
	return fmt.Errorf("Invalid pairing step %v\n", t)
}

var errInvalidInternalPairStep = func(t pairStepType) error {
	return fmt.Errorf("Invalid internal pairing step %v\n", t)
}

var errInvalidVerifyStep = func(t VerifyStepType) error {
	return fmt.Errorf("Invalid pairing verify step %v\n", t)
}

var errInvalidInternalVerifyStep = func(t VerifyStepType) error {
	return fmt.Errorf("Invalid internal pairing verify step %v\n", t)
}
