package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"
	"github.com/brutella/log"
)

// Implements pairing json of format
//     {
//       "guestName": <string>,
//       "guestPublicKey": <string>
//     }
type Pairing struct {
	GuestName      string `json:"guestName"`
	GuestPublicKey string `json:"guestPublicKey"`
}

// PairingController handles pairing with a client. The client's public key is stored in the database.
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
	method := tlv8.GetByte(TLVMethod)
	username := tlv8.GetString(TLVUsername)
	publicKey := tlv8.GetBytes(TLVPublicKey)

	log.Println("[VERB] ->   Method:", method)
	log.Println("[VERB] -> Username:", username)
	log.Println("[VERB] ->     LTPK:", publicKey)

	client := db.NewClient(username, publicKey)

	switch method {
	case MethodDelete:
		log.Printf("[INFO] Remove LTPK for client '%s'\n", username)
		c.database.DeleteClient(client)
	case MethodAdd:
		err := c.database.SaveClient(client)
		if err != nil {
			log.Println("[ERRO]", err)
			return nil, err
		}
	}

	out := common.NewTLV8Container()
	out.SetByte(TLVSequenceNumber, 0x2)

	return out, nil
}
