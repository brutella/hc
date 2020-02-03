package characteristic

import (
	"testing"
)

func TestNumberIntOutOfBounds(t *testing.T) {
	number := NewBrightness()
	number.SetValue(120)
	if is, want := number.GetValue(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	number.SetValue(-40)
	if is, want := number.GetValue(), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
