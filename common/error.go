package common

import (
	"errors"
	"fmt"
)

// NewErrorf returns a new error. Arguments are handled in the manner of fmt.Sprintf.
func NewErrorf(format string, a ...interface{}) error {
	return errors.New(fmt.Sprintf(format, a))
}

// NewError returns a new error using erros.New()
func NewError(message string) error {
	return errors.New(message)
}
