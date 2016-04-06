package characteristic

import (
	"encoding/base64"
	"reflect"
	"testing"
)

func TestBytesEncoding(t *testing.T) {
	val := []byte{0xFA, 0xAA}
	b := NewBytes(TypeLogs)
	b.SetValue(val)

	tlv8 := []byte{0x00, 0x02, 0xFA, 0xAA}
	expect := base64.StdEncoding.EncodeToString(tlv8)

	if x := b.Value; reflect.DeepEqual(x, expect) == false {
		t.Fatal(x)
	}

	if x := b.GetValue(); reflect.DeepEqual(x, val) == false {
		t.Fatal(x)
	}
}
