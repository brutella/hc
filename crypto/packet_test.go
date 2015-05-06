package crypto

import (
	"bytes"
	"testing"
)

func TestPacketFromBytes(t *testing.T) {
	b := bytes.NewBuffer([]byte{0x01, 0x02, 0x03})
	packets := packetsFromBytes(b)

	if x := len(packets); x != 1 {
		t.Fatal(x)
	}

	p := packets[0]

	if x := p.length; x != 3 {
		t.Fatal(x)
	}
	if x := len(p.value); x != p.length {
		t.Fatal(x)
	}
}

func TestMultiplePacketFromBytes(t *testing.T) {
	b := bytes.NewBuffer([]byte{0x01, 0x02, 0x03, 0x04, 0x05})
	p := packetsWithSizeFromBytes(3, b)

	if x := len(p); x != 2 {
		t.Fatal(x)
	}
	if x := p[0].length; x != 3 {
		t.Fatal(x)
	}
	if x := p[1].length; x != 2 {
		t.Fatal(x)
	}
}
