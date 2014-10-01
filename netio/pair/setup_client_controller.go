package pair

import(
    "github.com/brutella/hap"
    "github.com/brutella/hap/crypto"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/netio"
    
    "io"
    "fmt"
    "encoding/hex"
    "bytes"
)

type SetupClientController struct {
    Handler
    context *hap.Context
    bridge *hap.Bridge
    username string
    session *netio.PairSetupClientSession
}

func NewSetupClientController(context *hap.Context, bridge *hap.Bridge, username string) (*SetupClientController) {
    
    session := netio.NewPairSetupClientSession("Pair-Setup", bridge.Password())
    
    controller := SetupClientController{
                                    username: username,
                                    context: context,
                                    bridge: bridge,
                                    session: session,
                                }
    return &controller
}

func (c *SetupClientController) InitialPairingRequest() (io.Reader) {
    tlvPairStart := common.NewTLV8Container()
    tlvPairStart.SetByte(TLVType_Method, 0)
    tlvPairStart.SetByte(TLVType_SequenceNumber, PairStartRequest)
    
    return tlvPairStart.BytesBuffer()
}

func (c *SetupClientController) Handle(cont_in common.Container) (common.Container, error) {
    method := cont_in.GetByte(TLVType_Method)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, common.NewErrorf("Cannot handle auth method %b", method)
    }
    
    err_code := cont_in.GetByte(TLVType_ErrorCode)
    if err_code != 0x00 {
        return nil, common.NewErrorf("Received error %d", err_code)
    }
    
    seq := cont_in.GetByte(TLVType_SequenceNumber)
    fmt.Println("->     Seq:", seq)
    
    var cont_out common.Container
    var err error
    
    switch seq {
    case PairStartRespond:        
        cont_out, err = c.handlePairStartRespond(cont_in)
    case PairVerifyRespond:
        cont_out, err = c.handlePairVerifyRespond(cont_in)
    case PairKeyExchangeRespond:        
        cont_out, err = c.handleKeyExchange(cont_in)
    default:
        return nil, common.NewErrorf("Cannot handle sequence number %d", seq)
    }
        
    return cont_out, err
}

// Server -> Client
// - B: server public key
// - s: salt
//
// Client -> Server
// - A: client public key
// - M1: proof
func (c *SetupClientController) handlePairStartRespond(cont_in common.Container) (common.Container, error) {    
    salt := cont_in.GetBytes(TLVType_Salt)
    serverPublicKey := cont_in.GetBytes(TLVType_PublicKey)
    
    if len(salt) != 16 {
        return nil, common.NewErrorf("Salt is invalid (%d bytes)", len(salt))
    }
    
    if len(serverPublicKey) != 384 {
        return nil, common.NewErrorf("B is invalid (%d bytes)", len(serverPublicKey))
    }
    
    fmt.Println("->     B:", hex.EncodeToString(serverPublicKey))
    fmt.Println("->     s:", hex.EncodeToString(salt))
    
    // Client
    // 1) Receive salt `s` and public key `B` and generates `S` and `A`
    err := c.session.GenerateKeys(salt, serverPublicKey)
    if err != nil {
        return nil, err
    }
    fmt.Println("        S:", hex.EncodeToString(c.session.SecretKey))
    
    // 2) Send public key `A` and proof `M1`
    publicKey := c.session.PublicKey // SRP public key
    proof := c.session.Proof // M1
    
    fmt.Println("<-     A:", hex.EncodeToString(publicKey))
    fmt.Println("<-     M1:", hex.EncodeToString(proof))
    
    cont_out := common.NewTLV8Container()
    cont_out.SetByte(TLVType_Method, 0)
    cont_out.SetByte(TLVType_SequenceNumber, PairVerifyRequest)
    cont_out.SetBytes(TLVType_PublicKey, publicKey)
    cont_out.SetBytes(TLVType_Proof, proof)
    
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
func (c *SetupClientController) handlePairVerifyRespond(cont_in common.Container) (common.Container, error) {
    serverProof := cont_in.GetBytes(TLVType_Proof)
    fmt.Println("->     M2:", hex.EncodeToString(serverProof))
    
    if c.session.IsServerProofValid(serverProof) == false {
        return nil, common.NewErrorf("M2 %s is invalid", hex.EncodeToString(serverProof))
    }
    
    err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
    if err != nil {
        return nil, err
    }
    
    fmt.Println("        K:", hex.EncodeToString(c.session.EncryptionKey[:]))
    
    // 2) Send username, LTPK, signature as encrypted message
    H, err := crypto.HKDF_SHA512(c.session.SecretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
    material := make([]byte, 0)
    material = append(material, H[:]...)
    material = append(material, c.username...)
    material = append(material, c.session.LTPK...)
    
    signature, err := crypto.ED25519Signature(c.session.LTSK, material)
    if err != nil {
        return nil, err
    }
    
    tlvPairKeyExchange := common.NewTLV8Container()
    tlvPairKeyExchange.SetString(TLVType_Username, c.username)
    tlvPairKeyExchange.SetBytes(TLVType_PublicKey, []byte(c.session.LTPK))
    tlvPairKeyExchange.SetBytes(TLVType_Ed25519Signature, []byte(signature))
    
    encrypted, tag, err := crypto.Chacha20EncryptAndPoly1305Seal(c.session.EncryptionKey[:], []byte("PS-Msg05"), tlvPairKeyExchange.BytesBuffer().Bytes(), nil)
    if err != nil {
        return nil, err
    }
    
    cont_out := common.NewTLV8Container()
    cont_out.SetByte(TLVType_Method, 0)
    cont_out.SetByte(TLVType_SequenceNumber, PairKeyExchangeRequest)
    cont_out.SetBytes(TLVType_EncryptedData, append(encrypted, tag[:]...))
    
    fmt.Println("<-   Encrypted:", hex.EncodeToString(cont_out.GetBytes(TLVType_EncryptedData)))
    
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
func (c *SetupClientController) handleKeyExchange(cont_in common.Container) (common.Container, error) {
    data := cont_in.GetBytes(TLVType_EncryptedData)    
    message := data[:(len(data) - 16)]
    var mac [16]byte
    copy(mac[:], data[len(message):]) // 16 byte (MAC)
    fmt.Println("->     Message:", hex.EncodeToString(message))
    fmt.Println("->     MAC:", hex.EncodeToString(mac[:]))
    
    decrypted, err := crypto.Chacha20DecryptAndPoly1305Verify(c.session.EncryptionKey[:], []byte("PS-Msg06"), message, mac, nil)
    
    if err != nil {
        fmt.Println(err)
    } else {
        decrypted_buffer := bytes.NewBuffer(decrypted)
        cont_in, err := common.NewTLV8ContainerFromReader(decrypted_buffer)
        if err != nil {
            fmt.Println(err)
        }
        
        username  := cont_in.GetString(TLVType_Username)
        ltpk      := cont_in.GetBytes(TLVType_PublicKey)
        signature := cont_in.GetBytes(TLVType_Ed25519Signature)
        fmt.Println("->     Username:", username)
        fmt.Println("->     LTPK:", hex.EncodeToString(ltpk))
        fmt.Println("->     Signature:", hex.EncodeToString(signature))
    }
    
    return nil, err
}