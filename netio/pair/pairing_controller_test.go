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

	tlv8_out, err := controller.Handle(tlv8)
	assert.NotNil(t, err)
	assert.Nil(t, tlv8_out)
}

func TestAddPairing(t *testing.T) {
	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TagPairingMethod, PairingMethodAdd)
	tlv8.SetByte(TagSequence, 0x01)
	tlv8.SetString(TagUsername, "Unit Test")
	tlv8.SetBytes(TagPublicKey, []byte{0x01, 0x02})

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	tlv8_out, err := controller.Handle(tlv8)
	assert.Nil(t, err)
	assert.NotNil(t, tlv8_out)
	assert.Equal(t, tlv8_out.GetByte(TagSequence), byte(0x2))
}

func TestDeletePairing(t *testing.T) {
	username := "Unit Test"
	client := db.NewClient(username, []byte{0x01, 0x02})
	database, _ := db.NewDatabase(os.TempDir())
	database.SaveClient(client)

	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TagPairingMethod, PairingMethodDelete)
	tlv8.SetByte(TagSequence, 0x01)
	tlv8.SetString(TagUsername, username)

	controller := NewPairingController(database)

	tlv8_out, err := controller.Handle(tlv8)
	assert.Nil(t, err)
	assert.NotNil(t, tlv8_out)
	assert.Equal(t, tlv8_out.GetByte(TagSequence), byte(0x2))

	saved_client := database.ClientWithName(username)
	assert.Nil(t, saved_client)
}
