package pair

import(
    "github.com/brutella/hap"
    "io"
    "fmt"
    "encoding/hex"
    "bytes"
)

type SetupServerController struct {
    context *hap.Context
    bridge *hap.Bridge
    session *SetupServerSession
    curSeq byte
}

func NewSetupServerController(context *hap.Context, bridge *hap.Bridge) (*SetupServerController, error) {
    
    session, err := NewSetupServerSession("Pair-Setup", bridge.Password)
    if err != nil {
        return nil, err
    }
    
    controller := SetupServerController{
                                    context: context,
                                    bridge: bridge,
                                    session: session,
                                    curSeq: WaitingForRequest,
                                }
    
    return &controller, nil
}

func (c *SetupServerController) Handle(r io.Reader) (io.Reader, error) {
    var tlv_out *TLV8Container
    var err error
    
    tlv_in, err := ReadTLV8(r)
    if err != nil {
        return nil, err
    }
    
    method := tlv_in.Byte(TLVType_AuthMethod)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, hap.NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := tlv_in.Byte(TLVType_SequenceNumber)
    fmt.Println("->     Seq:", seq)
    
    switch seq {
    case PairStartRequest:
        if c.curSeq != WaitingForRequest {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        tlv_out, err = c.handlePairStart(tlv_in)
    case PairVerifyRequest:
        if c.curSeq != PairStartRespond {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        tlv_out, err = c.handlePairVerify(tlv_in)
    case PairKeyExchangeRequest:        
        if c.curSeq != PairVerifyRespond {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        tlv_out, err = c.handleKeyExchange(tlv_in)
    default:
        return nil, hap.NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    if err != nil {
        fmt.Println("[ERROR]", err)
        return nil, err
    } else {
        fmt.Println("<-     Seq:", tlv_out.Byte(TLVType_SequenceNumber))
        fmt.Println("-------------")
        return tlv_out.BytesBuffer(), nil
    }
}

// Client -> Server
// - Auth start
//
// Server -> Client
// - B: server public key
// - s: salt
func (c *SetupServerController) handlePairStart(tlv_in *TLV8Container) (*TLV8Container, error) {
    tlv_out := TLV8Container{}
    c.curSeq = PairStartRespond
    
    tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    tlv_out.SetBytes(TLVType_PublicKey, c.session.publicKey)
    tlv_out.SetBytes(TLVType_Salt, c.session.salt)
    
    fmt.Println("<-     B:", hex.EncodeToString(tlv_out.Bytes(TLVType_PublicKey)))
    fmt.Println("<-     s:", hex.EncodeToString(tlv_out.Bytes(TLVType_Salt)))
    
    return &tlv_out, nil
}

// Client -> Server
// - A: client public key
// - M1: proof
// 
// Server -> client
// - M2: proof
// or
// - auth error
func (c *SetupServerController) handlePairVerify(tlv_in *TLV8Container) (*TLV8Container, error) {
    tlv_out := TLV8Container{}
    c.curSeq = PairVerifyRespond
    
    tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    cpublicKey := tlv_in.Bytes(TLVType_PublicKey)
    fmt.Println("->     A:", hex.EncodeToString(cpublicKey))
    
    err := c.session.SetupSecretKeyFromClientPublicKey(cpublicKey)
    if err != nil {
        return nil, err
    }
    
    cproof := tlv_in.Bytes(TLVType_Proof)
    fmt.Println("->     M1:", hex.EncodeToString(cproof))
    
    sproof, err := c.session.ProofFromClientProof(cproof)
    if err != nil || len(sproof) == 0 { // proof `M1` is wrong
        fmt.Println("[Failed] Proof M1 is wrong")
        c.reset()
        tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
    } else {
        fmt.Println("[Success] Proof M1 is valid")
        err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
        if err != nil {
            return nil, err
        }
        
        // Return proof `M1`
        tlv_out.SetBytes(TLVType_Proof, sproof)
    }
    
    fmt.Println("<-     M2:", hex.EncodeToString(tlv_out.Bytes(TLVType_Proof)))
    fmt.Println("        S:", hex.EncodeToString(c.session.secretKey))
    fmt.Println("        K:", hex.EncodeToString(c.session.encryptionKey[:]))
    
    return &tlv_out, nil
}

// Client -> Server
// - encrypted tlv8: client LTPK, client name and signature (of H, client name, LTPK)
// - auth tag (mac)
// 
// Server
// - Validate signature of encrpyted tlv8
// - Read and store client LTPK and name
// 
// Server -> Client
// - encrpyted tlv8: bridge LTPK, bridge name, signature (of H2, bridge name, LTPK)
func (c *SetupServerController) handleKeyExchange(tlv_in *TLV8Container) (*TLV8Container, error) {
    tlv_out := TLV8Container{}
    
    c.curSeq = PairKeyExchangeRespond
    
    tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    data := tlv_in.Bytes(TLVType_EncryptedData)    
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := hap.Chacha20DecryptAndPoly1305Verify(c.session.encryptionKey[:], []byte("PS-Msg05"), message, mac, nil)
    
    if err != nil {
        c.reset()
        fmt.Println(err)
        tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_UnkownError) // return error 1
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        tlv_in, err := ReadTLV8(decrypted_buffer)
        if err != nil {
            return nil, err
        }
        
        username  := tlv_in.String(TLVType_Username)
        ltpk      := tlv_in.Bytes(TLVType_PublicKey)
        signature := tlv_in.Bytes(TLVType_Ed25519Signature)
        fmt.Println("->     Username:", username)
        fmt.Println("->     LTPK:", hex.EncodeToString(ltpk))
        fmt.Println("->     Signature:", hex.EncodeToString(signature))
        
        // Calculate `H`
        H, _ := hap.HKDF_SHA512(c.session.secretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
        material := make([]byte, 0)
        material = append(material, H[:]...)
        material = append(material, []byte(username)...)
        material = append(material, ltpk...)
        
        if hap.ValidateED25519Signature(ltpk, material, signature) == false {
            fmt.Println("[Failed] ed25519 signature is invalid")
            c.reset()
            tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
        } else {
            fmt.Println("[Success] ed25519 signature is valid")
            // Store client LTPK and name
            client := hap.NewClient(username, ltpk)
            c.context.SaveClient(client)
            fmt.Printf("[Storage] Stored LTPK '%s' for client '%s'\n", hex.EncodeToString(ltpk), username)
            
            LTPK := c.context.PublicKeyForAccessory(c.bridge)
            LTSK := c.context.SecretKeyForAccessory(c.bridge)
            
            // Send username, LTPK, signature as encrypted message
            H2, err := hap.HKDF_SHA512(c.session.secretKey, []byte("Pair-Setup-Accessory-Sign-Salt"), []byte("Pair-Setup-Accessory-Sign-Info"))
            material = make([]byte, 0)
            material = append(material, H2[:]...)
            material = append(material, []byte(c.bridge.Name)...)
            material = append(material, LTPK...)

            signature, err := hap.ED25519Signature(LTSK, material)
            if err != nil {
                return nil, err
            }
            
            tlvPairKeyExchange := TLV8Container{}
            tlvPairKeyExchange.SetString(TLVType_Username, c.bridge.Name)
            tlvPairKeyExchange.SetBytes(TLVType_PublicKey, LTPK)
            tlvPairKeyExchange.SetBytes(TLVType_Ed25519Signature, []byte(signature))
            
            fmt.Println("<-     Username:", tlvPairKeyExchange.String(TLVType_Username))
            fmt.Println("<-     LTPK:", hex.EncodeToString(tlvPairKeyExchange.Bytes(TLVType_PublicKey)))
            fmt.Println("<-     Signature:", hex.EncodeToString(tlvPairKeyExchange.Bytes(TLVType_Ed25519Signature)))
            
            encrypted, mac, _ := hap.Chacha20EncryptAndPoly1305Seal(c.session.encryptionKey[:], []byte("PS-Msg06"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)    
            tlv_out.SetByte(TLVType_AuthMethod, 0)
            tlv_out.SetByte(TLVType_SequenceNumber, PairKeyExchangeRequest)
            tlv_out.SetBytes(TLVType_EncryptedData, append(encrypted, mac[:]...))
            
            c.reset()
        }
    }
    
    return &tlv_out, nil
}

func (c *SetupServerController) reset() {
    c.curSeq = WaitingForRequest
    // TODO: reset session
}