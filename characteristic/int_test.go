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

func TestNumberIntOutOfValidValues(t *testing.T) {
	number := NewInt("")
	number.Format = FormatUInt8
	number.Perms = PermsAll()
	number.SetValidValues([]int{
		0,
		10,
	})

	number.SetValue(5)
	if is, want := number.GetValue(), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestNumberIntOutOfValidValuesRange(t *testing.T) {
	number := NewInt("")
	number.Format = FormatUInt8
	number.Perms = PermsAll()
	number.SetValidValuesRange(0, 100)

	number.SetValue(120)
	if is, want := number.GetValue(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	number.SetValue(-40)
	if is, want := number.GetValue(), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
