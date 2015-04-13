package hap

import (
	"bytes"
	"errors"
)

// NewPassword creates a HomeKit compatible password string used for the bridge pairing setup.
// The argument string must be excatly 8 characters e.g. '01020304'
func NewPassword(str string) (string, error) {
	var password string
	if str == "12345678" {
		return password, errors.New("Password must not be 12345678")
	}

	if len(str) != 8 {
		return password, errors.New("Password must be 8 characters long")
	}
	bs := []byte(str)
	for _, b := range bs {
		if b < byte('0') || b > byte('9') {
			return password, errors.New("Password must only contain numbers")
		}
	}
	runes := bytes.Runes(bs)
	first := string(runes[:3])
	second := string(runes[3:5])
	third := string(runes[5:])
	password = first + "-" + second + "-" + third

	return password, nil
}
