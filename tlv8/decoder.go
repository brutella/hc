package tlv8

import (
	"github.com/xiam/to"

	"bytes"
	"io"
	"reflect"
	"strings"
)

type decoder struct {
	r *reader
}

func newDecoder(b []byte) (*decoder, error) {
	r, err := newReader(bytes.NewBuffer(b))
	return &decoder{r}, err
}

func (d *decoder) decodeSlice(v interface{}) error {
	vValue := reflect.ValueOf(v)

	eValue := vValue.Elem()
	e := eValue.Interface()
	eType := reflect.TypeOf(e)

	if eType.Kind() != reflect.Slice {
		return &UnexpectedTypeError{reflect.TypeOf(v)}
	}

	for i := 0; i < eValue.Len(); i++ {
		eValue := eValue.Index(i)
		e := interfaceOf(eValue)
		if err := d.decode(e); err != nil {
			return err
		}
	}

	return nil
}

func (d *decoder) decode(v interface{}) error {

	vValue := reflect.ValueOf(v)

	if vValue.Kind() != reflect.Ptr || vValue.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	eValue := vValue.Elem()
	e := eValue.Interface()
	eType := reflect.TypeOf(e)

	if eType.Kind() == reflect.Slice {
		return d.decodeSlice(v)
	}

	for i := 0; i < eValue.NumField(); i++ {
		if tlv8, ok := eType.Field(i).Tag.Lookup("tlv8"); ok {
			values := strings.Split(tlv8, ",")
			tag := uint8(to.Uint64(values[0]))

			field := eValue.Field(i)
			switch value := field.Interface().(type) {
			case uint8:
				if v, err := d.r.readByte(tag); err == nil {
					field.SetUint(uint64(v))
				} else if err == io.EOF {
					continue
				} else {
					return err
				}
			case uint16:
				if v, err := d.r.readUint16(tag); err == nil {
					field.SetUint(uint64(v))
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case int16:
				if v, err := d.r.readint16(tag); err == nil {
					field.SetInt(int64(v))
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case uint32:
				if v, err := d.r.readUint32(tag); err == nil {
					field.SetUint(uint64(v))
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case int32:
				if v, err := d.r.readint32(tag); err == nil {
					field.SetInt(int64(v))
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case int64:
				if v, err := d.r.readint64(tag); err == nil {
					field.SetInt(v)
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case uint64:
				if v, err := d.r.readUint64(tag); err == nil {
					field.SetUint(v)
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case float32:
				if v, err := d.r.readFloat32(tag); err == nil {
					field.SetFloat(float64(v))
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case []byte:
				if v, err := d.r.readBytes(tag); err == nil {
					field.SetBytes(v)
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			case string:
				if v, err := d.r.readString(tag); err == nil {
					field.SetString(v)
				} else if err == io.EOF {
					continue
				} else {
					return err
				}
			case bool:
				if v, err := d.r.readBool(tag); err == nil {
					field.SetBool(v)
				} else if err == io.EOF {
					continue
				} else {
					return err
				}

			default:
				valueType := reflect.TypeOf(value)
				// elemValue is the Value of a new instance
				var elemValue reflect.Value
				if valueType.Kind() == reflect.Slice {
					var slice = reflect.MakeSlice(valueType, 0, 0)
					for {
						instanceValue := newValueOf(valueType)
						v := instanceValue.Interface()
						var err error
						if tlv8 == "-" {
							// unnamed slices are inline encoded
							err = d.decode(v)
							if isEmptyStruct(v) {
								// step out of loop
								break
							}
						} else {
							b, e := d.r.readBytes(tag)
							if e == io.EOF {
								break
							}

							err = e
							if err == nil {
								structDecoder, e := newDecoder(b)
								if e != nil {
									err = e
									break
								}

								err = structDecoder.decode(v)
							}
						}

						if err == nil || err == io.EOF {
							slice = reflect.Append(slice, instanceValue.Elem())
						}

						if d.r.eof() {
							// step out of loop
							break
						}

						// stop decoding
						if err != nil {
							return err
						}
					}

					elemValue = slice
				} else {
					elemValue = newValueOf(valueType)
					data, err := d.r.readBytes(tag)
					if err == nil {
						err = unmarshal(data, elemValue.Interface())
					}

					if err == io.EOF {
						break
					}

					if err != nil {
						return err
					}
				}

				if field.Kind() == reflect.Ptr || field.Kind() == reflect.Slice {
					field.Set(elemValue)
				} else {
					field.Set(elemValue.Elem())
				}
			}
		}
	}

	return nil
}

func newValueOf(t reflect.Type) reflect.Value {
	if t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice {
		return reflect.New(t.Elem())
	}

	return reflect.New(t)
}

type InvalidUnmarshalError struct {
	Type reflect.Type
}

func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "tlv8: Unmarshal(nil)"
	}
	if e.Type.Kind() != reflect.Ptr {
		return "tlv8: Unmarshal(non-pointer " + e.Type.String() + ")"
	}

	return "tlv8: Unmarshal(nil " + e.Type.String() + ")"
}

func isEmptyStruct(v interface{}) bool {
	vValue := reflect.ValueOf(v)
	if vValue.Kind() == reflect.Ptr {
		vValue = vValue.Elem()
	}

	for i := 0; i < vValue.NumField(); i++ {
		field := vValue.Field(i)
		if !isEmptyValue(field) {
			return false
		}
	}

	return true
}

func isEmptyValue(v reflect.Value) bool {
	switch v.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return v.Len() == 0
	case reflect.Bool:
		return !v.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return v.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return v.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return v.Float() == 0
	case reflect.Interface, reflect.Ptr:
		return v.IsNil()
	}
	return false
}
