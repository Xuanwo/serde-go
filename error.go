package serde

import (
	"fmt"
)

type ErrInvalidType struct {
	unexpected string
	expected   fmt.Stringer
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("invalid type: %s, expected %s", e.unexpected, e.expected)
}

func NewInvalidTypeError(unexpected string, expected fmt.Stringer) *ErrInvalidType {
	return &ErrInvalidType{
		unexpected: unexpected,
		expected:   expected,
	}
}
