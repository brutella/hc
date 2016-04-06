package characteristic

import (
	"testing"
)

func TestString(t *testing.T) {
	str := NewString(TypeName)
	str.Value = "A String"

	if is, want := str.GetValue(), "A String"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	str.SetValue("My String")

	if is, want := str.GetValue(), "My String"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
