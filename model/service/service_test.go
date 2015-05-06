package service

import (
	"github.com/brutella/hc/model"
	"testing"
)

func TestService(t *testing.T) {
	s := New()

	// TODO(brutella): int64 cast should not be required!
	if is, want := s.GetID(), int64(model.InvalidID); is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := len(s.GetCharacteristics()), 0; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
