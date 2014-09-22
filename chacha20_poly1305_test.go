package gohap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "encoding/binary"
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