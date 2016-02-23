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

func TestNumberFloatOutOfBounds(t *testing.T) {
	number := NewFloatMinMaxSteps(20.2, 0, 100, 0.1, PermsAll())

	number.SetFloat(120)
	if is, want := number.FloatValue(), 100.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	number.SetFloat(-40)
	if is, want := number.FloatValue(), 0.0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
