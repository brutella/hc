package characteristic

import (
	"testing"
)

func TestBool(t *testing.T) {
	b := NewBool(TypeOn)
	b.Value = true

	if is, want := b.GetValue(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b.SetValue(false)

	if is, want := b.GetValue(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestValueUpdate(t *testing.T) {
	b := NewBool(TypeOn)
	b.Perms = PermsWrite()

	var newValue bool
	b.OnValueRemoteUpdate(func(value bool) {
		newValue = value
	})

	b.UpdateValueFromConnection(true, TestConn)

	if is, want := b.GetValue(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	if is, want := newValue, true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
