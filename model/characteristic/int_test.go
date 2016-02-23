package characteristic

import (
	"testing"
)

func TestNumberIntOutOfBounds(t *testing.T) {
	number := NewInt(20, 0, 100, 1, PermsAll())

	number.SetInt(120)
	if is, want := number.IntValue(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	number.SetInt(-40)
	if is, want := number.IntValue(), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
