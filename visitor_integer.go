package serde

import (
	"errors"
)

type Int8Visitor struct {
	v *int8

	DummyVisitor
}

func NewInt8Visitor(v *int8) Int8Visitor {
	return Int8Visitor{v: v}
}

func (s Int8Visitor) VisitInt8(v int8) (err error) {
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitUint8(v uint8) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitInt16(v int16) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitUint16(v uint16) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitInt32(v int32) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitInt(v int) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitUint(v uint) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitInt64(v int64) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	if v < MinInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

func (s Int8Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt8 {
		return errors.New("overflow")
	}
	*s.v = int8(v)
	return nil
}

type Uint8Visitor struct {
	v *uint8

	DummyVisitor
}

func NewUint8Visitor(v *uint8) Uint8Visitor {
	return Uint8Visitor{v: v}
}

func (s Uint8Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitUint8(v uint8) (err error) {
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitInt16(v int16) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitUint16(v uint16) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitInt32(v int32) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitInt(v int) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitUint(v uint) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitInt64(v int64) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	if v < MinUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

func (s Uint8Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint8 {
		return errors.New("overflow")
	}
	*s.v = uint8(v)
	return nil
}

type Int16Visitor struct {
	v *int16

	DummyVisitor
}

func NewInt16Visitor(v *int16) Int16Visitor {
	return Int16Visitor{v: v}
}

func (s Int16Visitor) VisitInt8(v int8) (err error) {
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitUint8(v uint8) (err error) {
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitInt16(v int16) (err error) {
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitUint16(v uint16) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitInt32(v int32) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	if v < MinInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitInt(v int) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	if v < MinInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitUint(v uint) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitInt64(v int64) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	if v < MinInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

func (s Int16Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt16 {
		return errors.New("overflow")
	}
	*s.v = int16(v)
	return nil
}

type Uint16Visitor struct {
	v *uint16

	DummyVisitor
}

func NewUint16Visitor(v *uint16) Uint16Visitor {
	return Uint16Visitor{v: v}
}

func (s Uint16Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitUint8(v uint8) (err error) {
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitInt16(v int16) (err error) {
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitUint16(v uint16) (err error) {
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitInt32(v int32) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitInt(v int) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitUint(v uint) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitInt64(v int64) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	if v < MinUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

func (s Uint16Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint16 {
		return errors.New("overflow")
	}
	*s.v = uint16(v)
	return nil
}

type Int32Visitor struct {
	v *int32

	DummyVisitor
}

func NewInt32Visitor(v *int32) Int32Visitor {
	return Int32Visitor{v: v}
}

func (s Int32Visitor) VisitInt8(v int8) (err error) {
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitUint8(v uint8) (err error) {
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitInt16(v int16) (err error) {
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitUint16(v uint16) (err error) {
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitInt32(v int32) (err error) {
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitInt(v int) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	if v < MinInt32 {
		return errors.New("overflow")
	}
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitUint32(v uint32) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitUint(v uint) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitInt64(v int64) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	if v < MinInt32 {
		return errors.New("overflow")
	}
	*s.v = int32(v)
	return nil
}

func (s Int32Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt32 {
		return errors.New("overflow")
	}
	*s.v = int32(v)
	return nil
}

type IntVisitor struct {
	v *int

	DummyVisitor
}

func NewIntVisitor(v *int) IntVisitor {
	return IntVisitor{v: v}
}

func (s IntVisitor) VisitInt8(v int8) (err error) {
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitUint8(v uint8) (err error) {
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitInt16(v int16) (err error) {
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitUint16(v uint16) (err error) {
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitInt32(v int32) (err error) {
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitInt(v int) (err error) {
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitUint32(v uint32) (err error) {
	if UintSize == 32 && v > MaxInt32 {
		return errors.New("overflow")
	}
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitUint(v uint) (err error) {
	if v > MaxInt {
		return errors.New("overflow")
	}
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitInt64(v int64) (err error) {
	if v > MaxInt {
		return errors.New("overflow")
	}
	if v < MinInt {
		return errors.New("overflow")
	}
	*s.v = int(v)
	return nil
}

func (s IntVisitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt {
		return errors.New("overflow")
	}
	*s.v = int(v)
	return nil
}

type Uint32Visitor struct {
	v *uint32

	DummyVisitor
}

func NewUint32Visitor(v *uint32) Uint32Visitor {
	return Uint32Visitor{v: v}
}

func (s Uint32Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitUint8(v uint8) (err error) {
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitInt16(v int16) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitUint16(v uint16) (err error) {
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitInt32(v int32) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitInt(v int) (err error) {
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitUint32(v uint32) (err error) {
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitUint(v uint) (err error) {
	if v > MaxUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitInt64(v int64) (err error) {
	if v > MaxUint32 {
		return errors.New("overflow")
	}
	if v < MinUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

func (s Uint32Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint32 {
		return errors.New("overflow")
	}
	*s.v = uint32(v)
	return nil
}

type UintVisitor struct {
	v *uint

	DummyVisitor
}

func NewUintVisitor(v *uint) UintVisitor {
	return UintVisitor{v: v}
}

func (s UintVisitor) VisitInt8(v int8) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitUint8(v uint8) (err error) {
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitInt16(v int16) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitUint16(v uint16) (err error) {
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitInt32(v int32) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitInt(v int) (err error) {
	if v < MinUint {
		return errors.New("overflow")
	}
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitUint32(v uint32) (err error) {
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitUint(v uint) (err error) {
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitInt64(v int64) (err error) {
	if UintSize == 32 && v > MaxUint32 {
		return errors.New("overflow")
	}
	if v < MinUint {
		return errors.New("overflow")
	}
	*s.v = uint(v)
	return nil
}

func (s UintVisitor) VisitUint64(v uint64) (err error) {
	if v > MaxUint {
		return errors.New("overflow")
	}
	*s.v = uint(v)
	return nil
}

type Int64Visitor struct {
	v *int64

	DummyVisitor
}

func NewInt64Visitor(v *int64) Int64Visitor {
	return Int64Visitor{v: v}
}

func (s Int64Visitor) VisitInt8(v int8) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitUint8(v uint8) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitInt16(v int16) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitUint16(v uint16) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitInt32(v int32) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitInt(v int) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitUint32(v uint32) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitUint(v uint) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitInt64(v int64) (err error) {
	*s.v = int64(v)
	return nil
}

func (s Int64Visitor) VisitUint64(v uint64) (err error) {
	if v > MaxInt64 {
		return errors.New("overflow")
	}
	*s.v = int64(v)
	return nil
}

type Uint64Visitor struct {
	v *uint64

	DummyVisitor
}

func NewUint64Visitor(v *uint64) Uint64Visitor {
	return Uint64Visitor{v: v}
}

func (s Uint64Visitor) VisitInt8(v int8) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitUint8(v uint8) (err error) {
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitInt16(v int16) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitUint16(v uint16) (err error) {
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitInt32(v int32) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitInt(v int) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitUint32(v uint32) (err error) {
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitUint(v uint) (err error) {
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitInt64(v int64) (err error) {
	if v < MinUint64 {
		return errors.New("overflow")
	}
	*s.v = uint64(v)
	return nil
}

func (s Uint64Visitor) VisitUint64(v uint64) (err error) {
	*s.v = uint64(v)
	return nil
}
