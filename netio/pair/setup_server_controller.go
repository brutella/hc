package pair

import(
    "github.com/brutella/hap/db"
    "github.com/brutella/hap/crypto"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/netio"
    "github.com/brutella/log"
    
    "encoding/hex"
    "bytes"
)

type SetupServerController struct {
    bridge *netio.Bridge
    session *PairSetupServerSession
    curSeq byte
    database db.Database
}

func NewSetupServerController(bridge *netio.Bridge, database db.Database) (*SetupServerController, error) {
    
    session, err := NewPairSetupServerSession(bridge.Id(), bridge.Password())
    if err != nil {
        return nil, err
    }
    
    controller := SetupServerController{
                                    bridge: bridge,
                                    session: session,
                                    database: database,
                                    curSeq: WaitingForRequest,
                                }
    
    return &controller, nil
}

func (c *SetupServerController) Handle(cont_in common.Container) (common.Container, error) {
    var cont_out common.Container
    var err error
    
    method := cont_in.GetByte(TLVType_Method)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, common.NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := cont_in.GetByte(TLVType_SequenceNumber)
    
    switch seq {
    case PairStartRequest:
        if c.curSeq != WaitingForRequest {
            c.reset()
            return nil, common.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handlePairStart(cont_in)
    case PairVerifyRequest:
        if c.curSeq != PairStartRespond {
            c.reset()
            return nil, common.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handlePairVerify(cont_in)
    case PairKeyExchangeRequest:        
        if c.curSeq != PairVerifyRespond {
            c.reset()
            return nil, common.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handleKeyExchange(cont_in)
    default:
        return nil, common.NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    return cont_out, err
}

// Client -> Server
// - Auth start
//
// Server -> Client
// - B: server public key
// - s: salt
func (c *SetupServerController) handlePairStart(cont_in common.Container) (common.Container, error) {
    cont_out := common.NewTLV8Container()
    c.curSeq = PairStartRespond
    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    cont_out.SetBytes(TLVType_PublicKey, c.session.PublicKey)
    cont_out.SetBytes(TLVType_Salt, c.session.Salt)
    
    log.Println("[INFO] <-     B:", hex.EncodeToString(cont_out.GetBytes(TLVType_PublicKey)))
    log.Println("[INFO] <-     s:", hex.EncodeToString(cont_out.GetBytes(TLVType_Salt)))
    
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
func (c *SetupServerController) handlePairVerify(cont_in common.Container) (common.Container, error) {
    cont_out := common.NewTLV8Container()
    c.curSeq = PairVerifyRespond
    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    cpublicKey := cont_in.GetBytes(TLVType_PublicKey)
    log.Println("[INFO] ->     A:", hex.EncodeToString(cpublicKey))
    
    err := c.session.SetupSecretKeyFromClientPublicKey(cpublicKey)
    if err != nil {
        return nil, err
    }
    
    cproof := cont_in.GetBytes(TLVType_Proof)
    log.Println("[INFO] ->     M1:", hex.EncodeToString(cproof))
    
    sproof, err := c.session.ProofFromClientProof(cproof)
    if err != nil || len(sproof) == 0 { // proof `M1` is wrong
        log.Println("[WARN] Proof M1 is wrong")
        c.reset()
        cont_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
    } else {
        log.Println("[INFO] Proof M1 is valid")
        err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
        if err != nil {
            return nil, err
        }
        
        // Return proof `M1`
        cont_out.SetBytes(TLVType_Proof, sproof)
    }
    
    log.Println("[INFO] <-     M2:", hex.EncodeToString(cont_out.GetBytes(TLVType_Proof)))
    log.Println("[INFO]         S:", hex.EncodeToString(c.session.SecretKey))
    log.Println("[INFO]         K:", hex.EncodeToString(c.session.EncryptionKey[:]))
    
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
func (c *SetupServerController) handleKeyExchange(cont_in common.Container) (common.Container, error) {
    cont_out := common.NewTLV8Container()
    
    c.curSeq = PairKeyExchangeRespond
    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    data := cont_in.GetBytes(TLVType_EncryptedData)    
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    log.Println("[INFO] ->     Message:", hex.EncodeToString(message))
    log.Println("[INFO] ->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(c.session.EncryptionKey[:], []byte("PS-Msg05"), message, mac, nil)
    
    if err != nil {
        c.reset()
        log.Println("[ERROR]", err)
        cont_out.SetByte(TLVType_ErrorCode, TLVStatus_UnkownError) // return error 1
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        cont_in, err := common.NewTLV8ContainerFromReader(decrypted_buffer)
        if err != nil {
            return nil, err
        }
        
        username  := cont_in.GetString(TLVType_Username)
        ltpk      := cont_in.GetBytes(TLVType_PublicKey)
        signature := cont_in.GetBytes(TLVType_Ed25519Signature)
        log.Println("[INFO] ->     Username:", username)
        log.Println("[INFO] ->     LTPK:", hex.EncodeToString(ltpk))
        log.Println("[INFO] ->     Signature:", hex.EncodeToString(signature))
        
        // Calculate `H`
        H, _ := crypto.HKDF_SHA512(c.session.SecretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
        material := make([]byte, 0)
        material = append(material, H[:]...)
        material = append(material, []byte(username)...)
        material = append(material, ltpk...)
        
        if crypto.ValidateED25519Signature(ltpk, material, signature) == false {
            log.Println("[WARN] ed25519 signature is invalid")
            c.reset()
            cont_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
        } else {
            log.Println("[INFO] ed25519 signature is valid")
            // Store client LTPK and name
            client := db.NewClient(username, ltpk)
            c.database.SaveClient(client)
            log.Printf("[INFO] Stored LTPK '%s' for client '%s'\n", hex.EncodeToString(ltpk), username)
            
            LTPK := c.bridge.PublicKey
            LTSK := c.bridge.SecretKey
            
            // Send username, LTPK, signature as encrypted message
            H2, err := crypto.HKDF_SHA512(c.session.SecretKey, []byte("Pair-Setup-Accessory-Sign-Salt"), []byte("Pair-Setup-Accessory-Sign-Info"))
            material = make([]byte, 0)
            material = append(material, H2[:]...)
            material = append(material, []byte(c.session.Username)...)
            material = append(material, LTPK...)

            signature, err := crypto.ED25519Signature(LTSK, material)
            if err != nil {
                return nil, err
            }
            
            tlvPairKeyExchange := common.NewTLV8Container()
            tlvPairKeyExchange.SetBytes(TLVType_Username, c.session.Username)
            tlvPairKeyExchange.SetBytes(TLVType_PublicKey, LTPK)
            tlvPairKeyExchange.SetBytes(TLVType_Ed25519Signature, []byte(signature))
            
            log.Println("[INFO] <-     Username:", tlvPairKeyExchange.GetString(TLVType_Username))
            log.Println("[INFO] <-     LTPK:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TLVType_PublicKey)))
            log.Println("[INFO] <-     Signature:", hex.EncodeToString(tlvPairKeyExchange.GetBytes(TLVType_Ed25519Signature)))
            
            encrypted, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(c.session.EncryptionKey[:], []byte("PS-Msg06"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)    
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