package gohap

import(
    "io"
    "fmt"
    "encoding/hex"
    "bytes"
)

const (
    WaitingForRequest   = 0x00
    VerifyStartRequest  = 0x01
    VerifyStartRespond  = 0x02
    VerifyFinishRequest = 0x03
    VerifyFinishRespond = 0x04
)

type VerifyController struct {
    context *Context
    accessory *Accessory
    session *PairVerifySession
    curSeq byte
}

func NewVerifyController(context *Context, accessory *Accessory) (*VerifyController, error) {    
    controller := VerifyController{
                                    context: context,
                                    accessory: accessory,
                                    session: NewPairVerifySession(),
                                    curSeq: WaitingForRequest,
                                }
    
    return &controller, nil
}

func (c *VerifyController) Handle(r io.Reader) (io.Reader, error) {
    var tlv_out *TLV8Container
    var err error
    
    tlv_in, err := ReadTLV8(r)
    if err != nil {
        return nil, err
    }
    
    method := tlv_in.GetByte(TLVType_AuthMethod)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := tlv_in.GetByte(TLVType_SequenceNumber)
    fmt.Println("->     Seq:", seq)
    
    switch seq {
    case VerifyStartRequest:
        if c.curSeq != WaitingForRequest {
            return nil, NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        tlv_out, err = c.handlePairVerifyStart(tlv_in)
    case VerifyFinishRequest:
        if c.curSeq != VerifyStartRespond {
            return nil, NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        tlv_out, err = c.handlePairVerifyFinish(tlv_in)
    default:
        return nil, NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    fmt.Println("<-     Seq:", tlv_out.GetByte(TLVType_SequenceNumber))
    fmt.Println("-------------")
    return tlv_out.BytesBuffer(), err
}

// Client -> Server
// - Public key `A`
//
// Server
// - Create public `B` and secret key `S` based on `A`

// Server -> Client
// - B: server public key
// - signature: from server publickey, server name, client public key
func (c *VerifyController) handlePairVerifyStart(tlv_in *TLV8Container) (*TLV8Container, error) {
    c.curSeq = VerifyStartRespond
    
    clientPublicKey := tlv_in.GetBytes(TLVType_PublicKey)
    fmt.Println("->     A:", hex.EncodeToString(clientPublicKey))
    if len(clientPublicKey) != 32 {
        return nil, NewErrorf("Invalid client public key size %d", len(clientPublicKey))
    }
    var otherPublicKey [32]byte
    copy(otherPublicKey[:], clientPublicKey)
    
    c.session.GenerateKeysWithOtherPublicKey(otherPublicKey)
    
    material := make([]byte, 0)
    material = append(material, c.session.publicKey[:]...)
    material = append(material, c.accessory.name...)
    material = append(material, clientPublicKey...)
    signature, _ := ED25519Signature(c.accessory.secretKey, material)
    
    K, _ := HKDF_SHA512(c.session.sharedKey[:], []byte("Pair-Verify-Encrypt-Salt"), []byte("Pair-Verify-Encrypt-Info"))
    c.session.encryptionKey = K
    
    // Encrypt
    tlv_encrypt := TLV8Container{}
    tlv_encrypt.SetString(TLVType_Username, c.accessory.name)
    tlv_encrypt.SetBytes(TLVType_Ed25519Signature, signature)
    
    var mac [16]byte
    encrypted, mac, _ := Chacha20EncryptAndPoly1305Seal(c.session.encryptionKey[:], []byte("PV-Msg02"), tlv_encrypt.BytesBuffer().Bytes(), mac, nil)
    
    tlv_out := TLV8Container{}    
    tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    tlv_out.SetBytes(TLVType_PublicKey, c.session.publicKey[:])
    tlv_out.SetBytes(TLVType_EncryptedData, append(encrypted, mac[:]...))
    
    fmt.Println("       K:", hex.EncodeToString(K[:]))
    fmt.Println("       B:", hex.EncodeToString(c.session.publicKey[:]))
    fmt.Println("       S:", hex.EncodeToString(c.session.secretKey[:]))
    
    fmt.Println("<-     B:", hex.EncodeToString(tlv_out.GetBytes(TLVType_PublicKey)))
    
    return &tlv_out, nil
}

// Client -> Server
// - Encrypted tlv8: username and signature
//
// Server
// - Decrypt tlv8 and validate signature

// Server -> Client
// - only sequence number
// - error code (on error)
func (c *VerifyController) handlePairVerifyFinish(tlv_in *TLV8Container) (*TLV8Container, error) {
    c.curSeq = VerifyFinishRespond
    
    data := tlv_in.GetBytes(TLVType_EncryptedData)
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := Chacha20DecryptAndPoly1305Verify(c.session.encryptionKey[:], []byte("PV-Msg03"), message, mac, nil)
    
    tlv_out := TLV8Container{}    
    tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
    
    if err != nil {
        c.reset()
        fmt.Println(err)
        tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // return error 2
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        tlv_in, err := ReadTLV8(decrypted_buffer)
        if err != nil {
            return nil, err
        }
        
        username  := tlv_in.GetString(TLVType_Username)
        signature := tlv_in.GetBytes(TLVType_Ed25519Signature)
        fmt.Println("    client:", username)
        fmt.Println(" signature:", hex.EncodeToString(signature))
        
        client := c.context.ClientForName(username)
        if client == nil {
            return nil, NewErrorf("Client %s is unknown", username)
        }
        
        if len(client.publicKey) == 0 {
            return nil, NewErrorf("No LTPK available for client %s", username)
        }
        
        material := make([]byte, 0)
        material = append(material, c.session.otherPublicKey[:]...)
        material = append(material, c.session.publicKey[:]...)
        
        if ValidateED25519Signature(client.publicKey, material, signature) == false {
            fmt.Println("[Failed] ed25519 signature is invalid")
            c.reset()
            tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_UnkownPeerError) // return error 4
        } else {
            fmt.Println("[Success] ed25519 signature is valid")
        }
    }
    
    return &tlv_out, nil
}

func (c *VerifyController) reset() {
    c.curSeq = WaitingForRequest
}