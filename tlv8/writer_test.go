package tlv8

import (
	"bytes"
	"reflect"
	"testing"
)

func TestWriteBytes(t *testing.T) {
	wr := newWriter()

	buf := make([]byte, 256)
	wr.writeBytes(1, buf)

	expected := make([]byte, 260)
	expected[0] = 0x1
	expected[1] = 0xFF
	expected[257] = 0x1
	expected[258] = 0x1

	if is, want := wr.bytes(), expected; !reflect.DeepEqual(is, want) {
		t.Fatalf("%v != %v", is, want)
	}

	rd, err := newReader(bytes.NewBuffer(wr.bytes()))
	if err != nil {
		t.Fatal(err)
	}

	read, err := rd.readBytes(0x1)
	if err != nil {
		t.Fatal(err)
	}
	if is, want := read, buf; !reflect.DeepEqual(is, want) {
		t.Fatalf("%v len(%d) != %v len(%d)", is, len(is), want, len(want))
	}
}
