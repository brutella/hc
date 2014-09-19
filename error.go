package gohap

import (
    "errors"
    "fmt"
)

func NewErrorf(format string, a ...interface{}) error {
    return  errors.New(fmt.Sprintf(format, a))
}

func NewError(message string) error {
    return  errors.New(message)
}
