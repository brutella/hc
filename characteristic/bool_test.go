package characteristic

import (
	"testing"
)

func TestBool(t *testing.T) {
	b := NewOn()
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
	b := NewOn()

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
