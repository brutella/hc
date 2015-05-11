package characteristic

import (
	"testing"
)

func TestBool(t *testing.T) {
	b := NewBool(true, PermsAll())

	if is, want := b.BoolValue(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b.SetBool(false)

	if is, want := b.BoolValue(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
