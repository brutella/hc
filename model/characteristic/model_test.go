package characteristic

import (
	"testing"
)

func TestModel(t *testing.T) {
	m := NewModel("Late 2014")

	if is, want := m.Type, TypeModel; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := m.Model(), "Late 2014"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	m.SetModel("Early 2015")

	if is, want := m.Model(), "Early 2015"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
