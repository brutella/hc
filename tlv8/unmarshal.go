package tlv8

import ()

func Unmarshal(data []byte, v interface{}) error {
	return unmarshal(data, v)
}

func unmarshal(data []byte, v interface{}) error {
	d, err := newDecoder(data)
	if err != nil {
		return err
	}

	d.decode(v)

	return d.err
}
