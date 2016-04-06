package characteristic

import (
	"testing"
)

func TestFloat(t *testing.T) {
	float := NewFloat(TypeBrightness)
	float.Format = FormatFloat
	float.Value = 20.2

	if is, want := float.GetValue(), 20.2; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	float.SetValue(10.1)

	if is, want := float.GetValue(), 10.1; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestNumberFloatOutOfBounds(t *testing.T) {
	float := NewFloat(TypeBrightness)
	float.Format = FormatFloat

	float.Value = 20.2
	float.SetMinValue(0.0)
	float.SetMaxValue(100.0)
	float.SetStepValue(0.1)

	float.SetValue(120)
	if is, want := float.GetValue(), 100.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	float.SetValue(-40)
	if is, want := float.GetValue(), 0.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
