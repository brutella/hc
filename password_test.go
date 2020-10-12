package hc

import (
	"testing"
)

func TestPin(t *testing.T) {
	pwd, err := ValidatePin("00011222")
	if err != nil {
		t.Fatal(err)
	}
	if pwd != "000-11-222" {
		t.Fatal(pwd)
	}
}

func TestShortPin(t *testing.T) {
	if _, err := ValidatePin("0001122"); err == nil {
		t.Fatal("expected error")
	}
}

func TestLongPin(t *testing.T) {
	if _, err := ValidatePin("000112221"); err == nil {
		t.Fatal("expected error")
	}
}

func TestNonNumberPin(t *testing.T) {
	if _, err := ValidatePin("0001122a"); err == nil {
		t.Fatal("expected error")
	}
}

func TestInvalidPins(t *testing.T) {
	if _, err := ValidatePin("12345678"); err == nil {
		t.Fatal("expected error")
	}

	if _, err := ValidatePin("87654321"); err == nil {
		t.Fatal("expected error")
	}

	if _, err := ValidatePin("11111111"); err == nil {
		t.Fatal("expected error")
	}

	if _, err := ValidatePin("99999999"); err == nil {
		t.Fatal("expected error")
	}
}
