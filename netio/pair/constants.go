package pair

type TLVType byte
// TLV types
const (
	TLVMethod           = 0x00 // integer, either 0x00 (uncertified) or 0x01 (MFi compliant)
	TLVUsername         = 0x01 // string
	TLVSalt             = 0x02 // 16 bytes
	TLVPublicKey        = 0x03 // either SRP client public key (384 bytes) or ED25519 LTPK (32 bytes)
	TLVProof            = 0x04 // 64 bytes (SRP proof)
	TLVEncryptedData    = 0x05 // data with auth tag (mac)
	TLVSequenceNumber   = 0x06 // integer
	TLVErrorCode        = 0x07 // integer, see TLVStatus
	TLVEd25519Signature = 0x0A // 64 bytes

	TLVMFiCertificate = 0x09
	TLVMFiSignature   = 0x0A
)

type MethodType byte
const (
	MethodDefault = 0x00
	MethodMFi     = 0x01
	MethodAdd     = 0x03
	MethodDelete  = 0x04
)

// TLV errors
const (
	TLVStatus_NoError              = 0x00
	TLVStatus_UnkownError          = 0x01
	TLVStatus_AuthError            = 0x02 // e.g. client proof `M1` is wrong
	TLVStatus_TooManyAttemptsError = 0x03
	TLVStatus_UnkownPeerError      = 0x04
	TLVStatus_MaxPeerError         = 0x05
	TLVStatus_MaxAuthAttempsError  = 0x06
)
