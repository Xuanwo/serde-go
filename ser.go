package serde

type Serializable interface {
	Serialize(serializer Serializer) (err error)
}

type Serializer interface {
	SerializeBool(v bool) (err error)
	SerializeInt(v int) (err error)
	SerializeInt8(v int8) (err error)
	SerializeInt16(v int16) (err error)
	SerializeInt32(v int32) (err error)
	SerializeInt64(v int64) (err error)
	SerializeUint(v uint) (err error)
	SerializeUint8(v uint8) (err error)
	SerializeUint16(v uint16) (err error)
	SerializeUint32(v uint32) (err error)
	SerializeUint64(v uint64) (err error)
	SerializeFloat32(v float32) (err error)
	SerializeFloat64(v float64) (err error)
	SerializeComplex64(v complex64) (err error)
	SerializeComplex128(v complex128) (err error)
	SerializeString(v string) (err error)
	SerializeBytes(v []byte) (err error)

	SerializeSlice(length int) (s SliceSerializer, err error)
	SerializeMap(length int) (m MapSerializer, err error)
	SerializeStruct(name string, length int) (s StructSerializer, err error)
}

type SliceSerializer interface {
	SerializeElement(v Serializable) (err error)
	EndSlice() (err error)
}

type MapSerializer interface {
	SerializeEntry(k, v Serializable) (err error)
	EndMap() (err error)
}

type StructSerializer interface {
	SerializeField(k, v Serializable) (err error)
	EndStruct() (err error)
}
