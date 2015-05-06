package hap

import (
	"bytes"
	"errors"
)

// NewPassword returns a HomeKit compatible password string from a 8-numbers strings e.g. '01020304'.
func NewPassword(password string) (string, error) {
	var fmtPassword string
	if password == "12345678" {
		return fmtPassword, errors.New("fmtPassword must not be 12345678")
	}

	if len(password) != 8 {
		return fmtPassword, errors.New("fmtPassword must be 8 characters long")
	}
	bs := []byte(password)
	for _, b := range bs {
		if b < byte('0') || b > byte('9') {
			return fmtPassword, errors.New("fmtPassword must only contain numbers")
		}
	}
	runes := bytes.Runes(bs)
	first := string(runes[:3])
	second := string(runes[3:5])
	third := string(runes[5:])
	fmtPassword = first + "-" + second + "-" + third

	return fmtPassword, nil
}
