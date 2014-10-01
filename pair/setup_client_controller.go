package pair

import(
    "github.com/brutella/hap"
    
    "io"
    "fmt"
    "encoding/hex"
    "bytes"
)

type SetupClientController struct {
    PairingHandler
    context *hap.Context
    bridge *hap.Bridge
    username string
    session *SetupClientSession
}

func NewSetupClientController(context *hap.Context, bridge *hap.Bridge, username string) (*SetupClientController) {
    
    session := NewSetupClientSession("Pair-Setup", bridge.Password())
    
    controller := SetupClientController{
                                    username: username,
                                    context: context,
                                    bridge: bridge,
                                    session: session,
                                }
    return &controller
}

func (c *SetupClientController) InitialPairingRequest() (io.Reader) {
    tlvPairStart := NewTLV8Container()
    tlvPairStart.SetByte(TLVType_Method, 0)
    tlvPairStart.SetByte(TLVType_SequenceNumber, PairStartRequest)
    
    return tlvPairStart.BytesBuffer()
}

func (c *SetupClientController) HandleReader(r io.Reader) (io.Reader, error) {
    tlv_in, err := NewTLV8ContainerFromReader(r)
    if err != nil {
        return nil, err
    }
    fmt.Println("->     Seq:", tlv_in.GetByte(TLVType_SequenceNumber))
    
    tlv_out, err := c.Handle(tlv_in)
    
    if err != nil {
        fmt.Println("[ERROR]", err)
        return nil, err
    } else {
        if tlv_out != nil {
            fmt.Println("<-     Seq:", tlv_out.GetByte(TLVType_SequenceNumber))
            fmt.Println("-------------")
            return tlv_out.BytesBuffer(), nil
        }
    }
    
    return nil, err
}

func (c *SetupClientController) Handle(tlv_in Container) (Container, error) {
    method := tlv_in.GetByte(TLVType_Method)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, hap.NewErrorf("Cannot handle auth method %b", method)
    }
    
    err_code := tlv_in.GetByte(TLVType_ErrorCode)
    if err_code != 0x00 {
        return nil, hap.NewErrorf("Received error %d", err_code)
    }
    
    seq := tlv_in.GetByte(TLVType_SequenceNumber)
    fmt.Println("->     Seq:", seq)
    
    var tlv_out Container
    var err error
    
    switch seq {
    case PairStartRespond:        
        tlv_out, err = c.handlePairStartRespond(tlv_in)
    case PairVerifyRespond:
        tlv_out, err = c.handlePairVerifyRespond(tlv_in)
    case PairKeyExchangeRespond:        
        tlv_out, err = c.handleKeyExchange(tlv_in)
    default:
        return nil, hap.NewErrorf("Cannot handle sequence number %d", seq)
    }
        
    return tlv_out, err
}

// Server -> Client
// - B: server public key
// - s: salt
//
// Client -> Server
// - A: client public key
// - M1: proof
func (c *SetupClientController) handlePairStartRespond(tlv_in Container) (Container, error) {    
    salt := tlv_in.GetBytes(TLVType_Salt)
    serverPublicKey := tlv_in.GetBytes(TLVType_PublicKey)
    
    if len(salt) != 16 {
        return nil, hap.NewErrorf("Salt is invalid (%d bytes)", len(salt))
    }
    
    if len(serverPublicKey) != 384 {
        return nil, hap.NewErrorf("B is invalid (%d bytes)", len(serverPublicKey))
    }
    
    fmt.Println("->     B:", hex.EncodeToString(serverPublicKey))
    fmt.Println("->     s:", hex.EncodeToString(salt))
    
    // Client
    // 1) Receive salt `s` and public key `B` and generates `S` and `A`
    err := c.session.GenerateKeys(salt, serverPublicKey)
    if err != nil {
        return nil, err
    }
    fmt.Println("        S:", hex.EncodeToString(c.session.secretKey))
    
    // 2) Send public key `A` and proof `M1`
    publicKey := c.session.publicKey // SRP public key
    proof := c.session.proof // M1
    
    fmt.Println("<-     A:", hex.EncodeToString(publicKey))
    fmt.Println("<-     M1:", hex.EncodeToString(proof))
    
    tlv_out := NewTLV8Container()
    tlv_out.SetByte(TLVType_Method, 0)
    tlv_out.SetByte(TLVType_SequenceNumber, PairVerifyRequest)
    tlv_out.SetBytes(TLVType_PublicKey, publicKey)
    tlv_out.SetBytes(TLVType_Proof, proof)
    
    return tlv_out, nil
}

// Client -> Server
// - A: client public key
// - M1: proof
// 
// Server -> client
// - M2: proof
// or
// - auth error
func (c *SetupClientController) handlePairVerifyRespond(tlv_in Container) (Container, error) {
    serverProof := tlv_in.GetBytes(TLVType_Proof)
    fmt.Println("->     M2:", hex.EncodeToString(serverProof))
    
    if c.session.IsServerProofValid(serverProof) == false {
        return nil, hap.NewErrorf("M2 %s is invalid", hex.EncodeToString(serverProof))
    }
    
    err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
    if err != nil {
        return nil, err
    }
    
    fmt.Println("        K:", hex.EncodeToString(c.session.encryptionKey[:]))
    
    // 2) Send username, LTPK, signature as encrypted message
    H, err := hap.HKDF_SHA512(c.session.secretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
    material := make([]byte, 0)
    material = append(material, H[:]...)
    material = append(material, c.username...)
    material = append(material, c.session.LTPK...)
    
    signature, err := hap.ED25519Signature(c.session.LTSK, material)
    if err != nil {
        return nil, err
    }
    
    tlvPairKeyExchange := NewTLV8Container()
    tlvPairKeyExchange.SetString(TLVType_Username, c.username)
    tlvPairKeyExchange.SetBytes(TLVType_PublicKey, []byte(c.session.LTPK))
    tlvPairKeyExchange.SetBytes(TLVType_Ed25519Signature, []byte(signature))
    
    encrypted, tag, err := hap.Chacha20EncryptAndPoly1305Seal(c.session.encryptionKey[:], []byte("PS-Msg05"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)
    if err != nil {
        return nil, err
    }
    
    tlv_out := NewTLV8Container()
    tlv_out.SetByte(TLVType_Method, 0)
    tlv_out.SetByte(TLVType_SequenceNumber, PairKeyExchangeRequest)
    tlv_out.SetBytes(TLVType_EncryptedData, append(encrypted, tag[:]...))
    
    fmt.Println("<-   Encrypted:", hex.EncodeToString(tlv_out.GetBytes(TLVType_EncryptedData)))
    
    return tlv_out, nil
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
func (c *SetupClientController) handleKeyExchange(tlv_in Container) (Container, error) {
    data := tlv_in.GetBytes(TLVType_EncryptedData)    
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := hap.Chacha20DecryptAndPoly1305Verify(c.session.encryptionKey[:], []byte("PS-Msg06"), message, mac, nil)
    
    if err != nil {
        fmt.Println(err)
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        tlv_in, err := NewTLV8ContainerFromReader(decrypted_buffer)
        if err != nil {
            fmt.Println(err)
        }
        
        username  := tlv_in.GetString(TLVType_Username)
        ltpk      := tlv_in.GetBytes(TLVType_PublicKey)
        signature := tlv_in.GetBytes(TLVType_Ed25519Signature)
        fmt.Println("->     Username:", username)
        fmt.Println("->     LTPK:", hex.EncodeToString(ltpk))
        fmt.Println("->     Signature:", hex.EncodeToString(signature))
    }
    
    return nil, err
}