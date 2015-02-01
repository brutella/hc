package crypto

import (
	"io"
)

const (
	PacketLengthMax = 0x400
)

type packet struct {
	length int
	value  []byte
}

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

// Returns packets with the default HAP length of 1024
func PacketsFromBytes(r io.Reader) []packet {
	return PacketsWithSizeFromBytes(1024, r)
}
