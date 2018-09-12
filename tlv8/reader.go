package tlv8

import (
	"bytes"
	"encoding/binary"
	"io"
)

type bucket []byte
type hash map[byte][]bucket

type reader struct {
	m hash
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

func (r *reader) readBytes(tag byte) ([]byte, error) {
	list := r.m[tag]
	for _, b := range list {
		list = append(list[:0], list[1:]...)
		if len(list) > 0 {
			r.m[tag] = list
		} else {
			delete(r.m, tag)
		}

		return b, nil
	}

	return nil, io.EOF
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
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint16(b), nil
	}
}

func (r *reader) readUint32(tag byte) (uint32, error) {
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		return binary.LittleEndian.Uint32(b), nil
	}
}

func (r *reader) readint16(tag byte) (int16, error) {
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		var v int16
		// little endian
		v |= int16(b[0])
		v |= int16(b[1]) << 8

		return v, nil
	}
}

func (r *reader) readint32(tag byte) (int32, error) {
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		var v int32
		// little endian
		v |= int32(b[0])
		v |= int32(b[1]) << 8
		v |= int32(b[2]) << 16
		v |= int32(b[3]) << 24

		return v, nil
	}
}

func (r *reader) readFloat32(tag byte) (float32, error) {
	if b, err := r.readBytes(tag); err != nil {
		return 0, err
	} else {
		var buf = bytes.NewReader(b)
		var result float32
		binary.Read(buf, binary.LittleEndian, result)
		return result, nil
	}
}

func read(r io.Reader) (hash, error) {
	var h = hash{}
	for r != nil {
		var tag byte
		var n byte
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

		var list []bucket
		if l, ok := h[tag]; ok {
			list = l
		} else {
			list = []bucket{}
		}
		list = append(list, v)
		h[tag] = list
	}

	return h, nil
}
