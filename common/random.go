package util

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

// RandomHexString returns a random hex string.
func RandomHexString() string {
	t := time.Now().Format(time.RFC3339Nano)

	h := md5.New()
	h.Write([]byte(t))
	result := h.Sum(nil)

	return hex.EncodeToString(result)
}
