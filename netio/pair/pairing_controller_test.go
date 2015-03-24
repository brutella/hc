package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestUnknownPairingMethod(t *testing.T) {
	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TagPairingMethod, 0x09)

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	out, err := controller.Handle(tlv8)
	assert.NotNil(t, err)
	assert.Nil(t, out)
}

func TestAddPairing(t *testing.T) {
	in := common.NewTLV8Container()
	in.SetByte(TagPairingMethod, PairingMethodAdd.Byte())
	in.SetByte(TagSequence, 0x01)
	in.SetString(TagUsername, "Unit Test")
	in.SetBytes(TagPublicKey, []byte{0x01, 0x02})

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	out, err := controller.Handle(in)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, out.GetByte(TagSequence), byte(0x2))
}

func TestDeletePairing(t *testing.T) {
	username := "Unit Test"
	entity := db.NewEntity(username, []byte{0x01, 0x02}, nil)
	database, _ := db.NewDatabase(os.TempDir())
	database.SaveEntity(entity)

	in := common.NewTLV8Container()
	in.SetByte(TagPairingMethod, PairingMethodDelete.Byte())
	in.SetByte(TagSequence, 0x01)
	in.SetString(TagUsername, username)

	controller := NewPairingController(database)

	out, err := controller.Handle(in)
	assert.Nil(t, err)
	assert.NotNil(t, out)
	assert.Equal(t, out.GetByte(TagSequence), byte(0x2))

	savedEntity := database.EntityWithName(username)
	assert.Nil(t, savedEntity)
}
