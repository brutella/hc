package characteristic

import (
	"testing"
)

func TestOn(t *testing.T) {
	b := NewOn(true)

	if is, want := b.On(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	b.SetOn(false)

	if is, want := b.On(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
