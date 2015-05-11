package characteristic

import (
	"testing"
)

func TestName(t *testing.T) {
	n := NewName("Test")

	if is, want := n.Type, CharTypeName; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := n.Name(), "Test"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	n.SetName("My Name")

	if is, want := n.Name(), "My Name"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
