package pair

import (
	"errors"
	"fmt"
)

type errCode byte

const (
	// ErrCodeNo is code for no error
	ErrCodeNo errCode = 0x00

	// ErrCodeUnknown is code for unknown error
	ErrCodeUnknown errCode = 0x01

	// ErrCodeAuthenticationFailed is code for authentication error e.g. client proof is wrong
	ErrCodeAuthenticationFailed errCode = 0x02

	// ErrCodeTooManyAttempts is code for too many attempts error (not used)
	ErrCodeTooManyAttempts errCode = 0x03

	// ErrCodeUnknownPeer is code for unknown accessory or client error
	ErrCodeUnknownPeer errCode = 0x04

	// ErrCodeMaxPeer is code for reaching maximum number of peers error (not used)
	ErrCodeMaxPeer errCode = 0x05

	// ErrCodeMaxAuthenticationAttempts is code for reaching maximum number of authentication attempts error (not used)
	ErrCodeMaxAuthenticationAttempts errCode = 0x06
)

func (t errCode) Byte() byte {
	return byte(t)
}

func (t errCode) Error() error {
	return errors.New(t.String())
}

func (t errCode) String() string {
	switch t {
	case ErrCodeNo:
		return "None"
	case ErrCodeUnknown:
		return "Unknown"
	case ErrCodeAuthenticationFailed:
		return "Authentication Failed"
	case ErrCodeTooManyAttempts:
		return "Too Many Attempts"
	case ErrCodeUnknownPeer:
		return "Unknown Peer"
	case ErrCodeMaxPeer:
		return "Max Peer"
	case ErrCodeMaxAuthenticationAttempts:
		return "Max Authentication Attempts"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}
