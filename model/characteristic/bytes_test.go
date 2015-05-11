package characteristic

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestSetBytes(t *testing.T) {
	b := NewBytes(nil)
	b.SetBytes([]byte{0xAB, 0xBA})

	if x := b.Bytes(); reflect.DeepEqual(x, []byte{0xAB, 0xBA}) == false {
		t.Fatal(x)
	}
}

func TestInitBytes(t *testing.T) {
	b := NewBytes([]byte{0xFA, 0xAA})

	if x := b.Bytes(); reflect.DeepEqual(x, []byte{0xFA, 0xAA}) == false {
		t.Fatal(x)
	}
}

func TestBytesEncoding(t *testing.T) {
	b := NewBytes([]byte{0xFA, 0xAA})
	tlv8 := []byte{0x00, 0x02, 0xFA, 0xAA}
	expect := base64.StdEncoding.EncodeToString(tlv8)

	if x := b.GetValue(); reflect.DeepEqual(x, expect) == false {
		t.Fatal(x)
	}
}
