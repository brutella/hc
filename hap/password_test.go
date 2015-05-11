package hap

import (
	"testing"
)

func TestPassword(t *testing.T) {
	pwd, err := NewPassword("00011222")
	if err != nil {
		t.Fatal(err)
	}
	if pwd != "000-11-222" {
		t.Fatal(pwd)
	}
}

func TestShortPassword(t *testing.T) {
	if _, err := NewPassword("0001122"); err == nil {
		t.Fatal("expected error")
	}
}

func TestLongPassword(t *testing.T) {
	if _, err := NewPassword("000112221"); err == nil {
		t.Fatal("expected error")
	}
}

func TestNonNumberPassword(t *testing.T) {
	if _, err := NewPassword("0001122a"); err == nil {
		t.Fatal("expected error")
	}
}

func TestInvalidPassword(t *testing.T) {
	if _, err := NewPassword("12345678"); err == nil {
		t.Fatal("expected error")
	}
}
