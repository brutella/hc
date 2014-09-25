package pair

import(
    "github.com/brutella/hap"
    "io"
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

func (c *VerifyServerController) Handle(r io.Reader) (io.Reader, error) {
    var tlv_out *hap.TLV8Container
    var err error
    
    tlv_in, err := hap.ReadTLV8(r)
    if err != nil {
        return nil, err
    }
    
    method := tlv_in.GetByte(hap.TLVType_AuthMethod)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, hap.NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := tlv_in.GetByte(hap.TLVType_SequenceNumber)
    fmt.Println("->     Seq:", seq)
    
    switch seq {
    case VerifyStartRequest:
        if c.curSeq != WaitingForRequest {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        tlv_out, err = c.handlePairVerifyStart(tlv_in)
    case VerifyFinishRequest:
        if c.curSeq != VerifyStartRespond {
            c.reset()
            return nil, hap.NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        tlv_out, err = c.handlePairVerifyFinish(tlv_in)
    default:
        return nil, hap.NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    if err != nil {
        fmt.Println("[ERROR]", err)
        return nil, err
    } else {
        fmt.Println("<-     Seq:", tlv_out.GetByte(hap.TLVType_SequenceNumber))
        fmt.Println("-------------")
        return tlv_out.BytesBuffer(), nil
    }
}

// Client -> Server
// - Public key `A`
//
// Server
// - Create public `B` and secret key `S` based on `A`

// Server -> Client
// - B: server public key
// - signature: from server session public key, server name, client session public key
func (c *VerifyServerController) handlePairVerifyStart(tlv_in *hap.TLV8Container) (*hap.TLV8Container, error) {
    c.curSeq = VerifyStartRespond
    
    clientPublicKey := tlv_in.GetBytes(hap.TLVType_PublicKey)
    fmt.Println("->     A:", hex.EncodeToString(clientPublicKey))
    if len(clientPublicKey) != 32 {
        return nil, hap.NewErrorf("Invalid client public key size %d", len(clientPublicKey))
    }
    
    var otherPublicKey [32]byte
    copy(otherPublicKey[:], clientPublicKey)
    
    c.session.GenerateSharedKeyWithOtherPublicKey(otherPublicKey)
    c.session.SetupEncryptionKey([]byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))
    
    LTSK := c.context.SecretKeyForAccessory(c.bridge)
    
    material := make([]byte, 0)
    material = append(material, c.session.publicKey[:]...)
    material = append(material, c.bridge.Name...)
    material = append(material, clientPublicKey...)
    signature, _ := hap.ED25519Signature(LTSK, material)
    
    // Encrypt
    tlv_encrypt := hap.TLV8Container{}
    tlv_encrypt.SetString(hap.TLVType_Username, c.bridge.Name)
    tlv_encrypt.SetBytes(hap.TLVType_Ed25519Signature, signature)
    
    encrypted, mac, _ := hap.Chacha20EncryptAndPoly1305Seal(c.session.encryptionKey[:], []byte("PV-Msg02"), tlv_encrypt.BytesBuffer().Bytes(), nil)
    
    tlv_out := hap.TLV8Container{}    
    tlv_out.SetByte(hap.TLVType_SequenceNumber, c.curSeq)
    tlv_out.SetBytes(hap.TLVType_PublicKey, c.session.publicKey[:])
    tlv_out.SetBytes(hap.TLVType_EncryptedData, append(encrypted, mac[:]...))
    
    fmt.Println("       K:", hex.EncodeToString(c.session.encryptionKey[:]))
    fmt.Println("       B:", hex.EncodeToString(c.session.publicKey[:]))
    fmt.Println("       S:", hex.EncodeToString(c.session.secretKey[:]))
    fmt.Println("  Shared:", hex.EncodeToString(c.session.sharedKey[:]))
    
    fmt.Println("<-     B:", hex.EncodeToString(tlv_out.GetBytes(hap.TLVType_PublicKey)))
    
    return &tlv_out, nil
}

// Client -> Server
// - Encrypted tlv8: username and signature
//
// Server enrypty tlv8 and validates signature

// Server -> Client
// - only sequence number
// - error code (optional)
func (c *VerifyServerController) handlePairVerifyFinish(tlv_in *hap.TLV8Container) (*hap.TLV8Container, error) {
    c.curSeq = VerifyFinishRespond
    
    data := tlv_in.GetBytes(hap.TLVType_EncryptedData)
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := hap.Chacha20DecryptAndPoly1305Verify(c.session.encryptionKey[:], []byte("PV-Msg03"), message, mac, nil)
    
    tlv_out := hap.TLV8Container{}    
    tlv_out.SetByte(hap.TLVType_SequenceNumber, c.curSeq)
    
    if err != nil {
        c.reset()
        fmt.Println(err)
        tlv_out.SetByte(hap.TLVType_ErrorCode, hap.TLVStatus_AuthError) // return error 2
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        tlv_in, err := hap.ReadTLV8(decrypted_buffer)
        if err != nil {
            return nil, err
        }
        
        username  := tlv_in.GetString(hap.TLVType_Username)
        signature := tlv_in.GetBytes(hap.TLVType_Ed25519Signature)
        fmt.Println("    client:", username)
        fmt.Println(" signature:", hex.EncodeToString(signature))
        
        client := c.context.ClientForName(username)
        if client == nil {
            return nil, hap.NewErrorf("Client %s is unknown", username)
        }
        
        if len(client.PublicKey) == 0 {
            return nil, hap.NewErrorf("No LTPK available for client %s", username)
        }
        
        material := make([]byte, 0)
        material = append(material, c.session.otherPublicKey[:]...)
        // TODO Report that material does not include username in docs
        material = append(material, []byte(username)...)
        material = append(material, c.session.publicKey[:]...)
        
        if hap.ValidateED25519Signature(client.PublicKey, material, signature) == false {
            fmt.Println("[Failed] signature is invalid")
            c.reset()
            tlv_out.SetByte(hap.TLVType_ErrorCode, hap.TLVStatus_UnkownPeerError) // return error 4
        } else {
            fmt.Println("[Success] signature is valid")
            // Verification is done now generation incoming and outgoing encryption keys
            c.context.GenerateEncryptionKeysWithSharedkey(c.session.sharedKey)
        }
    }
    
    return &tlv_out, nil
}

func (c *VerifyServerController) reset() {
    c.curSeq = WaitingForRequest
}