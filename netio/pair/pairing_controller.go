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

func (c *PairingController) Handle(cont common.Container) (common.Container, error) {
	method := PairMethodType(cont.GetByte(TagPairingMethod))
	username := cont.GetString(TagUsername)
	publicKey := cont.GetBytes(TagPublicKey)

	log.Println("[VERB] ->   Method:", method)
	log.Println("[VERB] -> Username:", username)
	log.Println("[VERB] ->     LTPK:", publicKey)

	client := db.NewClient(username, publicKey)

	switch method {
	case PairingMethodDelete:
		log.Printf("[INFO] Remove LTPK for client '%s'\n", username)
		c.database.DeleteClient(client)
	case PairingMethodAdd:
		err := c.database.SaveClient(client)
		if err != nil {
			log.Println("[ERRO]", err)
			return nil, err
		}
	default:
		return nil, common.NewErrorf("Invalid pairing method type %v", method)
	}

	out := common.NewTLV8Container()
	out.SetByte(TagSequence, 0x2)

	return out, nil
}
