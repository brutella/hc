package tlv8

import (
	"bytes"
	"github.com/xiam/to"
	"reflect"
)

type encoder struct {
	wr  *writer
	err error
}

func newEncoder() *encoder {
	return &encoder{newWriter(), nil}
}

func (coder *encoder) encodeSlice(v interface{}) {
	if b, err := slicePayload(v); err != nil {
		coder.err = err
	} else {
		coder.write(b)
	}
}

func (coder *encoder) encode(v interface{}) {
	vValue := reflect.ValueOf(v)

	if vValue.Kind() == reflect.Slice {
		coder.encodeSlice(v)
		return
	}

	e := interfaceOf(vValue)
	if b, err := structPayload(e); err != nil {
		coder.err = err
	} else {
		coder.write(b)
	}
}

func (coder *encoder) write(b []byte) {
	if _, err := coder.wr.write(b); err != nil {
		coder.err = err
	}
}

func structPayload(v interface{}) ([]byte, error) {
	wr := newWriter()

	vValue := reflect.ValueOf(v)
	vType := reflect.TypeOf(v)

	if vValue.Kind() == reflect.Ptr {
		return nil, &UnexpectedTypeError{vType}
	}

	for i := 0; i < vType.NumField(); i++ {
		if tlv8, ok := vType.Field(i).Tag.Lookup("tlv8"); ok {
			tag := uint8(to.Uint64(tlv8))
			field := vValue.Field(i)
			switch v := field.Interface().(type) {
			case uint8:
				wr.writeByte(tag, v)
			case []byte:
				wr.writeBytes(tag, v)
			case string:
				wr.writeString(tag, v)
			case uint16:
				wr.writeUint16(tag, v)
			case uint32:
				wr.writeUint32(tag, v)
			case int16:
				wr.writeInt16(tag, v)
			case int32:
				wr.writeInt32(tag, v)
			case float32:
				wr.writeFloat32(tag, v)
			case int64:
				wr.writeInt64(tag, v)
			case uint64:
				wr.writeUint64(tag, v)
			case bool:
				wr.writeBool(tag, v)
			default:
				vValue := reflect.ValueOf(v)
				if vValue.Kind() == reflect.Slice {
					for i := 0; i < vValue.Len(); i++ {
						eValue := vValue.Index(i)
						b, err := structPayload(interfaceOf(eValue))
						if err != nil {
							return nil, err
						}

						if i > 0 {
							// delimit elements with {0x00, 0x00}
							wr.write([]byte{0x0, 0x0})
						}

						if tlv8 == "-" {
							wr.write(b)
						} else {
							// every element in a named slice is encoded by the slice field tlv8 tag
							wr.writeBytes(tag, b)
						}
					}
				} else {
					e := interfaceOf(vValue)
					if b, err := structPayload(e); err != nil {
						return nil, err
					} else {
						wr.writeBytes(tag, b)
					}
				}
			}
		}
	}

	return wr.bytes(), nil
}

func slicePayload(v interface{}) ([]byte, error) {
	vValue := reflect.ValueOf(v)

	if vValue.Kind() != reflect.Slice {
		return nil, &UnexpectedTypeError{reflect.TypeOf(v)}
	}

	var buf bytes.Buffer
	for i := 0; i < vValue.Len(); i++ {
		eValue := vValue.Index(i)

		if b, err := structPayload(interfaceOf(eValue)); err != nil {
			return nil, err
		} else {
			buf.Write(b)
		}
	}

	return buf.Bytes(), nil
}

// interfaceOf returns the value of v as interface{}.
func interfaceOf(v reflect.Value) interface{} {
	if v.Kind() == reflect.Ptr {
		return v.Elem().Interface()
	}

	return v.Interface()
}

type UnexpectedTypeError struct {
	Type reflect.Type
}

func (e *UnexpectedTypeError) Error() string {
	return "tlv8: " + e.Type.String() + " unexpected type"
}
