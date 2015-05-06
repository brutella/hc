package util

import (
	"crypto/rand"
)

// RandomHexString returns a random hex string.
func RandomHexString() string {
	var b [16]byte
	_, err := rand.Read(b[:])
	if err != nil {
		panic(err)
	}
	var out [32]byte
	for i := 0; i < len(b); i++ {
		out[i*2] = btoh((b[i] >> 4) & 0xF)
		out[i*2+1] = btoh(b[i] & 0xF)
	}
	return string(out[:])
}

func btoh(i byte) byte {
	if i > 9 {
		return 0x61 + (i - 10)
	}
	return 0x30 + i
}
