package gohap

import(
    "io"
)

const (
    SequenceStartRequest       = 1
    SequenceStartRespond       = 2
    SequenceVerifyRequest      = 3
    SequenceVerifyRespond      = 4
    SequenceKeyExchangeRequest = 5
    SequenceKeyExchangeRepond  = 6
)

type PairingController struct {
    session *PairingSession
}

func NewPairingController(username string, password string) (*PairingController, error) {
    if len(username) == 0 || len(password) == 0 {
        return nil, NewErrorf("Invalid username %s or password %s", username, password)
    }
    
    session, err := NewPairingSession(username, password)
    if err != nil {
        return nil, err
    }
    
    controller := PairingController{session: session}
    
    return &controller, nil
}

func (c *PairingController) Handle(r io.Reader) (io.Reader, error) {
    container, err := ReadTLV8(r)
    
    if err != nil {
        return nil, err
    }
    method := container.GetUInt64(TLVType_AuthMethod)
    
    // It is valid that method is not sent
    // If method is sent then it must be 0x00
    if method != 0 {
        return nil, NewErrorf("Cannot handle auth method %d", method)
    }
    
    seq := container.GetUInt64(TLVType_SequenceNumber)
    
    switch seq {
    case SequenceStartRequest:
        // Return random salt `s` and public key `B`
        tlv := TLV8Container{}
        tlv.Set(TLVType_Salt, c.session.Salt())
        tlv.Set(TLVType_PublicKey, c.session.PublicKey())
        tlv.Set(TLVType_SequenceNumber, []byte{SequenceStartRespond})
        return tlv.BytesBuffer(), nil
        
    case SequenceVerifyRequest:
    case SequenceKeyExchangeRequest:
    default:
        return nil, NewErrorf("Cannot handle sequence number %d", seq)
    }
    
    return nil, NewErrorf("Not handled")
}