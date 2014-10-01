package pair

import (
    "github.com/brutella/hap"

    "testing"
    "github.com/stretchr/testify/assert"
    "os"
)

func TestAddPairing(t *testing.T) {
    tlv8 := &TLV8Container{}
    tlv8.SetByte(TLVType_Method, TLVType_Method_PairingAdd)
    tlv8.SetByte(TLVType_SequenceNumber, 0x01)
    tlv8.SetString(TLVType_Username, "Unit Test")
    tlv8.SetBytes(TLVType_PublicKey, []byte{0x01, 0x02})
    
    storage, _  := hap.NewFileStorage(os.TempDir())
    context     := hap.NewContext(storage)
    controller  := NewPairingController(context)
    
    err := controller.Handle(tlv8)
    assert.Nil(t, err)
}

func TestDeletePairing(t *testing.T) {
    username := "Unit Test"
    client := hap.NewClient(username, []byte{0x01, 0x02})
    storage, _  := hap.NewFileStorage(os.TempDir())
    context     := hap.NewContext(storage)
    context.SaveClient(client)
    
    tlv8 := &TLV8Container{}
    tlv8.SetByte(TLVType_Method, TLVType_Method_PairingDelete)
    tlv8.SetByte(TLVType_SequenceNumber, 0x01)
    tlv8.SetString(TLVType_Username, username)
    
    controller := NewPairingController(context)
    
    err := controller.Handle(tlv8)
    assert.Nil(t, err)
    
    saved_client := context.ClientForName(username)
    assert.Nil(t, saved_client)
}