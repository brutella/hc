package pair

import "fmt"

type ErrorType byte

const (
	ErrorNone                      ErrorType = 0x00
	ErrorUnknown                   ErrorType = 0x01
	ErrorAuthenticationFailed      ErrorType = 0x02 // e.g. client proof `M1` is wrong
	ErrorTooManyAttempts           ErrorType = 0x03
	ErrorUnknownPeer               ErrorType = 0x04
	ErrorMaxPeer                   ErrorType = 0x05
	ErrorMaxAuthenticationAttempts ErrorType = 0x06
)

func (t ErrorType) Byte() byte {
	return byte(t)
}

func (t ErrorType) String() string {
	switch t {
	case ErrorNone:
		return "None"
	case ErrorUnknown:
		return "Unknown"
	case ErrorAuthenticationFailed:
		return "Authentication Failed"
	case ErrorTooManyAttempts:
		return "Too Many Attemps"
	case ErrorUnknownPeer:
		return "Unknown Peer"
	case ErrorMaxPeer:
		return "Max Peer"
	case ErrorMaxAuthenticationAttempts:
		return "Max Authentication Attempts"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}
