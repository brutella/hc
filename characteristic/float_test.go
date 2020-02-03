package characteristic

import (
	"testing"
)

func TestFloat(t *testing.T) {
	float := NewHue()
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
	float := NewHue()

	float.SetValue(400)
	if is, want := float.GetValue(), 360.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	float.SetValue(-40)
	if is, want := float.GetValue(), 0.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
