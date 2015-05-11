package util

import (
	"reflect"
	"testing"
)

func TestSerialForName(t *testing.T) {
	storage, err := NewTempFileStorage()
	if err != nil {
		t.Fatal(err)
	}

	name := "My Accessory"
	serial := GetSerialNumberForAccessoryName(name, storage)
	same := GetSerialNumberForAccessoryName(name, storage)

	if is, want := serial, same; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
