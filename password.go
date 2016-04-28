package hc

import (
	"bytes"
	"errors"
)

// NewPin returns a HomeKit compatible pin string from a 8-numbers strings e.g. '01020304'.
func NewPin(pin string) (string, error) {
	var fmtPin string
	if pin == "12345678" {
		return fmtPin, errors.New("Pin must not be 12345678")
	}

	if len(pin) != 8 {
		return fmtPin, errors.New("Pin must be 8 characters long")
	}
	bs := []byte(pin)
	for _, b := range bs {
		if b < byte('0') || b > byte('9') {
			return fmtPin, errors.New("Pin must only contain numbers")
		}
	}
	runes := bytes.Runes(bs)
	first := string(runes[:3])
	second := string(runes[3:5])
	third := string(runes[5:])
	fmtPin = first + "-" + second + "-" + third

	return fmtPin, nil
}
