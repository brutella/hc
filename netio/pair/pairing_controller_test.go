package pair

import (
	"github.com/brutella/hc/common"
	"github.com/brutella/hc/db"

	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestAddPairing(t *testing.T) {
	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TLVMethod, MethodAdd)
	tlv8.SetByte(TLVSequenceNumber, 0x01)
	tlv8.SetString(TLVUsername, "Unit Test")
	tlv8.SetBytes(TLVPublicKey, []byte{0x01, 0x02})

	database, _ := db.NewDatabase(os.TempDir())
	controller := NewPairingController(database)

	tlv8_out, err := controller.Handle(tlv8)
	assert.Nil(t, err)
	assert.NotNil(t, tlv8_out)
	assert.Equal(t, tlv8_out.GetByte(TLVSequenceNumber), byte(0x2))
}

func TestDeletePairing(t *testing.T) {
	username := "Unit Test"
	client := db.NewClient(username, []byte{0x01, 0x02})
	database, _ := db.NewDatabase(os.TempDir())
	database.SaveClient(client)

	tlv8 := common.NewTLV8Container()
	tlv8.SetByte(TLVMethod, MethodDelete)
	tlv8.SetByte(TLVSequenceNumber, 0x01)
	tlv8.SetString(TLVUsername, username)

	controller := NewPairingController(database)

	tlv8_out, err := controller.Handle(tlv8)
	assert.Nil(t, err)
	assert.NotNil(t, tlv8_out)
	assert.Equal(t, tlv8_out.GetByte(TLVSequenceNumber), byte(0x2))

	saved_client := database.ClientWithName(username)
	assert.Nil(t, saved_client)
}
