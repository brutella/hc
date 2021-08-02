package tlv8

import (
	"encoding/binary"
	"io"
	"math"
)

type bucket []byte

type reader struct {
	m map[byte][]bucket
}

func newReader(r io.Reader) (*reader, error) {
	m, err := read(r)

	return &reader{m}, err
}

func (r *reader) readByte(tag byte) (byte, error) {
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		return b[0], nil
	}
}

func (r *reader) eof() bool {
	return len(r.m) == 0
}

func (r *reader) len(tag byte) int {
	if list := r.m[tag]; list != nil && len(list) > 0 {
		return len(list[0])
	}

	return 0
}

func (r *reader) readBytes(tag byte) ([]byte, error) {
	list := r.m[tag]

	if len(list) == 0 {
		return nil, io.EOF
	}

	b := list[0]
	if len(list) > 1 {
		r.m[tag] = append(list[:0], list[1:]...)
	} else {
		delete(r.m, tag)
	}

	return b, nil
}

func (r *reader) readBool(tag byte) (bool, error) {
	if b, err := r.readByte(tag); err != nil {
		return false, err
	} else {
		return (b == 1), nil
	}

}

func (r *reader) readString(tag byte) (string, error) {
	if b, err := r.readBytes(tag); err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

func (r *reader) readUint16(tag byte) (uint16, error) {
	if r.len(tag) < 2 {
		v, err := r.readByte(tag)
		return uint16(v), err
	}

	b, err := r.readBytes(tag)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint16(b), nil
}

func (r *reader) readUint32(tag byte) (uint32, error) {
	if r.len(tag) < 4 {
		v, err := r.readUint16(tag)
		return uint32(v), err
	}

	b, err := r.readBytes(tag)
	if err != nil {
		return 0, err
	}

	return binary.LittleEndian.Uint32(b), nil
}

func (r *reader) readUint64(tag byte) (uint64, error) {
	if r.len(tag) < 8 {
		v, err := r.readUint32(tag)
		return uint64(v), err
	}

	b, err := r.readBytes(tag)
	if err != nil {
		return 0, err
	}

	if len(b) < 8 {
		v, err := r.readUint32(tag)
		return uint64(v), err
	}

	return binary.LittleEndian.Uint64(b), nil
}

func (r *reader) readint16(tag byte) (int16, error) {
	if r.len(tag) < 2 {
		v, err := r.readByte(tag)
		return int16(v), err
	}

	b, err := r.readBytes(tag)
	if err != nil {
		return 0, err
	}

	var v int16
	// little endian
	v |= int16(b[0])
	v |= int16(b[1]) << 8

	return v, nil
}

func (r *reader) readint32(tag byte) (int32, error) {
	if r.len(tag) < 4 {
		v, err := r.readint16(tag)
		return int32(v), err
	}

	b, err := r.readBytes(tag)
	if err != nil {
		return 0, err
	}

	var v int32
	// little endian
	v |= int32(b[0])
	v |= int32(b[1]) << 8
	v |= int32(b[2]) << 16
	v |= int32(b[3]) << 24

	return v, nil
}

func (r *reader) readint64(tag byte) (int64, error) {
	if r.len(tag) < 8 {
		v, err := r.readint32(tag)
		return int64(v), err
	}

	b, err := r.readBytes(tag)

	if err != nil {
		return 0, err
	}

	var v int64
	// little endian
	v |= int64(b[0])
	v |= int64(b[1]) << 8
	v |= int64(b[2]) << 16
	v |= int64(b[3]) << 24
	v |= int64(b[4]) << 32
	v |= int64(b[5]) << 40
	v |= int64(b[6]) << 48
	v |= int64(b[7]) << 56

	return v, nil
}

func (r *reader) readFloat32(tag byte) (float32, error) {
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		bits := binary.LittleEndian.Uint32(b)
		return math.Float32frombits(bits), nil
	}
}

func read(r io.Reader) (map[byte][]bucket, error) {
	var h = map[byte][]bucket{}

	var tag, n byte
	var lastItemWasDelimiter bool
	for {
		if err := binary.Read(r, binary.LittleEndian, &tag); err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		if err := binary.Read(r, binary.LittleEndian, &n); err != nil {
			return nil, err
		}

		var v = make([]byte, n)
		if err := binary.Read(r, binary.LittleEndian, &v); err != nil {
			return nil, err
		}

		if len(v) > 0 {
			if l, ok := h[tag]; ok {
				if lastItemWasDelimiter {
					h[tag] = append(l, v)
				} else {
					h[tag] = []bucket{append(l[0], v...)}
				}
			} else {
				h[tag] = []bucket{v}
			}
		}

		lastItemWasDelimiter = tag == 0 && n == 0
	}

	return h, nil
}
