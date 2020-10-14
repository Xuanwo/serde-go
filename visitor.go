package serde

import (
	"time"
)

type DummyVisitor struct {
}

// FIXME: returning error cloud be misleading, maybe we can generate those dummy functions?
func (d DummyVisitor) String() string {
	return "DummyVisitor"
}

func (d DummyVisitor) VisitNil() (err error) {
	return NewInvalidTypeError("nil", d)
}

func (d DummyVisitor) VisitBool(v bool) (err error) {
	return NewInvalidTypeError("bool", d)
}

func (d DummyVisitor) VisitInt(v int) (err error) {
	return NewInvalidTypeError("int", d)
}

func (d DummyVisitor) VisitInt8(v int8) (err error) {
	return NewInvalidTypeError("int8", d)
}

func (d DummyVisitor) VisitInt16(v int16) (err error) {
	return NewInvalidTypeError("int16", d)
}

func (d DummyVisitor) VisitInt32(v int32) (err error) {
	return NewInvalidTypeError("int32", d)
}

func (d DummyVisitor) VisitInt64(v int64) (err error) {
	return NewInvalidTypeError("int64", d)
}

func (d DummyVisitor) VisitUint(v uint) (err error) {
	return NewInvalidTypeError("nil", d)
}

func (d DummyVisitor) VisitUint8(v uint8) (err error) {
	return NewInvalidTypeError("uint8", d)
}

func (d DummyVisitor) VisitUint16(v uint16) (err error) {
	return NewInvalidTypeError("uint16", d)
}

func (d DummyVisitor) VisitUint32(v uint32) (err error) {
	return NewInvalidTypeError("uint32", d)
}

func (d DummyVisitor) VisitUint64(v uint64) (err error) {
	return NewInvalidTypeError("uint64", d)
}

func (d DummyVisitor) VisitFloat32(v float32) (err error) {
	return NewInvalidTypeError("float32", d)
}

func (d DummyVisitor) VisitFloat64(v float64) (err error) {
	return NewInvalidTypeError("float64", d)
}

func (d DummyVisitor) VisitComplex64(v complex64) (err error) {
	return NewInvalidTypeError("complex64", d)
}

func (d DummyVisitor) VisitComplex128(v complex128) (err error) {
	return NewInvalidTypeError("complex128", d)
}

func (d DummyVisitor) VisitRune(v rune) (err error) {
	return NewInvalidTypeError("rune", d)
}

func (d DummyVisitor) VisitString(v string) (err error) {
	return NewInvalidTypeError("string", d)
}

func (d DummyVisitor) VisitByte(v byte) (err error) {
	return NewInvalidTypeError("byte", d)
}

func (d DummyVisitor) VisitBytes(v []byte) (err error) {
	return NewInvalidTypeError("[]byte", d)
}

func (d DummyVisitor) VisitTime(v time.Time) (err error) {
	return NewInvalidTypeError("time.Time", d)
}

func (d DummyVisitor) VisitSlice(s SliceAccess) (err error) {
	return NewInvalidTypeError("slice", d)
}

func (d DummyVisitor) VisitMap(m MapAccess) (err error) {
	return NewInvalidTypeError("map", d)
}

type StringVisitor struct {
	v *string

	DummyVisitor
}

func NewStringVisitor(v *string) StringVisitor {
	return StringVisitor{v: v}
}

func (s StringVisitor) VisitString(v string) (err error) {
	*s.v = v
	return nil
}

func (s StringVisitor) VisitBytes(v []byte) (err error) {
	*s.v = string(v)
	return nil
}
