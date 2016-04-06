package characteristic

import (
	"testing"
)

func TestName(t *testing.T) {
	n := NewName()
	n.Value = "Test"

	if is, want := n.Type, TypeName; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := n.GetValue(), "Test"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	n.SetValue("My Name")

	if is, want := n.GetValue(), "My Name"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
