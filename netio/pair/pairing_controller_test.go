package pair

import (
    "github.com/brutella/hap"
    "github.com/brutella/hap/common"
    "github.com/brutella/hap/db"

    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestAddPairing(t *testing.T) {
    tlv8 := common.NewTLV8Container()
    tlv8.SetByte(TLVType_Method, TLVType_Method_PairingAdd)
    tlv8.SetByte(TLVType_SequenceNumber, 0x01)
    tlv8.SetString(TLVType_Username, "Unit Test")
    tlv8.SetBytes(TLVType_PublicKey, []byte{0x01, 0x02})
    
    storage, _  := hap.NewFileStorage(os.TempDir())
    database    := db.NewManager(storage)
    controller  := NewPairingController(database)
    
    tlv8_out, err := controller.Handle(tlv8)
    assert.Nil(t, err)
    assert.Nil(t, tlv8_out)
}

func TestDeletePairing(t *testing.T) {
    username := "Unit Test"
    client := db.NewClient(username, []byte{0x01, 0x02})
    storage, _  := hap.NewFileStorage(os.TempDir())
    database    := db.NewManager(storage)
    database.SaveClient(client)
    
    tlv8 := common.NewTLV8Container()
    tlv8.SetByte(TLVType_Method, TLVType_Method_PairingDelete)
    tlv8.SetByte(TLVType_SequenceNumber, 0x01)
    tlv8.SetString(TLVType_Username, username)
    
    controller := NewPairingController(database)
    
    tlv8_out, err := controller.Handle(tlv8)
    assert.Nil(t, err)
    assert.Nil(t, tlv8_out)
    
    saved_client := database.ClientForName(username)
    assert.Nil(t, saved_client)
}