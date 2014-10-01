package pair

import (
    "github.com/brutella/hap"
    "github.com/brutella/hap/common"

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
    context     := hap.NewContext(storage)
    controller  := NewPairingController(context)
    
    tlv8_out, err := controller.Handle(tlv8)
    assert.Nil(t, err)
    assert.Nil(t, tlv8_out)
}

func TestDeletePairing(t *testing.T) {
    username := "Unit Test"
    client := hap.NewClient(username, []byte{0x01, 0x02})
    storage, _  := hap.NewFileStorage(os.TempDir())
    context     := hap.NewContext(storage)
    context.SaveClient(client)
    
    tlv8 := common.NewTLV8Container()
    tlv8.SetByte(TLVType_Method, TLVType_Method_PairingDelete)
    tlv8.SetByte(TLVType_SequenceNumber, 0x01)
    tlv8.SetString(TLVType_Username, username)
    
    controller := NewPairingController(context)
    
    tlv8_out, err := controller.Handle(tlv8)
    assert.Nil(t, err)
    assert.Nil(t, tlv8_out)
    
    saved_client := context.ClientForName(username)
    assert.Nil(t, saved_client)
}