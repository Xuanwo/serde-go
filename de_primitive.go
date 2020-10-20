package serde

import (
	"errors"
)

type DummyVisitor struct {
	expect string
}

func (vi DummyVisitor) String() string {
	if vi.expect == "" {
		return "dummy"
	}

	return vi.expect
}

func NewDummyVisitor(expect string) DummyVisitor {
	return DummyVisitor{expect: expect}
}

type BoolVisitor struct {
	v *bool
}

func NewBoolVisitor(v *bool) BoolVisitor {
	return BoolVisitor{v: v}
}

func (vi BoolVisitor) VisitBool(v bool) (err error) {
	*vi.v = v
	return nil
}

type Float32Visitor struct {
	v *float32
}

func NewFloat32Visitor(v *float32) Float32Visitor {
	return Float32Visitor{v: v}
}

func (vi Float32Visitor) VisitFloat32(v float32) (err error) {
	*vi.v = v
	return nil
}

func (vi Float32Visitor) VisitFloat64(v float64) (err error) {
	if v < MinFloat32 || v > MaxFloat32 {
		return errors.New("overflow")
	}
	*vi.v = float32(v)
	return nil
}

type Float64Visitor struct {
	v *float64
}

func NewFloat64Visitor(v *float64) Float64Visitor {
	return Float64Visitor{v: v}
}

func (vi Float64Visitor) VisitFloat32(v float32) (err error) {
	*vi.v = float64(v)
	return nil
}

func (vi Float64Visitor) VisitFloat64(v float64) (err error) {
	*vi.v = v
	return nil
}

type Complex64Visitor struct {
	v *complex64
}

func NewComplex64Visitor(v *complex64) Complex64Visitor {
	return Complex64Visitor{v: v}
}

func (vi Complex64Visitor) VisitComplex64(v complex64) (err error) {
	*vi.v = v
	return nil
}

func (vi Complex64Visitor) VisitComplex128(v complex128) (err error) {
	r, i := real(v), imag(v)
	if r < MinFloat32 || r > MaxFloat32 || i < MinFloat32 || i > MaxFloat32 {
		return errors.New("overflow")
	}
	*vi.v = complex64(v)
	return nil
}

type Complex128Visitor struct {
	v *complex128
}

func NewComplex128Visitor(v *complex128) Complex128Visitor {
	return Complex128Visitor{v: v}
}

func (vi Complex128Visitor) VisitComplex64(v complex64) (err error) {
	*vi.v = complex128(v)
	return nil
}

func (vi Complex128Visitor) VisitComplex128(v complex128) (err error) {
	*vi.v = v
	return nil
}

type RuneVisitor struct {
	v *rune
}

func NewRuneVisitor(v *rune) RuneVisitor {
	return RuneVisitor{v: v}
}

func (vi RuneVisitor) VisitRune(v rune) (err error) {
	*vi.v = v
	return nil
}

func (vi RuneVisitor) VisitInt32(v int32) (err error) {
	*vi.v = v
	return nil
}

type StringVisitor struct {
	v *string
}

func NewStringVisitor(v *string) StringVisitor {
	return StringVisitor{v: v}
}

func (vi StringVisitor) String() string {
	return "string"
}

func (vi StringVisitor) VisitString(v string) (err error) {
	*vi.v = v
	return nil
}

func (vi StringVisitor) VisitBytes(v []byte) (err error) {
	*vi.v = string(v)
	return nil
}

type ByteVisitor struct {
	v *byte
}

func NewByteVisitor(v *byte) ByteVisitor {
	return ByteVisitor{v: v}
}

func (vi ByteVisitor) VisitByte(v byte) (err error) {
	*vi.v = v
	return nil
}

func (vi ByteVisitor) VisitUint8(v uint8) (err error) {
	*vi.v = v
	return nil
}

type BytesVisitor struct {
	v *[]byte
}

func NewBytesVisitor(v *[]byte) BytesVisitor {
	return BytesVisitor{v: v}
}

func (vi BytesVisitor) VisitBytes(v []byte) (err error) {
	*vi.v = v
	return nil
}
