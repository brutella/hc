package characteristic

import (
	"encoding/base64"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSetBytes(t *testing.T) {
	b := NewBytes(nil)
	b.SetBytes([]byte{0xAB, 0xBA})
	assert.Equal(t, b.Bytes(), []byte{0xAB, 0xBA})
}

func TestInitBytes(t *testing.T) {
	b := NewBytes([]byte{0xFA, 0xAA})
	assert.Equal(t, b.Bytes(), []byte{0xFA, 0xAA})
}

func TestBytesEncoding(t *testing.T) {
	b := NewBytes([]byte{0xFA, 0xAA})
	tlv8 := []byte{0x00, 0x02, 0xFA, 0xAA}
	assert.Equal(t, b.GetValue(), base64.StdEncoding.EncodeToString(tlv8))
}
