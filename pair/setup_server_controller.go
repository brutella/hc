package pair

import(
    "github.com/brutella/hap"
    "fmt"
    "encoding/hex"
    "bytes"
)

type SetupServerController struct {
    Handler
    context *hap.Context
    bridge *hap.Bridge
    session *SetupServerSession
    curSeq byte
}

func NewSetupServerController(context *hap.Context, bridge *hap.Bridge) (*SetupServerController, error) {
    
    session, err := NewSetupServerSession("Pair-Setup", bridge.Password())
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

func (c *SetupServerController) Handle(cont_in Container) (Container, error) {
    var cont_out Container
    var err error
    
    method := cont_in.GetByte(TLVType_Method)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, hap.NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := cont_in.GetByte(TLVType_SequenceNumber)
    
    switch seq {
    case PairStartRequest:
        if c.curSeq != WaitingForRequest {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handlePairStart(cont_in)
    case PairVerifyRequest:
        if c.curSeq != PairStartRespond {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handlePairVerify(cont_in)
    case PairKeyExchangeRequest:        
        if c.curSeq != PairVerifyRespond {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handleKeyExchange(cont_in)
    default:
        return nil, hap.NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    return cont_out, err
}

// Client -> Server
// - Auth start
//
// Server -> Client
// - B: server public key
// - s: salt
func (c *SetupServerController) handlePairStart(cont_in Container) (Container, error) {
    cont_out := NewTLV8Container()
    c.curSeq = PairStartRespond
    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    cont_out.SetBytes(TLVType_PublicKey, c.session.publicKey)
    cont_out.SetBytes(TLVType_Salt, c.session.salt)
    
    fmt.Println("<-     B:", hex.EncodeToString(cont_out.GetBytes(TLVType_PublicKey)))
    fmt.Println("<-     s:", hex.EncodeToString(cont_out.GetBytes(TLVType_Salt)))
    
    return cont_out, nil
}

// Client -> Server
// - A: client public key
// - M1: proof
// 
// Server -> client
// - M2: proof
// or
// - auth error
func (c *SetupServerController) handlePairVerify(cont_in Container) (Container, error) {
    cont_out := NewTLV8Container()
    c.curSeq = PairVerifyRespond
    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    cpublicKey := cont_in.GetBytes(TLVType_PublicKey)
    fmt.Println("->     A:", hex.EncodeToString(cpublicKey))
    
    err := c.session.SetupSecretKeyFromClientPublicKey(cpublicKey)
    if err != nil {
        return nil, err
    }
    
    cproof := cont_in.GetBytes(TLVType_Proof)
    fmt.Println("->     M1:", hex.EncodeToString(cproof))
    
    sproof, err := c.session.ProofFromClientProof(cproof)
    if err != nil || len(sproof) == 0 { // proof `M1` is wrong
        fmt.Println("[Failed] Proof M1 is wrong")
        c.reset()
        cont_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
    } else {
        fmt.Println("[Success] Proof M1 is valid")
        err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
        if err != nil {
            return nil, err
        }
        
        // Return proof `M1`
        cont_out.SetBytes(TLVType_Proof, sproof)
    }
    
    fmt.Println("<-     M2:", hex.EncodeToString(cont_out.GetBytes(TLVType_Proof)))
    fmt.Println("        S:", hex.EncodeToString(c.session.secretKey))
    fmt.Println("        K:", hex.EncodeToString(c.session.encryptionKey[:]))
    
    return cont_out, nil
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
func (c *SetupServerController) handleKeyExchange(cont_in Container) (Container, error) {
    cont_out := NewTLV8Container()
    
    c.curSeq = PairKeyExchangeRespond
    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    data := cont_in.GetBytes(TLVType_EncryptedData)    
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := hap.Chacha20DecryptAndPoly1305Verify(c.session.encryptionKey[:], []byte("PS-Msg05"), message, mac, nil)
    
    if err != nil {
        c.reset()
        fmt.Println(err)
        cont_out.SetByte(TLVType_ErrorCode, TLVStatus_UnkownError) // return error 1
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        cont_in, err := NewTLV8ContainerFromReader(decrypted_buffer)
        if err != nil {
            return nil, err
        }
        
        username  := cont_in.GetString(TLVType_Username)
        ltpk      := cont_in.GetBytes(TLVType_PublicKey)
        signature := cont_in.GetBytes(TLVType_Ed25519Signature)
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
            cont_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
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
            material = append(material, []byte(c.bridge.Id())...)
            material = append(material, LTPK...)

            signature, err := hap.ED25519Signature(LTSK, material)
            if err != nil {
                return nil, err
            }
            
            tlvPairKeyExchange := NewTLV8Container()
            tlvPairKeyExchange.SetString(TLVType_Username, c.bridge.Id())
            tlvPairKeyExchange.SetBytes(TLVType_PublicKey, LTPK)
            tlvPairKeyExchange.SetBytes(TLVType_Ed25519Signature, []byte(signature))
            
            fmt.Println("<-     Username:", tlvPairKeyExchange.GetString(TLVType_Username))
            fmt.Println("<-     LTPK:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TLVType_PublicKey)))
            fmt.Println("<-     Signature:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TLVType_Ed25519Signature)))
            
            encrypted, mac, _ := hap.Chacha20EncryptAndPoly1305Seal(c.session.encryptionKey[:], []byte("PS-Msg06"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)    
            cont_out.SetByte(TLVType_Method, 0)
            cont_out.SetByte(TLVType_SequenceNumber, PairKeyExchangeRequest)
            cont_out.SetBytes(TLVType_EncryptedData, append(encrypted, mac[:]...))
            
            c.reset()
        }
    }
    
    return cont_out, nil
}

func (c *SetupServerController) reset() {
    c.curSeq = WaitingForRequest
    // TODO: reset session
}