package crypto

import (
	"io"
)

const (
	// PacketLengthMax is the max length of encrypted packets
	PacketLengthMax = 0x400
)

type packet struct {
	length int
	value  []byte
}

// packetsWithSizeFromBytes returns lv (tlv without t(ype)) packets
func packetsWithSizeFromBytes(length int, r io.Reader) []packet {
	var packets []packet
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

// packetsFromBytes returns packets with length PacketLengthMax
func packetsFromBytes(r io.Reader) []packet {
	return packetsWithSizeFromBytes(PacketLengthMax, r)
}
