package pair

import(
    "github.com/brutella/hap"
    "fmt"
)

/*
{
  guestName: <string>,
  guestPublicKey: <string>
}
*/
type Pairing struct {
    GuestName string `json:"guestName"`
    GuestPublicKey string `json:"guestPublicKey"`
}

type PairingController struct {
    context *hap.Context
}

func NewPairingController(context *hap.Context) *PairingController {
    c := PairingController{
        context: context,
    }
    
    return &c
}

func (c *PairingController) Handle(tlv8 *TLV8Container) error {
    seq         := tlv8.Byte(TLVType_SequenceNumber)
    method      := tlv8.Byte(TLVType_Method)
    username    := tlv8.String(TLVType_Username)
    publicKey   := tlv8.Bytes(TLVType_PublicKey)
    
    fmt.Println("->      Seq:", seq)
    fmt.Println("->   Method:", method)
    fmt.Println("-> Username:", username)
    fmt.Println("->     LTPK:", publicKey)
    
    client := hap.NewClient(username, publicKey)
    
    switch method {
    case TLVType_Method_PairingDelete:
        c.context.DeleteClient(client)
    case TLVType_Method_PairingAdd:
        err := c.context.SaveClient(client)
        if err != nil {
            fmt.Println(err)
            return err
        }
    }
    
    return nil
}