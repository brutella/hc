package util

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// MAC48Address returns a MAC-48-like address from the argument string
func MAC48Address(input string) string {
	h := md5.New()
	h.Write([]byte(input))
	result := h.Sum(nil)

	var c []string
	c = append(c, toHex(result[0]))
	c = append(c, toHex(result[1]))
	c = append(c, toHex(result[2]))
	c = append(c, toHex(result[3]))
	c = append(c, toHex(result[4]))
	c = append(c, toHex(result[5]))

	// setup id needs the mac address in upper case
	return strings.ToUpper(strings.Join(c, ":"))
}

func toHex(b byte) string {
	return hex.EncodeToString([]byte{b})
}
