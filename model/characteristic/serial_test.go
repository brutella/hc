package characteristic

import (
	"testing"
)

func TestSerialCharacteristic(t *testing.T) {
	str := NewSerialNumber("001002")

	if is, want := str.Type, TypeSerialNumber; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := str.SerialNumber(), "001002"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	str.SetSerialNumber("001003")

	if is, want := str.SerialNumber(), "001003"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
