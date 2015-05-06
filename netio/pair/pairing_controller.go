package pair

import (
	"fmt"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/util"
	"github.com/brutella/log"
)

// Pairing implements pairing json of format
//     {
//       "guestName": <string>,
//       "guestPublicKey": <string>
//     }
type Pairing struct {
	GuestName      string `json:"guestName"`
	GuestPublicKey string `json:"guestPublicKey"`
}

// PairingController handles un-/pairing with a client by simply exchanging
// the keys going through the pairing process.
type PairingController struct {
	database db.Database
}

// NewPairingController returns a pairing controller.
func NewPairingController(database db.Database) *PairingController {
	c := PairingController{
		database: database,
	}

	return &c
}

// Handle processes a container to pair with a new client without going through the pairing process.
func (c *PairingController) Handle(cont util.Container) (util.Container, error) {
	method := pairMethodType(cont.GetByte(TagPairingMethod))
	username := cont.GetString(TagUsername)
	publicKey := cont.GetBytes(TagPublicKey)

	log.Println("[VERB] ->   Method:", method)
	log.Println("[VERB] -> Username:", username)
	log.Println("[VERB] ->     LTPK:", publicKey)

	entity := db.NewEntity(username, publicKey, nil)

	switch method {
	case PairingMethodDelete:
		log.Printf("[INFO] Remove LTPK for client '%s'\n", username)
		c.database.DeleteEntity(entity)
	case PairingMethodAdd:
		err := c.database.SaveEntity(entity)
		if err != nil {
			log.Println("[ERRO]", err)
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Invalid pairing method type %v", method)
	}

	out := util.NewTLV8Container()
	out.SetByte(TagSequence, 0x2)

	return out, nil
}
