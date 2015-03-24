package pair

// These constants are used to access values in a TLV8 container.
const (
	// TagPairingMethod is the paring method tag. The value is of type pairMethodType.
	TagPairingMethod = 0x00

	// TagUsername is the username tag. The value is of type string.
	TagUsername = 0x01

	// TagSalt is sthe salt tag. The value is of type 16 bytes.
	TagSalt = 0x02

	// TagPublicKey is the public key tag. The value is either SRP client public key (384 bytes) or ED25519 public key (32 bytes) - depending on the context.
	TagPublicKey = 0x03

	// TagProof is the SRP proof tag. The value is of type 64 bytes.
	TagProof = 0x04

	// TagEncryptedData is the encrypted data tag. The value includes the encrypted message and auth tag.
	TagEncryptedData = 0x05

	// TagSequence is the sequence tag. The value is of type pairStepType or VerifyStepType - depending on the context.
	TagSequence = 0x06

	// TagErrCode is the error tag. The value is of type ErrCode.
	TagErrCode = 0x07

	// TagSignature is the Ed25519 signature tag. The value is of type 64 bytes.
	TagSignature = 0x0A

	// TagMFiCertificate is the MFi certificate tag (currently not used).
	TagMFiCertificate = 0x09

	// TagMFiSignature is the MFi signature tag (currently not used).
	TagMFiSignature = 0x0A
)
