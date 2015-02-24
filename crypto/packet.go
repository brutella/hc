package crypto

import (
	"io"
)

const (
	// Defines the max length of encrypted packets
	PacketLengthMax = 0x400
)

type packet struct {
	length int
	value  []byte
}

// PacketsWithSizeFromBytes returns lv (tlv without t(ype)) packets
func PacketsWithSizeFromBytes(length int, r io.Reader) []packet {
	packets := make([]packet, 0)

	for {
		var value = make([]byte, length)
		n, err := r.Read(value)
		if n == 0 {
			break
		}

		if n > length {
			panic("Invalid length")
		}

		p := packet{length: n, value: value[:n]}
		packets = append(packets, p)

		if n < length || err == io.EOF {
			break
		}
	}

	return packets
}

// PacketsFromBytes returns packets with length PacketLengthMax
func PacketsFromBytes(r io.Reader) []packet {
	return PacketsWithSizeFromBytes(PacketLengthMax, r)
}
