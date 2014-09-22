package gohap

const (
    TLVType_AuthMethod       = 0x00 // integer, either 0x00 (uncertified) or 0x01 (MFi compliant)
    TLVType_Username         = 0x01 // string
    TLVType_Salt             = 0x02 // 16 bytes
    TLVType_PublicKey        = 0x03 // either SRP client public key (384 bytes) or ED25519 LTPK (32 bytes)
    TLVType_Proof            = 0x04 // SRP proof
    TLVType_EncryptedData    = 0x05 // data with auth tag
    TLVType_SequenceNumber   = 0x06 // integer
    TLVType_ErrorCode        = 0x07 // integer, see TLVStatus
    TLVType_Ed25519Signature = 0x0A // 64 bytes
    
    TLVType_MFiCertificate   = 0x09
    TLVType_MFiSignature     = 0x0A
)

const (
    TLVStatus_NoError               = 0x00
    TLVStatus_UnkownError           = 0x01
    TLVStatus_AuthError             = 0x02 // e.g. client proof `M1` is wrong
    TLVStatus_TooManyAttemptsError  = 0x03
    TLVStatus_UnkownPeerError       = 0x04
    TLVStatus_MaxPeerError          = 0x05
    TLVStatus_MaxAuthAttempsError   = 0x06
)

const (
    SRPUsername = "Pair-Setup"
)

const (
    HKDFSalt = "Pair-Setup-Encrypt-Salt"
    HKDFInfo = "Pair-Setup-Encrypt-Info"
)

const (
    HTTPEndpointPairSetup = "pair-setup"
)

const (
    HTTPContentTypePairingTLV8 = "application/pairing+tlv8"
)

