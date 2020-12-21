package tests

import (
	"fmt"

	"github.com/Xuanwo/serde-go"
)

func DeserializeFromInterfaces(vv []interface{}, v serde.Deserializable) error {
	de := &De{v: vv}
	return v.Deserialize(de)
}

type De struct {
	idx int
	v   []interface{}
}

func (d *De) next() interface{} {
	if d.idx > len(d.v) {
		panic(fmt.Errorf("idx overflow: length %d, actual %d", len(d.v), d.idx))
	}

	defer func() {
		d.idx += 1
	}()
	return d.v[d.idx]
}

func (d *De) peek() interface{} {
	if d.idx > len(d.v) {
		panic(fmt.Errorf("idx overflow: length %d, actual %d", len(d.v), d.idx))
	}

	return d.v[d.idx]
}

func (d *De) DeserializeAny(v serde.Visitor) (err error) {
	i := d.peek()

	switch i.(type) {
	case bool:
		err = d.DeserializeBool(v)
	case int:
		err = d.DeserializeInt(v)
	case int8:
		err = d.DeserializeInt8(v)
	case int16:
		err = d.DeserializeInt16(v)
	case int32:
		err = d.DeserializeInt32(v)
	case int64:
		err = d.DeserializeInt64(v)
	case uint:
		err = d.DeserializeUint(v)
	case uint8:
		err = d.DeserializeUint8(v)
	case uint16:
		err = d.DeserializeUint16(v)
	case uint32:
		err = d.DeserializeUint32(v)
	case uint64:
		err = d.DeserializeUint64(v)
	case float32:
		err = d.DeserializeFloat32(v)
	case float64:
		err = d.DeserializeFloat64(v)
	case complex64:
		err = d.DeserializeComplex64(v)
	case complex128:
		err = d.DeserializeComplex128(v)
	case string:
		err = d.DeserializeString(v)
	case []byte:
		err = d.DeserializeBytes(v)
	case testSlice:
		err = d.DeserializeSlice(v)
	case testMap:
		err = d.DeserializeMap(v)
	case nil:
		err = d.DeserializeNil(v)
	default:
		err = fmt.Errorf("not supported type: %#v", i)
	}

	return err
}

func (d *De) DeserializeNil(v serde.Visitor) (err error) {
	d.next()
	return v.VisitNil()
}

func (d *De) DeserializeBool(v serde.Visitor) (err error) {
	return v.VisitBool(d.next().(bool))
}

func (d *De) DeserializeInt(v serde.Visitor) (err error) {
	return v.VisitInt(d.next().(int))
}

func (d *De) DeserializeInt8(v serde.Visitor) (err error) {
	return v.VisitInt8(d.next().(int8))
}

func (d *De) DeserializeInt16(v serde.Visitor) (err error) {
	return v.VisitInt16(d.next().(int16))
}

func (d *De) DeserializeInt32(v serde.Visitor) (err error) {
	return v.VisitInt32(d.next().(int32))
}

func (d *De) DeserializeInt64(v serde.Visitor) (err error) {
	return v.VisitInt64(d.next().(int64))
}

func (d *De) DeserializeUint(v serde.Visitor) (err error) {
	return v.VisitUint(d.next().(uint))
}

func (d *De) DeserializeUint8(v serde.Visitor) (err error) {
	return v.VisitUint8(d.next().(uint8))
}

func (d *De) DeserializeUint16(v serde.Visitor) (err error) {
	return v.VisitUint16(d.next().(uint16))
}

func (d *De) DeserializeUint32(v serde.Visitor) (err error) {
	return v.VisitUint32(d.next().(uint32))
}

func (d *De) DeserializeUint64(v serde.Visitor) (err error) {
	return v.VisitUint64(d.next().(uint64))
}

func (d *De) DeserializeFloat32(v serde.Visitor) (err error) {
	return v.VisitFloat32(d.next().(float32))
}

func (d *De) DeserializeFloat64(v serde.Visitor) (err error) {
	return v.VisitFloat64(d.next().(float64))
}

func (d *De) DeserializeComplex64(v serde.Visitor) (err error) {
	return v.VisitComplex64(d.next().(complex64))
}

func (d *De) DeserializeComplex128(v serde.Visitor) (err error) {
	return v.VisitComplex128(d.next().(complex128))
}

func (d *De) DeserializeString(v serde.Visitor) (err error) {
	return v.VisitString(d.next().(string))
}

func (d *De) DeserializeBytes(v serde.Visitor) (err error) {
	return v.VisitBytes(d.next().([]byte))
}

func (d *De) DeserializeSlice(v serde.Visitor) (err error) {
	x := d.next()
	if _, ok := x.(testSlice); !ok {
		return fmt.Errorf("value should be a map, but it's %#v", x)
	}

	return v.VisitSlice(d)
}

func (d *De) DeserializeMap(v serde.Visitor) (err error) {
	x := d.next()
	tm, ok := x.(testMap)
	if !ok {
		return fmt.Errorf("value should be a map, but it's %#v", x)
	}
	if tm != false {
		return fmt.Errorf("testMap should be false, instead of %#v", tm)
	}

	return v.VisitMap(d)
}

func (d *De) DeserializeStruct(name string, fields []string, v serde.Visitor) (err error) {
	return d.DeserializeMap(v)
}

func (d *De) NextKey(v serde.Visitor) (ok bool, err error) {
	x := d.peek()

	if tm, ok := x.(testMap); ok {
		if tm == true {
			_ = d.next()
			return false, nil
		}
	}
	return true, d.DeserializeAny(v)
}

func (d *De) NextValue(v serde.Visitor) (err error) {
	return d.DeserializeAny(v)
}

func (d *De) NextElement(v serde.Visitor) (ok bool, err error) {
	x := d.peek()

	if tm, ok := x.(testSlice); ok {
		if tm == true {
			_ = d.next()
			return false, nil
		}
	}
	return true, d.DeserializeAny(v)
}
