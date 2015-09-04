package characteristic

import (
	"bytes"
	"io/ioutil"
	"os"
)

// fileMarshaler returns the file content as base64 string
type fileMarshaler struct {
	filePath string
}

// MarshalJSON is the required method in the json.Marshaler interface
func (m *fileMarshaler) MarshalJSON() ([]byte, error) {
	f, err := os.Open(m.filePath)
	if err != nil {
		return nil, err
	}

	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	b64 := tlv8Base64FromBytes(b)
	var buf bytes.Buffer
	buf.WriteString("\"")
	buf.Write([]byte(b64))
	buf.WriteString("\"")
	return buf.Bytes(), nil
}

type Log struct {
	*Bytes
}

func NewLog(filePath string) *Log {
	m := &fileMarshaler{filePath}
	c := NewBytes([]byte{})
	c.Type = TypeLogs
	c.Permissions = PermsRead()
	// Value is represented by an object which returns the file content as base64 string
	c.Value = m

	return &Log{c}
}
