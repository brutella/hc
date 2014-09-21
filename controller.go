package gohap

import(
    "io"
    "fmt"
    "encoding/hex"
    "bytes"
)

const (
    SequenceWaitingForRequest  = 0x00
    SequenceStartRequest       = 0x01
    SequenceStartRespond       = 0x02
    SequenceVerifyRequest      = 0x03
    SequenceVerifyRespond      = 0x04
    SequenceKeyExchangeRequest = 0x05
    SequenceKeyExchangeRepond  = 0x06
)

type PairingController struct {
    accessory *Accessory
    session *PairingSession
    curSeq byte
}

func NewPairingController(accessory *Accessory) (*PairingController, error) {
    
    session, err := NewPairingSession("Pair-Setup", accessory.password)
    if err != nil {
        return nil, err
    }
    
    controller := PairingController{
                                    accessory: accessory,
                                    session: session,
                                    curSeq: SequenceWaitingForRequest,
                                }
    
    return &controller, nil
}

func (c *PairingController) Handle(r io.Reader) (io.Reader, error) {
    container, err := ReadTLV8(r)
    
    if err != nil {
        return nil, err
    }
    method := container.GetByte(TLVType_AuthMethod)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0x00 {
        return nil, NewErrorf("Cannot handle auth method %b", method)
    }
    
    seq := container.GetByte(TLVType_SequenceNumber)
    
    switch seq {
    case SequenceStartRequest:
        fmt.Println("Receive pair start request...")
        // Client -> Server
        // - Auth start
        
        if c.curSeq != SequenceWaitingForRequest {
            return nil, NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        c.curSeq = SequenceStartRespond
        
        // Return random salt `s` and public key `B`
        
        // Server -> client
        // - B: server public key
        // - s: salt
        tlv_out := TLV8Container{}
        tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
        tlv_out.SetBytes(TLVType_Salt, c.session.Salt())
        tlv_out.SetBytes(TLVType_PublicKey, c.session.PublicKey())
        
        fmt.Println(" <- Seq:", tlv_out.GetByte(TLVType_SequenceNumber))
        fmt.Println(" <-   B:", hex.EncodeToString(tlv_out.GetBytes(TLVType_PublicKey)))
        fmt.Println(" <-   s:", hex.EncodeToString(tlv_out.GetBytes(TLVType_Salt)))
        fmt.Println("-------------")
        
        return tlv_out.BytesBuffer(), nil
    case SequenceVerifyRequest:
        fmt.Println("Receive pair verify request...")
        
        // Client -> Server
        // - A: client public key
        // - M1: proof
        if c.curSeq != SequenceStartRespond {
            return nil, NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        
        c.curSeq = SequenceVerifyRespond
        
        
        // Server -> client
        // - M2: proof
        // or
        // - auth error
        tlv_out := TLV8Container{}
        tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
        
        cpublicKey := container.GetBytes(TLVType_PublicKey)
        fmt.Println(" -> A:", hex.EncodeToString(cpublicKey))
        
        err := c.session.SetupSecretKeyFromClientPublicKey(cpublicKey)
        if err != nil {
            return nil, err
        }
        
        cproof := container.GetBytes(TLVType_Proof)
        fmt.Println(" -> M1:", hex.EncodeToString(cproof))
        
        sproof, err := c.session.ProofFromClientProof(cproof)
        if err != nil || len(sproof) == 0 { // proof `M1` is wrong
            c.reset()
            // Return auth error
            tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError)
        } else {
            err := c.session.SetupEncryptionKey([]byte("Pair-Setup-Encrypt-Salt"), []byte("Pair-Setup-Encrypt-Info"))
            if err != nil {
                return nil, err
            }
            
            // Return proof `M1`
            tlv_out.SetBytes(TLVType_Proof, sproof)
        }
        
        fmt.Println(" <- Seq:", tlv_out.GetByte(TLVType_SequenceNumber))
        fmt.Println(" <-  M2:", hex.EncodeToString(tlv_out.GetBytes(TLVType_Proof)))
        fmt.Println("      S:", hex.EncodeToString(c.session.secretKey))
        fmt.Println("      K:", hex.EncodeToString(c.session.encryptionKey))
        fmt.Println("-------------")
        return tlv_out.BytesBuffer(), nil
    case SequenceKeyExchangeRequest:
        fmt.Println("Receive pair key exchange request...")
        
        // Client -> Server
        // - encrypted message
        // - auth tag
        if c.curSeq != SequenceVerifyRespond {
            return nil, NewErrorf("Controller is in wrong state (%d)", c.curSeq)
        }
        c.curSeq = SequenceKeyExchangeRepond
        tlv_out := TLV8Container{}
        tlv_out.SetByte(TLVType_SequenceNumber, c.curSeq)
        
        data := container.GetBytes(TLVType_EncryptedData)
        message := data[:(len(data) - 16)]
        authTag := data[len(message):] // 16 byte (MAC)
        fmt.Println("  -> message:", hex.EncodeToString(message))
        fmt.Println(" -> auth tag:", hex.EncodeToString(authTag))
        
        decrypted, err := DecryptAndVerify(c.session.encryptionKey, []byte("PS-Msg05"), message, authTag, nil)
        
        if err != nil {
            c.reset()
            fmt.Println(err)
            // Return auth error
            tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_UnkownError) // send error 1
        } else {
            decrypted_buffer := bytes.NewBuffer(decrypted)
            tlv_in, err := ReadTLV8(decrypted_buffer)
            if err != nil {
                return nil, err
            }
            
            username  := tlv_in.GetBytes(TLVType_Username)
            ltpk      := tlv_in.GetBytes(TLVType_PublicKey)
            signature := tlv_in.GetBytes(TLVType_Ed25519Signature)
            fmt.Println("  -> Username:", string(username))
            fmt.Println("  -> LTPK:", hex.EncodeToString(ltpk))
            fmt.Println("  -> Signature:", hex.EncodeToString(signature))
            
            // Calculate `H`
            H, _ := HKDF_SHA512_256(c.session.secretKey, []byte("Pair-Setup-Controller-Sign-Salt"), []byte("Pair-Setup-Controller-Sign-Info"))
            material := make([]byte, 0)
            material = append(material, H...)
            material = append(material, username...)
            material = append(material, ltpk...)
            
            if ValidateED25519Signature(ltpk, material, signature) == false {
                c.reset()
                fmt.Println("[Failed] ed25519 signature is invalid")
                // Return auth error
                tlv_out.SetByte(TLVType_ErrorCode, TLVStatus_AuthError) // send error 2
            } else {
                fmt.Println("[Success] ed25519 signature is valid")
                // TODO Store ltpk, username and H
                
                // Send username, LTPK, proof as encrypted message
                H2, err := HKDF_SHA512_256(c.session.secretKey, []byte("Pair-Setup-Accessory-Sign-Salt"), []byte("Pair-Setup-Accessory-Sign-Info"))
                material = make([]byte, 0)
                material = append(material, H2...)
                material = append(material, []byte(c.accessory.name)...)
                material = append(material, c.accessory.publicKey...) // LTPK
    
                signature, err := ED25519Signature(c.accessory.secretKey, material)
                if err != nil {
                    return nil, err
                }
                
                tlvPairKeyExchange := TLV8Container{}
                tlvPairKeyExchange.SetBytes(TLVType_Username, []byte(c.accessory.name))
                tlvPairKeyExchange.SetBytes(TLVType_PublicKey, []byte(c.accessory.publicKey))
                tlvPairKeyExchange.SetBytes(TLVType_Proof, []byte(signature))
                
                var tag [16]byte
                encrypted, tag, _ := EncryptAndSeal(c.session.encryptionKey, []byte("PS-Msg06"), tlvPairKeyExchange.BytesBuffer().Bytes(), tag, nil)    
                tlv_out.SetByte(TLVType_AuthMethod, 0)
                tlv_out.SetByte(TLVType_SequenceNumber, SequenceKeyExchangeRequest)
                tlv_out.SetBytes(TLVType_EncryptedData, append(encrypted, tag[:]...))
            }
        }
        
        return tlv_out.BytesBuffer(), nil
    default:
        return nil, NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    return nil, NewErrorf("Not handled")
}

func (c *PairingController) reset() {
    c.curSeq = SequenceWaitingForRequest
    // TODO: reset session
}