package characteristic

import (
	"testing"
)

func TestOn(t *testing.T) {
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
