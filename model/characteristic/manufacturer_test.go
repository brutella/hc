package characteristic

import (
	"testing"
)

func TestManufacturer(t *testing.T) {
	m := NewManufacturer("Apple")

	if is, want := m.Type, TypeManufacturer; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := m.Manufacturer(), "Apple"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	m.SetManufacturer("Google")

	if is, want := m.Manufacturer(), "Google"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
