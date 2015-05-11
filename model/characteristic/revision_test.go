package characteristic

import (
	"testing"
)

func TestHardwareRevisionCharacteristic(t *testing.T) {
	hw := NewHardwareRevision("1.0")

	if is, want := hw.Type, CharTypeHardwareRevision; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
	if is, want := hw.Revision(), "1.0"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}

	hw.SetRevision("1.0.1")

	if is, want := hw.Revision(), "1.0.1"; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestFirmwareRevisionCharacteristic(t *testing.T) {
	hw := NewFirmwareRevision("1.0")

	if is, want := hw.Type, CharTypeFirmwareRevision; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}

func TestSoftwareRevisionCharacteristic(t *testing.T) {
	hw := NewSoftwareRevision("1.0")

	if is, want := hw.Type, CharTypeSoftwareRevision; is != want {
		t.Fatalf("is=%v want=%v", is, want)
	}
}
