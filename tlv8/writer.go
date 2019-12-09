package tlv8

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
)

type writer struct {
	buf bytes.Buffer
}

func newWriter() *writer {
	var buf bytes.Buffer
	return &writer{buf}
}

func (wr *writer) bytes() []byte {
	return wr.buf.Bytes()
}

func (wr *writer) writeBytes(tag uint8, value []byte) {
	buff := bytes.NewBuffer(value)

	for {
		var bytes = make([]byte, 255)
		n, err := io.ReadFull(buff, bytes)
		if err == nil || err == io.ErrUnexpectedEOF {
			v := bytes[:n]

			// Write tag, length, value
			b := append([]byte{tag, uint8(n)}, v[:]...)
			wr.write(b)

			if err == io.ErrUnexpectedEOF { // Fewer than 255 bytes read
				break
			}
		} else {
			break
		}
	}
}

func (wr *writer) writeUint16(tag uint8, v uint16) {
	var b [2]byte
	binary.LittleEndian.PutUint16(b[:], v)
	wr.writeBytes(tag, b[:])
}

func (wr *writer) writeUint32(tag uint8, v uint32) {
	var b [4]byte
	binary.LittleEndian.PutUint32(b[:], v)
	wr.writeBytes(tag, b[:])
}

func (wr *writer) writeInt16(tag uint8, v int16) {
	buf := make([]byte, 2)
	// little endian
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	wr.writeBytes(tag, buf[:2])
}

func (wr *writer) writeInt32(tag uint8, v int32) {
	buf := make([]byte, 4)
	// little endian
	buf[0] = byte(v)
	buf[1] = byte(v >> 8)
	buf[2] = byte(v >> 16)
	buf[3] = byte(v >> 24)
	wr.writeBytes(tag, buf[:4])
}

func (wr *writer) writeFloat32(tag uint8, v float32) {
	var b [4]byte
	math.Float32bits(v)
	wr.writeBytes(tag, b[:])
}

func (wr *writer) writeBool(tag uint8, b bool) {
	if b == true {
		wr.write([]byte{tag, 1, 1})
	} else {
		wr.write([]byte{tag, 1, 0})
	}
}

func (wr *writer) writeString(tag uint8, s string) {
	wr.writeBytes(tag, []byte(s))
}

func (wr *writer) writeByte(tag uint8, b byte) {
	wr.write([]byte{tag, 1, b})
}

func (wr *writer) write(b []byte) (int, error) {
	return wr.buf.Write(b)
}
