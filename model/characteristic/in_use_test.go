package characteristic

import (
	"reflect"
	"testing"
)

func TestInUse(t *testing.T) {
	use := NewInUse(true)

	if x := use.Permissions; reflect.DeepEqual(x, PermsRead()) == false {
		t.Fatal(x)
	}
	if is, want := use.InUse(), true; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	use.SetInUse(false)

	if is, want := use.InUse(), false; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
