package gohap

const (
    TLVType_AuthMethod       = 0x00 // integer, either 0x00 (uncertified) or 0x01 (MFi compliant)
    TLVType_Username         = 0x01 // string
    TLVType_Salt             = 0x02 // 16 bytes
    TLVType_PublicKey        = 0x03 // either Curve25519 (384 bytes)  or SRP (32 bytes) public key
    TLVType_Proof            = 0x04 // ED25519 or SRP proof
    TLVType_EncryptedData    = 0x05 // data with auth tag
    TLVType_SequenceNumber   = 0x06 // integer
    TLVType_ErrorCode        = 0x07 // integer
    TLVType_Ed25519Signature = 0x0A // 64 bytes
    
    TLVType_MFiCertificate   = 0x09
    TLVType_MFiSignature     = 0x0A
)

const (
    HTTPEndpointPairSetup = "pair-setup"
)

const (
    HTTPContentTypePairingTLV8 = "application/pairing+tlv8"
)

