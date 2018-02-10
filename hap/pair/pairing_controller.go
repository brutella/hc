package pair

import (
	"fmt"
	"github.com/brutella/hc/db"
	"github.com/brutella/hc/log"
	"github.com/brutella/hc/util"
)

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
	method := PairMethodType(cont.GetByte(TagPairingMethod))
	perm := cont.GetByte(TagPermission)
	username := cont.GetString(TagUsername)
	publicKey := cont.GetBytes(TagPublicKey)

	log.Debug.Println("->     Method:", method)
	log.Debug.Println("-> Permission:", perm)
	log.Debug.Println("->   Username:", username)
	log.Debug.Println("->       LTPK:", publicKey)

	entity := db.NewEntity(username, publicKey, nil)

	out := util.NewTLV8Container()
	out.SetByte(TagSequence, 0x2)

	switch method {
	case PairingMethodDelete:
		log.Debug.Printf("Remove LTPK for client '%s'\n", username)
		c.database.DeleteEntity(entity)
	case PairingMethodAdd:
		if perm != AdminPerm {
			log.Info.Println("Non-admin controllers are not allowed to add pairings")
			out.SetByte(TagErrCode, ErrCodeAuthenticationFailed.Byte())
			break
		}

		err := c.database.SaveEntity(entity)
		if err != nil {
			log.Info.Panic(err)
			return nil, err
		}
	default:
		return nil, fmt.Errorf("Invalid pairing method type %v", method)
	}

	return out, nil
}
