package serde

type Deserializable interface {
	Deserialize(deserializer Deserializer) (err error)
}

type Deserializer interface {
	DeserializeAny(v Visitor) (err error)

	DeserializeBool(v Visitor) (err error)
	DeserializeInt(v Visitor) (err error)
	DeserializeInt8(v Visitor) (err error)
	DeserializeInt16(v Visitor) (err error)
	DeserializeInt32(v Visitor) (err error)
	DeserializeInt64(v Visitor) (err error)
	DeserializeUint(v Visitor) (err error)
	DeserializeUint8(v Visitor) (err error)
	DeserializeUint16(v Visitor) (err error)
	DeserializeUint32(v Visitor) (err error)
	DeserializeUint64(v Visitor) (err error)
	DeserializeFloat32(v Visitor) (err error)
	DeserializeFloat64(v Visitor) (err error)
	DeserializeComplex64(v Visitor) (err error)
	DeserializeComplex128(v Visitor) (err error)
	DeserializeRune(v Visitor) (err error)
	DeserializeString(v Visitor) (err error)
	DeserializeByte(v Visitor) (err error)
	DeserializeBytes(v Visitor) (err error)

	DeserializeSlice(v Visitor) (err error)
	DeserializeMap(v Visitor) (err error)
	DeserializeStruct(name string, fields []string, v Visitor) (err error)
}

type Visitor interface {
	VisitNil() (err error)

	VisitBool(v bool) (err error)
	VisitInt(v int) (err error)
	VisitInt8(v int8) (err error)
	VisitInt16(v int16) (err error)
	VisitInt32(v int32) (err error)
	VisitInt64(v int64) (err error)
	VisitUint(v uint) (err error)
	VisitUint8(v uint8) (err error)
	VisitUint16(v uint16) (err error)
	VisitUint32(v uint32) (err error)
	VisitUint64(v uint64) (err error)
	VisitFloat32(v float32) (err error)
	VisitFloat64(v float64) (err error)
	VisitComplex64(v complex64) (err error)
	VisitComplex128(v complex128) (err error)
	VisitRune(v rune) (err error)
	VisitString(v string) (err error)
	VisitByte(v byte) (err error)
	VisitBytes(v []byte) (err error)

	VisitSlice(s SliceAccess) (err error)
	VisitMap(m MapAccess) (err error)
}

type MapAccess interface {
	NextKey(v Visitor) (ok bool, err error)
	NextValue(v Visitor) (err error)
}

type SliceAccess interface {
	NextElement(v Visitor) (ok bool, err error)
}
