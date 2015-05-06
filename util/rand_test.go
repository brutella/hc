package util

import (
	"testing"
)

func TestRandomHexString(t *testing.T) {
	s := RandomHexString()

	if x := len(s); x != 32 {
		t.Fatal(x)
	}

	for _, c := range s {
		if c >= 0x61 && c <= 0x66 { // a to f
			continue
		}
		if c >= 0x30 && c <= 0x39 { // 0 to 9
			continue
		}
		t.Fatalf("illegal hex character '%c'", c)
	}
}
