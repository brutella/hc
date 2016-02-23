package characteristic

import (
	"testing"
)

func TestNumber(t *testing.T) {
	number := NewNumber(20.2, 0, 100, 0.1, FormatFloat, PermsAll())

	if x := number.Format; x != FormatFloat {
		t.Fatal(x)
	}
	if is, want := number.GetValue(), 20.2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := number.GetMinValue(), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := number.GetMaxValue(), 100; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := number.GetMinStepValue(), 0.1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	number.SetNumber(17)
	number.SetMinValue(0.3)
	number.SetMaxValue(200)
	number.SetMinStepValue(0.3)

	if is, want := number.GetValue(), 17.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := number.GetMinValue(), 0.3; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := number.GetMaxValue(), 200; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := number.GetMinStepValue(), 0.3; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
