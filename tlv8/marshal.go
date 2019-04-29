package tlv8

func Marshal(v interface{}) ([]byte, error) {
	e := newEncoder()
	e.encode(v)
	return e.wr.bytes(), e.err
}
