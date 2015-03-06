package pair

import "fmt"

type ErrCode byte

const (
	ErrCodeNo                        ErrCode = 0x00
	ErrCodeUnknown                   ErrCode = 0x01
	ErrCodeAuthenticationFailed      ErrCode = 0x02 // e.g. client proof `M1` is wrong
	ErrCodeTooManyAttempts           ErrCode = 0x03
	ErrCodeUnknownPeer               ErrCode = 0x04
	ErrCodeMaxPeer                   ErrCode = 0x05
	ErrCodeMaxAuthenticationAttempts ErrCode = 0x06
)

func (t ErrCode) Byte() byte {
	return byte(t)
}

func (t ErrCode) String() string {
	switch t {
	case ErrCodeNo:
		return "None"
	case ErrCodeUnknown:
		return "Unknown"
	case ErrCodeAuthenticationFailed:
		return "Authentication Failed"
	case ErrCodeTooManyAttempts:
		return "Too Many Attemps"
	case ErrCodeUnknownPeer:
		return "Unknown Peer"
	case ErrCodeMaxPeer:
		return "Max Peer"
	case ErrCodeMaxAuthenticationAttempts:
		return "Max Authentication Attempts"
	}
	return fmt.Sprintf("%v Unknown", byte(t))
}
