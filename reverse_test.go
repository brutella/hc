package hap

import (
	"testing"
    "github.com/stretchr/testify/assert"
    "sort"
)

func TestReverseByteArray(t *testing.T) {
    bytes := ByteSequence([]byte{0x01, 0x02})
    reverse := ByteSequence([]byte{0x02, 0x01})
    sort.Sort(sort.Reverse(bytes))
    assert.Equal(t, bytes, reverse)
}
