package tests

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)

var _ = errors.New

type serdeMapVisitor_string_string struct {
	v *map[string]string

	serde.DummyVisitor
}

func serdeNewMapVisitor_string_string(v *map[string]string) *serdeMapVisitor_string_string {
	if *v == nil {
		*v = make(map[string]string)
	}
	return &serdeMapVisitor_string_string{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("map[string]string"),
	}
}

func (s *serdeMapVisitor_string_string) VisitMap(m serde.MapAccess) (err error) {
	var field string
	var value string
	for {
		ok, err := m.NextKey(serde.NewStringVisitor(&field))
		if !ok {
			break
		}
		if err != nil {
			return err
		}
		err = m.NextValue(serde.NewStringVisitor(&value))
		if err != nil {
			return err
		}
		(*s.v)[field] = value
	}
	return nil
}

type serdeMapSerializer_string_string map[string]string

func (s serdeMapSerializer_string_string) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeMap(len(s))
	if err != nil {
		return err
	}

	for k, v := range s {
		err = st.SerializeEntry(
			serde.StringSerializer(k),
			serde.StringSerializer(v),
		)
		if err != nil {
			return
		}
	}

	err = st.EndMap()
	if err != nil {
		return
	}
	return nil
}

type serdeStructEnum_TypeSupport = int

const (
	serdeStructEnumSkip_TypeSupport serdeStructEnum_TypeSupport = 0

	serdeStructEnum_TypeSupport_vint serdeStructEnum_TypeSupport = 1

	serdeStructEnum_TypeSupport_vmap serdeStructEnum_TypeSupport = 2
)

type serdeStructFieldVisitor_TypeSupport struct {
	e serdeStructEnum_TypeSupport

	serde.DummyVisitor
}

func serdeNewStructFieldVisitor_TypeSupport() *serdeStructFieldVisitor_TypeSupport {
	return &serdeStructFieldVisitor_TypeSupport{
		DummyVisitor: serde.NewDummyVisitor("TypeSupport Field"),
	}
}

func (s *serdeStructFieldVisitor_TypeSupport) VisitString(v string) (err error) {
	switch v {

	case "vint":
		s.e = serdeStructEnum_TypeSupport_vint

	case "vmap":
		s.e = serdeStructEnum_TypeSupport_vmap

	default:
		return errors.New("invalid field")
	}
	return nil
}

type serdeStructVisitor_TypeSupport struct {
	v *TypeSupport

	serde.DummyVisitor
}

func serdeNewStructVisitor_TypeSupport(v *TypeSupport) *serdeStructVisitor_TypeSupport {
	return &serdeStructVisitor_TypeSupport{
		v:            v,
		DummyVisitor: serde.NewDummyVisitor("TypeSupport"),
	}
}

func (s *serdeStructVisitor_TypeSupport) VisitMap(m serde.MapAccess) (err error) {
	field := serdeNewStructFieldVisitor_TypeSupport()
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
		case serdeStructEnumSkip_TypeSupport:
			v = serde.SkipVisitor{}

		case serdeStructEnum_TypeSupport_vint:
			v = serde.NewIntVisitor(&s.v.vint)

		case serdeStructEnum_TypeSupport_vmap:
			v = serdeNewMapVisitor_string_string(&s.v.vmap)

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

func (s *TypeSupport) Deserialize(de serde.Deserializer) (err error) {
	return de.DeserializeStruct("TypeSupport", nil, serdeNewStructVisitor_TypeSupport(s))
}

func (s *TypeSupport) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("TypeSupport", 2)
	if err != nil {
		return err
	}

	err = st.SerializeField(
		serde.StringSerializer("vint"),
		serde.IntSerializer(s.vint),
	)
	if err != nil {
		return
	}

	err = st.SerializeField(
		serde.StringSerializer("vmap"),
		serdeMapSerializer_string_string(s.vmap),
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
