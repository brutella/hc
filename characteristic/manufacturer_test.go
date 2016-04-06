package characteristic

import (
	"testing"
)

func TestManufacturer(t *testing.T) {
	m := NewManufacturer()
	m.Value = "Apple"

	if is, want := m.Type, TypeManufacturer; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := m.GetValue(), "Apple"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	m.SetValue("Google")

	if is, want := m.GetValue(), "Google"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
