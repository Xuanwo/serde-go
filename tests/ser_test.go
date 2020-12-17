package tests

import (
	"github.com/Xuanwo/serde-go"
)

func SerializeToInterfaces(v serde.Serializable) ([]interface{}, error) {
	ser := &Ser{v: make([]interface{}, 0)}

	err := v.Serialize(ser)
	if err != nil {
		return nil, err
	}

	return ser.v, nil
}

type Ser struct {
	v []interface{}
}

func (s *Ser) append(x interface{}) {
	s.v = append(s.v, x)
}

func (s *Ser) SerializeNil() (err error) {
	s.append(nil)
	return
}

func (s *Ser) SerializeBool(v bool) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeInt(v int) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeInt8(v int8) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeInt16(v int16) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeInt32(v int32) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeInt64(v int64) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeUint(v uint) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeUint8(v uint8) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeUint16(v uint16) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeUint32(v uint32) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeUint64(v uint64) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeFloat32(v float32) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeFloat64(v float64) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeComplex64(v complex64) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeComplex128(v complex128) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeString(v string) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeBytes(v []byte) (err error) {
	s.append(v)
	return
}

func (s *Ser) SerializeSlice(length int) (ss serde.SliceSerializer, err error) {
	s.append(testSlice(false))
	return s, nil
}

func (s *Ser) SerializeMap(length int) (m serde.MapSerializer, err error) {
	s.append(testMap(false))
	return s, nil
}

func (s *Ser) SerializeStruct(name string, length int) (ss serde.StructSerializer, err error) {
	s.append(testMap(false))
	return s, nil
}

func (s *Ser) SerializeElement(v serde.Serializable) (err error) {
	return v.Serialize(s)
}

func (s *Ser) EndSlice() (err error) {
	s.append(testSlice(true))
	return
}

func (s *Ser) SerializeEntry(k, v serde.Serializable) (err error) {
	err = k.Serialize(s)
	if err != nil {
		return err
	}
	return v.Serialize(s)
}

func (s *Ser) EndMap() (err error) {
	s.append(testMap(true))
	return
}

func (s *Ser) SerializeField(k, v serde.Serializable) (err error) {
	err = k.Serialize(s)
	if err != nil {
		return err
	}
	return v.Serialize(s)
}

func (s *Ser) EndStruct() (err error) {
	s.append(testMap(true))
	return
}

type testSlice bool
type testMap bool
