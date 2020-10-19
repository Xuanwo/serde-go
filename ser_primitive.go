package serde

type StringSerializer string

func (s StringSerializer) Serialize(ser Serializer) (err error) {
	return ser.SerializeString(string(s))
}
