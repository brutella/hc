package crypto

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPacketFromBytes(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03}
	var b bytes.Buffer
	b.Write(data)

	packets := packetsFromBytes(&b)
	assert.Equal(t, len(packets), 1)

	packet := packets[0]
	assert.Equal(t, packet.length, 3)
	assert.Equal(t, len(packet.value), packet.length)
}

func TestMultiplePacketFromBytes(t *testing.T) {
	data := []byte{0x01, 0x02, 0x03, 0x04, 0x05}
	var b bytes.Buffer
	b.Write(data)

	packets := packetsWithSizeFromBytes(3, &b)
	assert.Equal(t, len(packets), 2)

	assert.Equal(t, packets[0].length, 3)
	assert.Equal(t, packets[1].length, 2)
}
