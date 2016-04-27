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

	expect := base64.StdEncoding.EncodeToString(val)

	if x := b.Value; reflect.DeepEqual(x, expect) == false {
		t.Fatal(x)
	}

	if x := b.GetValue(); reflect.DeepEqual(x, val) == false {
		t.Fatal(x)
	}
}
