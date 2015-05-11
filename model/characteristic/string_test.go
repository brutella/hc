package characteristic

import (
	"testing"
)

func TestString(t *testing.T) {
	str := NewString("A String")
	if is, want := str.StringValue(), "A String"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	str.SetString("My String")

	if is, want := str.StringValue(), "My String"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
