package serde

type BoolIntSerializer bool

func (s BoolIntSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeBool(bool(s))
}

type IntSerializer int

func (s IntSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeInt(int(s))
}

type Int8Serializer int8

func (s Int8Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeInt8(int8(s))
}

type Int16Serializer int16

func (s Int16Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeInt16(int16(s))
}

type Int32Serializer int32

func (s Int32Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeInt32(int32(s))
}

type Int64Serializer int64

func (s Int64Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeInt64(int64(s))
}

type UintSerializer uint

func (s UintSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeUint(uint(s))
}

type Uint8Serializer uint8

func (s Uint8Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeUint8(uint8(s))
}

type Uint16Serializer uint16

func (s Uint16Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeUint16(uint16(s))
}

type Uint32Serializer uint32

func (s Uint32Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeUint32(uint32(s))
}

type Uint64Serializer uint64

func (s Uint64Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeUint64(uint64(s))
}

type Float32Serializer float32

func (s Float32Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeFloat32(float32(s))
}

type Float64Serializer float64

func (s Float64Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeFloat64(float64(s))
}

type Complex64Serializer complex64

func (s Complex64Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeComplex64(complex64(s))
}

type Complex128Serializer complex128

func (s Complex128Serializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeComplex128(complex128(s))
}

type RuneSerializer rune

func (s RuneSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeRune(rune(s))
}

type StringSerializer string

func (s StringSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeString(string(s))
}

type ByteSerializer byte

func (s ByteSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeByte(byte(s))
}

type BytesSerializer []byte

func (s BytesSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeBytes(s)
}
