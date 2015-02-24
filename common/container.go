package common

import (
	"bytes"
)

// Container is a dictionary using byte keys and values.
type Container interface {

	// SetByte sets one byte for a key
	SetByte(key, value byte)

	// SetBytes sets bytes for a key
	SetBytes(key byte, value []byte)

	// SetString sets a string for a key
	SetString(key byte, value string)

	// GetByte returns one byte for a key
	GetByte(key byte) byte

	// GetBytes returns bytes for a key
	GetBytes(key byte) []byte

	// GetString return a string for a key
	GetString(key byte) string

	// BytesBuffer returns the raw bytes
	BytesBuffer() *bytes.Buffer
}
