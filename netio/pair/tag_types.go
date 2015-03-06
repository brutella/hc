package pair

type TagType byte

const (
	TagPairingMethod    = 0x00 // PairingMethodType
	TagUsername         = 0x01 // string
	TagSalt             = 0x02 // 16 bytes
	TagPublicKey        = 0x03 // either SRP client public key (384 bytes) or ED25519 LTPK (32 bytes)
	TagProof            = 0x04 // 64 bytes (SRP proof)
	TagEncryptedData    = 0x05 // data with auth tag (mac)
	TagSequence         = 0x06 // PairSequenceType
	TagError            = 0x07 // ErrorType
	TagEd25519Signature = 0x0A // 64 bytes; TODO remove Ed25519 from name

	TagMFiCertificate = 0x09
	TagMFiSignature   = 0x0A
)
