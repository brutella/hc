package hc

import (
	"bytes"
	"errors"
	"fmt"
)

var invalidPins = []string{"12345678", "87654321", "00000000", "11111111", "22222222", "33333333", "44444444", "55555555", "66666666", "77777777", "88888888", "99999999"}

// ValidatePin validates a HomeKit pin.
func ValidatePin(pin string) (string, error) {
	var fmtPin string
	for _, invalidPin := range invalidPins {
		if pin == invalidPin {
			return fmtPin, fmt.Errorf("Pin must not be %s", pin)
		}
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
