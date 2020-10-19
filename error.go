package serde

import (
	"bytes"
	"fmt"
)

type Unexpected interface {
	unexpected() string
}

type UnexpectedBool bool

func (u UnexpectedBool) unexpected() string {
	return fmt.Sprintf("boolean `%v`", u)
}

type UnexpectedInt int

func (u UnexpectedInt) unexpected() string {
	return fmt.Sprintf("int `%v`", u)
}

type UnexpectedInt8 int8

func (u UnexpectedInt8) unexpected() string {
	return fmt.Sprintf("int8 `%v`", u)
}

type UnexpectedInt16 int16

func (u UnexpectedInt16) unexpected() string {
	return fmt.Sprintf("int16 `%v`", u)
}

type UnexpectedInt32 int32

func (u UnexpectedInt32) unexpected() string {
	return fmt.Sprintf("int32 `%v`", u)
}

type UnexpectedInt64 int64

func (u UnexpectedInt64) unexpected() string {
	return fmt.Sprintf("int64 `%v`", u)
}

type UnexpectedUint uint

func (u UnexpectedUint) unexpected() string {
	return fmt.Sprintf("uint `%v`", u)
}

type UnexpectedUint8 uint8

func (u UnexpectedUint8) unexpected() string {
	return fmt.Sprintf("uint8 `%v`", u)
}

type UnexpectedUint16 uint16

func (u UnexpectedUint16) unexpected() string {
	return fmt.Sprintf("uint16 `%v`", u)
}

type UnexpectedUint32 uint32

func (u UnexpectedUint32) unexpected() string {
	return fmt.Sprintf("uint32 `%v`", u)
}

type UnexpectedUint64 uint64

func (u UnexpectedUint64) unexpected() string {
	return fmt.Sprintf("uint64 `%v`", u)
}

type UnexpectedFloat32 float32

func (u UnexpectedFloat32) unexpected() string {
	return fmt.Sprintf("float32 `%v`", u)
}

type UnexpectedFloat64 float64

func (u UnexpectedFloat64) unexpected() string {
	return fmt.Sprintf("float64 `%v`", u)
}

type UnexpectedComplex64 complex64

func (u UnexpectedComplex64) unexpected() string {
	return fmt.Sprintf("complex64 `%v`", u)
}

type UnexpectedComplex128 complex128

func (u UnexpectedComplex128) unexpected() string {
	return fmt.Sprintf("complex128 `%v`", u)
}

type UnexpectedRune rune

func (u UnexpectedRune) unexpected() string {
	return fmt.Sprintf("rune `%v`", u)
}

type UnexpectedString string

func (u UnexpectedString) unexpected() string {
	return fmt.Sprintf("string `%v`", u)
}

type UnexpectedByte byte

func (u UnexpectedByte) unexpected() string {
	return fmt.Sprintf("byte `%v`", u)
}

type UnexpectedBytes struct {
}

func (u UnexpectedBytes) unexpected() string {
	return "bytes"
}

type UnexpectedSlice struct {
}

func (u UnexpectedSlice) unexpected() string {
	return "slice"
}

type UnexpectedMap struct {
}

func (u UnexpectedMap) unexpected() string {
	return "map"
}

type UnexpectedStruct struct {
}

func (u UnexpectedStruct) unexpected() string {
	return "struct"
}

type UnexpectedNil struct {
}

func (u UnexpectedNil) unexpected() string {
	return "nil"
}

type ErrInvalidType struct {
	unexpected Unexpected
	expected   fmt.Stringer
}

func (e *ErrInvalidType) Error() string {
	return fmt.Sprintf("invalid type: %s, expected %s", e.unexpected.unexpected(), e.expected)
}

func NewInvalidTypeError(unexpected Unexpected, expected fmt.Stringer) *ErrInvalidType {
	return &ErrInvalidType{
		unexpected: unexpected,
		expected:   expected,
	}
}

type ErrInvalidValue struct {
	unexpected Unexpected
	expected   fmt.Stringer
}

func (e *ErrInvalidValue) Error() string {
	return fmt.Sprintf("invalid value: %s, expected %s", e.unexpected.unexpected(), e.expected)
}

type ErrInvalidLength struct {
	length   uint
	expected fmt.Stringer
}

func (e *ErrInvalidLength) Error() string {
	return fmt.Sprintf("invalid length: %d, expected %s", e.length, e.expected)
}

type ErrUnknownField struct {
	field    string
	expected []string
}

func (e *ErrUnknownField) Error() string {
	return fmt.Sprintf("unknown field `%s`, expected %s", e.field, oneOfFields(e.expected))
}

type oneOfFields []string

func (o oneOfFields) String() string {
	switch len(o) {
	case 0:
		return "there are no fields"
	case 1:
		return fmt.Sprintf("expected `%s`", o[0])
	case 2:
		return fmt.Sprintf("expected `%s` or `%s`", o[0], o[1])
	default:
		var s bytes.Buffer
		s.WriteString("expected one of ")
		for i, v := range o {
			if i > 0 {
				s.WriteString(", ")
			}
			s.WriteRune('`')
			s.WriteString(v)
			s.WriteRune('`')
		}
		return s.String()
	}
}

type ErrMissingField struct {
	field string
}

func (e *ErrMissingField) Error() string {
	return fmt.Sprintf("missing field `%s`", e.field)
}

type ErrDuplicateField struct {
	field string
}

func (e *ErrDuplicateField) Error() string {
	return fmt.Sprintf("duplicate field `%s`", e.field)
}
