package characteristic

import (
	"testing"
)

func TestFloat(t *testing.T) {
	float := NewFloat(20.2, PermsAll())

	if is, want := float.FloatValue(), 20.2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	float.SetFloat(10.1)

	if is, want := float.FloatValue(), 10.1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
