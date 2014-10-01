package pair

import(
    "github.com/brutella/hap"
    "github.com/brutella/hap/crypto"
    "github.com/brutella/hap/common"
    
    "fmt"
    "encoding/hex"
    "bytes"
)

type VerifyServerController struct {
    context *hap.Context
    bridge *hap.Bridge
    session *VerifySession
    curSeq byte
}

func NewVerifyServerController(context *hap.Context, bridge *hap.Bridge) (*VerifyServerController, error) {    
    controller := VerifyServerController{
                                    context: context,
                                    bridge: bridge,
                                    session: NewVerifySession(),
                                    curSeq: WaitingForRequest,
                                }
    
    return &controller, nil
}
    
func (c *VerifyServerController) Handle(cont_in Container) (Container, error) {
    var cont_out Container
    var err error
    
    method := cont_in.GetByte(TLVType_Method)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, common.NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := cont_in.GetByte(TLVType_SequenceNumber)
    
    switch seq {
    case VerifyStartRequest:
        if c.curSeq != WaitingForRequest {
            c.reset()
            return nil, common.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        cont_out, err = c.handlePairVerifyStart(cont_in)
    case VerifyFinishRequest:
        if c.curSeq != VerifyStartRespond {
            c.reset()
            return nil, common.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        cont_out, err = c.handlePairVerifyFinish(cont_in)
    default:
        return nil, common.NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    return cont_out, err
}

// Client -> Server
// - Public key `A`
//
// Server
// - Create public `B` and secret key `S` based on `A`

// Server -> Client
// - B: server public key
// - signature: from server session public key, server name, client session public key
func (c *VerifyServerController) handlePairVerifyStart(cont_in Container) (Container, error) {
    c.curSeq = VerifyStartRespond
    
    clientPublicKey := cont_in.GetBytes(TLVType_PublicKey)
    fmt.Println("->     A:", hex.EncodeToString(clientPublicKey))
    if len(clientPublicKey) != 32 {
        return nil, common.NewErrorf("Invalid client public key size %d", len(clientPublicKey))
    }
    
    var otherPublicKey [32]byte
    copy(otherPublicKey[:], clientPublicKey)
    
    c.session.GenerateSharedKeyWithOtherPublicKey(otherPublicKey)
    c.session.SetupEncryptionKey([]byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))
    
    LTSK := c.context.SecretKeyForAccessory(c.bridge)
    
    material := make([]byte, 0)
    material = append(material, c.session.publicKey[:]...)
    material = append(material, c.bridge.Id()...)
    material = append(material, clientPublicKey...)
    signature, _ := crypto.ED25519Signature(LTSK, material)
    
    // Encrypt
    tlv_encrypt := NewTLV8Container()
    tlv_encrypt.SetString(TLVType_Username, c.bridge.Id())
    tlv_encrypt.SetBytes(TLVType_Ed25519Signature, signature)
    
    encrypted, mac, _ := crypto.Chacha20EncryptAndPoly1305Seal(c.session.encryptionKey[:], []byte("PV-Msg02"), tlv_encrypt.BytesBuffer().Bytes(), nil)
    
    cont_out := NewTLV8Container()    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    cont_out.SetBytes(TLVType_PublicKey, c.session.publicKey[:])
    cont_out.SetBytes(TLVType_EncryptedData, append(encrypted, mac[:]...))
    
    fmt.Println("       K:", hex.EncodeToString(c.session.encryptionKey[:]))
    fmt.Println("       B:", hex.EncodeToString(c.session.publicKey[:]))
    fmt.Println("       S:", hex.EncodeToString(c.session.secretKey[:]))
    fmt.Println("  Shared:", hex.EncodeToString(c.session.sharedKey[:]))
    
    fmt.Println("<-     B:", hex.EncodeToString(cont_out.GetBytes(TLVType_PublicKey)))
    
    return cont_out, nil
}

// Client -> Server
// - Encrypted tlv8: username and signature
//
// Server enrypty tlv8 and validates signature

// Server -> Client
// - only sequence number
// - error code (optional)
func (c *VerifyServerController) handlePairVerifyFinish(cont_in Container) (Container, error) {
    c.curSeq = VerifyFinishRespond
    
    data := cont_in.GetBytes(TLVType_EncryptedData)
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(c.session.encryptionKey[:], []byte("PV-Msg03"), message, mac, nil)
    
    cont_out := NewTLV8Container()    
    cont_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    if err != nil {
        c.reset()
        fmt.Println(err)
        cont_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        cont_in, err := NewTLV8ContainerFromReader(decrypted_buffer)
        if err != nil {
            return nil, err
        }
        
        username  := cont_in.GetString(TLVType_Username)
        signature := cont_in.GetBytes(TLVType_Ed25519Signature)
        fmt.Println("    client:", username)
        fmt.Println(" signature:", hex.EncodeToString(signature))
        
        client := c.context.ClientForName(username)
        if client == nil {
            return nil, common.NewErrorf("Client %s is unknown", username)
        }
        
        if len(client.PublicKey) == 0 {
            return nil, common.NewErrorf("No LTPK available for client %s", username)
        }
        
        material := make([]byte, 0)
        material = append(material, c.session.otherPublicKey[:]...)
        // TODO Report that material does not include username in docs
        material = append(material, []byte(username)...)
        material = append(material, c.session.publicKey[:]...)
        
        if crypto.ValidateED25519Signature(client.PublicKey, material, signature) == false {
            fmt.Println("[Failed] signature is invalid")
            c.reset()
            cont_out.SetByte(TLVType_ErrorCode, TLVStatus_UnkownPeerError) // return error 4
        } else {
            fmt.Println("[Success] signature is valid")
            
            // Verification is done
            // Switch to secure session
            secSession, err := crypto.NewSecureSessionFromSharedKey(c.session.sharedKey)
            if err != nil {
                fmt.Println("Could not setup secure session.", err)
            } else {
                fmt.Println("Setup secure session")
            }
            c.context.SetSecureSession(secSession)
            c.reset()
        }
    }
    
    return cont_out, nil
}

func (c *VerifyServerController) reset() {
    c.curSeq = WaitingForRequest
}