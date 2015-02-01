package netio

import (
	"strings"
	"time"
)

// Returns the current time in RFC1123 format
// The date string has the suffix GMT instead of UTC
func CurrentRFC1123Date() string {
	date := time.Now().UTC().Format(time.RFC1123)
	return strings.Replace(date, "UTC", "GMT", 1)
}
