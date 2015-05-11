package util

import (
	"bytes"
	"encoding/hex"
	"reflect"
	"testing"
)

func TestTLV8SetByte(t *testing.T) {
	container := NewTLV8Container()
	container.SetByte(1, 0xF)

	if is, want := container.GetByte(1), byte(0xF); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTLV8Bytes(t *testing.T) {
	data := "0102AFFA"
	b, _ := hex.DecodeString(data)
	buf := bytes.NewBuffer(b)
	container, err := NewTLV8ContainerFromReader(buf)

	if err != nil {
		t.Fatal(err)
	}
	if is, want := container.GetBytes(1), []byte{0xAF, 0xFA}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTLV8BytesFromMultipleSource(t *testing.T) {
	data := "0102AFFA0103BFFBAA"
	b, _ := hex.DecodeString(data)

	buf := bytes.NewBuffer(b)
	container, err := NewTLV8ContainerFromReader(buf)

	if err != nil {
		t.Fatal(err)
	}
	if is, want := container.GetBytes(1), []byte{0xAF, 0xFA, 0xBF, 0xFB, 0xAA}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTLV8SetMoreThanMaxBytes(t *testing.T) {
	container := NewTLV8Container()
	data := "00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF" // 384 bytes
	b, _ := hex.DecodeString(data)

	if is, want := len(b), 384; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	container.SetBytes(1, b)

	// split up in 255 chunks
	// 01(type)FF(length=255)bytes...01(type)81(length=129)bytes...
	expectedData := "01FF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEE0181FF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF00112233445566778899AABBCCDDEEFF" // 384 bytes
	expectedBytes, _ := hex.DecodeString(expectedData)
	if is, want := container.BytesBuffer().Bytes(), expectedBytes; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTLV8SetBytes(t *testing.T) {
	container := NewTLV8Container()
	container.SetBytes(1, []byte{0xAF, 0xFA})

	if is, want := container.GetBytes(1), []byte{0xAF, 0xFA}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTLV8BytesBuffer(t *testing.T) {
	container := NewTLV8Container()
	container.SetBytes(1, []byte{0xAF, 0xFA})

	if is, want := container.BytesBuffer().Bytes(), []byte{0x01, 0x02, 0xAF, 0xFA}; reflect.DeepEqual(is, want) == false {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestTLV8String(t *testing.T) {
	container := NewTLV8Container()
	container.SetString(1, "Hello World")

	if is, want := container.GetString(1), "Hello World"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
