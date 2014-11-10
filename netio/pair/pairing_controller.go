package pair

import(
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/db"
    "github.com/brutella/log"
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
    database db.Database
}

func NewPairingController(database db.Database) *PairingController {
    c := PairingController{
        database: database,
    }
    
    return &c
}

func (c *PairingController) Handle(tlv8 common.Container) (common.Container, error) {
    method      := tlv8.GetByte(TLVType_Method)
    username    := tlv8.GetString(TLVType_Username)
    publicKey   := tlv8.GetBytes(TLVType_PublicKey)
    
    log.Println("[INFO] ->   Method:", method)
    log.Println("[INFO] -> Username:", username)
    log.Println("[INFO] ->     LTPK:", publicKey)
    
    client := db.NewClient(username, publicKey)
    
    switch method {
    case TLVType_Method_PairingDelete:
        c.database.DeleteClient(client)
    case TLVType_Method_PairingAdd:
        err := c.database.SaveClient(client)
        if err != nil {
            log.Println("[ERROR]", err)
            return nil, err
        }
    }
    
    return nil, nil
}