package pair

import(
    _"github.com/brutella/hap"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/db"
    
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
    database *db.Manager
}

func NewPairingController(database *db.Manager) *PairingController {
    c := PairingController{
        database: database,
    }
    
    return &c
}

func (c *PairingController) Handle(tlv8 common.Container) (common.Container, error) {
    method      := tlv8.GetByte(TLVType_Method)
    username    := tlv8.GetString(TLVType_Username)
    publicKey   := tlv8.GetBytes(TLVType_PublicKey)
    
    fmt.Println("->   Method:", method)
    fmt.Println("-> Username:", username)
    fmt.Println("->     LTPK:", publicKey)
    
    client := db.NewClient(username, publicKey)
    
    switch method {
    case TLVType_Method_PairingDelete:
        c.database.DeleteClient(client)
    case TLVType_Method_PairingAdd:
        err := c.database.SaveClient(client)
        if err != nil {
            fmt.Println(err)
            return nil, err
        }
    }
    
    return nil, nil
}