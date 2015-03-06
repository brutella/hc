package pair

type TagType byte

const (
	TagPairingMethod    = 0x00 // PairingMethodType
	TagUsername         = 0x01 // string
	TagSalt             = 0x02 // 16 bytes
	TagPublicKey        = 0x03 // either SRP client public key (384 bytes) or ED25519 LTPK (32 bytes)
	TagProof            = 0x04 // 64 bytes (SRP proof)
	TagEncryptedData    = 0x05 // data with auth tag (mac)
	TagSequence         = 0x06 // SequenceType
	TagError            = 0x07 // ErrorType
	TagEd25519Signature = 0x0A // 64 bytes; TODO remove Ed25519 from name

	TagMFiCertificate = 0x09
	TagMFiSignature   = 0x0A
)

type PairingMethodType byte

const (
	PairingMethodDefault = 0x00
	PairingMethodMFi     = 0x01
	PairingMethodAdd     = 0x03
	PairingMethodDelete  = 0x04
)

type SequenceType byte

const (
	SequenceWaitingForRequest = 0x00

	SequencePairStartRequest        = 0x01
	SequencePairStartResponse       = 0x02
	SequencePairVerifyRequest       = 0x03
	SequencePairVerifyResponse      = 0x04
	SequencePairKeyExchangeRequest  = 0x05
	SequencePairKeyExchangeResponse = 0x06

	SequenceVerifyStartRequest   = 0x01
	SequenceVerifyStartResponse  = 0x02
	SequenceVerifyFinishRequest  = 0x03
	SequenceVerifyFinishResponse = 0x04
)

type ErrorType byte

const (
	ErrorNone                      = 0x00
	ErrorUnknown                   = 0x01
	ErrorAuthenticationFailed      = 0x02 // e.g. client proof `M1` is wrong
	ErrorTooManyAttempts           = 0x03
	ErrorUnknownPeer               = 0x04
	ErrorMaxPeer                   = 0x05
	ErrorMaxAuthenticationAttempts = 0x06
)
