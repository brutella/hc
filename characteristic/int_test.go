package characteristic

import (
	"testing"
)

func TestNumberIntOutOfBounds(t *testing.T) {
	number := NewInt(TypeBrightness)
	number.Format = FormatInt32
	number.Value = 2
	number.SetMinValue(0)
	number.SetMaxValue(100)
	number.SetStepValue(1)

	number.SetValue(120)
	if is, want := number.GetValue(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	number.SetValue(-40)
	if is, want := number.GetValue(), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
