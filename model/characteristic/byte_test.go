package characteristic

import (
	"reflect"
	"testing"
)

func TestByteCharacteristic(t *testing.T) {
	b := NewByteCharacteristic(0xFA, PermsAll())

	if x := b.Byte(); reflect.DeepEqual(x, byte(0xFA)) == false {
		t.Fatal(x)
	}

	b.SetByte(0xAF)

	if x := b.Byte(); reflect.DeepEqual(x, byte(0xAF)) == false {
		t.Fatal(x)
	}
}
