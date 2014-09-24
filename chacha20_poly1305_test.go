package hap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/binary"
    "encoding/hex"
    "github.com/codahale/chacha20"
)

func TestAddBytesMod8(t *testing.T) {
    b := []byte{}
    add := []byte{0xFF}
    
    assert.Equal(t, AddBytes(b, add, 8), []byte{0xFF, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
}


func TestAddBytesMod8FromUint64(t *testing.T) {
    b := []byte{}
    length := make([]byte, 8)
    binary.LittleEndian.PutUint64(length, uint64(1)) // [0x1 0x0 0x0 0x0 0x0 0x0 0x0 0x0 0x0]
    
    assert.Equal(t, AddBytes(b, length, 8), []byte{0x1, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0})
}

func TestChacha20(t *testing.T) {
    K, _ := hex.DecodeString("6a3bfd77d9efac53f8ef51712796bf7a37541f425a5dc5397c8a2c3c040d9301")
    message, _ := hex.DecodeString("8e685bd3237866e7a424b0f33df1a087a397a78e147042d2d17b159044d2ad1162dea13df2a119b61c90d62fc76335f49954557f2b07c463dca1664ca042599fca66068b16bc3e7e1896536ca2")
    
    c, err := chacha20.NewCipher(K, []byte("PS-Msg05"))
    assert.Nil(t, err)
    var out = make([]byte, len(message))
    c.XORKeyStream(out, message)
    
    c2, err := chacha20.NewCipher(K, []byte("PS-Msg05"))
    
    assert.Nil(t, err)
    var out2 = make([]byte, len(message))
    c2.XORKeyStream(out2, out)
    assert.Equal(t, out2, message)
}