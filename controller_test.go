package gohap

import (
	"testing"
    "github.com/stretchr/testify/assert"
)

func TestPairingStart(t *testing.T) {
    tlv := TLV8Container{}
    tlv.Set(TLVType_AuthMethod, []byte{0})
    tlv.Set(TLVType_SequenceNumber, []byte{SequenceStartRequest})
    
    controller, err := NewPairingController("Pair-Setup", "123-45-678")
    assert.Nil(t, err)
    
    reader, err := controller.Handle(tlv.BytesBuffer())
    assert.Nil(t, err)
    
    result, err := ReadTLV8(reader)
    assert.Nil(t, err)
    
    seq := result.GetUInt64(TLVType_SequenceNumber)
    assert.Equal(t, seq, uint64(SequenceStartRespond))
    salt := result.GetBytes(TLVType_Salt)
    assert.Equal(t, len(salt), 16) // must be 16 bytes long
    key := result.GetBytes(TLVType_PublicKey)
    assert.NotNil(t, key)
}    