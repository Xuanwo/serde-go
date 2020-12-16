package tests

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)

var _ = errors.New

type serdeStructEnum_TagSupport = int

const (
	serdeStructEnumSkip_TagSupport serdeStructEnum_TagSupport = 0

	serdeStructEnum_TagSupport_skipSerialize serdeStructEnum_TagSupport = 2

	serdeStructEnum_TagSupport_skipDeserialize serdeStructEnum_TagSupport = 3
)

type serdeStructFieldVisitor_TagSupport struct {
	e serdeStructEnum_TagSupport

	serde.DummyVisitor
}

func serdeNewStructFieldVisitor_TagSupport() *serdeStructFieldVisitor_TagSupport {
	return &serdeStructFieldVisitor_TagSupport{
		DummyVisitor: serde.NewDummyVisitor("TagSupport Field"),
	}
}

func (s *serdeStructFieldVisitor_TagSupport) VisitString(v string) (err error) {
	switch v {

	case "skip":
		s.e = serdeStructEnumSkip_TagSupport

	case "skipSerialize":
		s.e = serdeStructEnum_TagSupport_skipSerialize

	case "skipDeserialize":
		s.e = serdeStructEnumSkip_TagSupport

	default:
		return errors.New("invalid field")
	}
	return nil
}

type serdeStructVisitor_TagSupport struct {
	v *TagSupport

	serde.DummyVisitor
}

func serdeNewStructVisitor_TagSupport(v *TagSupport) *serdeStructVisitor_TagSupport {
	return &serdeStructVisitor_TagSupport{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("TagSupport"),
	}
}

func (s *serdeStructVisitor_TagSupport) VisitMap(m serde.MapAccess) (err error) {
	field := serdeNewStructFieldVisitor_TagSupport()
	for {
		ok, err := m.NextKey(field)
		if !ok {
			break
		}
		if err != nil {
			return err
		}

		var v serde.Visitor
		switch field.e {
		case serdeStructEnumSkip_TagSupport:
			v = serde.SkipVisitor{}

		case serdeStructEnum_TagSupport_skipSerialize:
			v = serde.NewIntVisitor(&s.v.skipSerialize)

		default:
			return errors.New("invalid field")
		}
		err = m.NextValue(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *TagSupport) Deserialize(de serde.Deserializer) (err error) {
	return de.DeserializeStruct("TagSupport", nil, serdeNewStructVisitor_TagSupport(s))
}

func (s *TagSupport) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("TagSupport", 3)
	if err != nil {
		return err
	}

	err = st.SerializeField(
		serde.StringSerializer("skipDeserialize"),
		serde.IntSerializer(s.skipDeserialize),
	)
	if err != nil {
		return
	}

	err = st.EndStruct()
	if err != nil {
		return
	}
	return nil
}
