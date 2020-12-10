// +build tools

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"log"
	"strings"
	"text/template"

	"github.com/Xuanwo/templateutils"
)

const SerdePrefix = "// serde:"

type serdeType interface {
	Name() string
	Visitor() string
	NewVisitor() string
	Serializer() string
	Generate() string
}

type serdeState struct {
	todo      []serdeType
	generated map[string]struct{}
}

func (s *serdeState) AppendTodo(g serdeType) {
	s.todo = append(s.todo, g)
}

func (s *serdeState) Generated(g serdeType) {
	s.generated[g.Name()] = struct{}{}
}

func (s *serdeState) IsGenerated(g serdeType) bool {
	_, ok := s.generated[g.Name()]
	return ok
}

func (s *serdeState) NeedGenerate() bool {
	return len(s.todo) > 0
}

func newSerdeState() *serdeState {
	return &serdeState{
		todo:      make([]serdeType, 0),
		generated: make(map[string]struct{}),
	}
}

type serdeStruct struct {
	structType

	Fields []structField
	Flags  map[string]bool

	comments *ast.CommentGroup
	decl     *ast.StructType
}

func newSerdeStruct(name string, comments *ast.CommentGroup, decl *ast.StructType) serdeStruct {
	s := serdeStruct{
		structType: structType(name),
		comments:   comments,
		decl:       decl,
		Flags:      map[string]bool{},
	}

	for _, comment := range s.comments.List {
		if !strings.HasPrefix(comment.Text, SerdePrefix) {
			continue
		}
		text := strings.TrimPrefix(comment.Text, SerdePrefix)

		for _, v := range strings.Split(text, ",") {
			s.Flags[strings.Trim(v, " ")] = true
		}
	}

	return s
}

func (s serdeStruct) NeedGenerate() bool {
	return len(s.Flags) > 0
}

func (s *serdeStruct) ParseFields(state *serdeState) {
	for _, v := range s.decl.Fields.List {
		for _, name := range v.Names {
			st := parseSerdeType(v.Type)

			switch st.(type) {
			case mapType:
				state.AppendTodo(st)
			}

			s.Fields = append(s.Fields, structField{
				Name:      name.Name,
				serdeType: st,
			})
		}
	}
}

var serdeStructTmpl = template.Must(template.New("struct").Parse(`
type serdeStructEnum_{{ $.Name }} = int

const (
{{- range $idx, $field := .Fields }}
	serdeStructEnum_{{ $.Name }}_{{ $field.Name }} {{ if eq $idx 0 }} serdeStructEnum_{{ $.Name }} = iota + 1 {{ end }}
{{- end }}
)

{{ if (index $.Flags "Deserialize") }}
type {{ $.FieldVisitor }} struct {
	e serdeStructEnum_{{ $.Name }}

	serde.DummyVisitor
}

func serdeNewStructFieldVisitor_{{ $.Name }}() *{{ $.FieldVisitor }} {
	return &serdeStructFieldVisitor_{{ $.Name }}{
		DummyVisitor: serde.NewDummyVisitor("{{ $.Name }} Field"),
	}
}

func (s *{{ $.FieldVisitor }}) VisitString(v string) (err error) {
	switch v {
{{- range $idx, $field := .Fields }}
	case "{{ $field.Name }}":
		s.e = serdeStructEnum_{{ $.Name }}_{{ $field.Name }}
{{- end }}
	default:
		return errors.New("invalid field")
	}
	return nil
}

type {{ $.Visitor }} struct {
	v *{{ $.Name }}

	serde.DummyVisitor
}

func serdeNewStructVisitor_{{ $.Name }}(v *{{ $.Name }}) *{{ $.Visitor }} {
	return &{{ $.Visitor }}{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("{{ $.Name }}"),
	}
}

func (s *{{ $.Visitor }}) VisitMap(m serde.MapAccess) (err error) {
	field := serdeNewStructFieldVisitor_{{ $.Name }}()
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
{{- range $idx, $field := .Fields }}
		case serdeStructEnum_{{ $.Name }}_{{ $field.Name }}:
			v = {{ $field.NewVisitor }}(&s.v.{{ $field.Name }})
{{- end }}
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

func (s *{{ $.Name }}) Deserialize(de serde.Deserializer) (err error) {
	return de.DeserializeStruct("{{ $.Name }}", nil, {{ $.NewVisitor }}(s))
}
{{end}}

{{ if (index $.Flags "Serialize") }}
func (s *{{ $.Name }}) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("{{ $.Name }}", {{ $.Fields | len }})
	if err != nil {
		return err
	}

{{- range $idx, $field := .Fields }}
	err = st.SerializeField(
		serde.StringSerializer("{{ $field.Name }}"),
		{{ $field.Serializer }}(s.{{ $field.Name }}),
	)
	if err != nil {
		return
	}
{{- end }}
	err = st.EndStruct()
	if err != nil {
		return
	}
	return nil
}
{{end}}
`))

func (s serdeStruct) Generate() string {
	var buf bytes.Buffer

	err := serdeStructTmpl.Execute(&buf, s)
	if err != nil {
		log.Fatalf(" struct %+v execute: %v", s, err)
	}

	return buf.String()
}

func parseSerdeType(t ast.Expr) serdeType {
	switch ty := t.(type) {
	case *ast.Ident:
		if isBasicType(ty.Name) {
			return basicType(ty.Name)
		} else {
			return structType(ty.Name)
		}
	case *ast.MapType:
		return mapType{
			key:   parseSerdeType(ty.Key),
			value: parseSerdeType(ty.Value),
		}
	default:
		log.Panicf("Expr %#+v is not supported for now", ty)
		return nil
	}
}

type structField struct {
	Name string
	serdeType
}

type structType string

func (s structType) Name() string {
	return string(s)
}

func (s structType) FieldVisitor() string {
	return fmt.Sprintf("serdeStructFieldVisitor_%s", s)
}

func (s structType) Visitor() string {
	return fmt.Sprintf("serdeStructVisitor_%s", s)
}

func (s structType) NewVisitor() string {
	return fmt.Sprintf("serdeNewStructVisitor_%s", s)
}

func (s structType) Serializer() string {
	return fmt.Sprintf("%s", string(s))
}

func (s structType) Generate() string {
	panic("struct type should not be generated")
}

type basicType string

func isBasicType(v string) bool {
	switch v {
	case "bool", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"rune", "string", "byte", "bytes":
		return true
	default:
		return false
	}
}

func (bt basicType) Name() string {
	return string(bt)
}

func (bt basicType) Visitor() string {
	switch bt {
	case "bool", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"rune", "string", "byte", "bytes":
		return fmt.Sprintf("serde.%sVisitor", templateutils.ToUpperFirst(string(bt)))
	default:
		panic(fmt.Errorf("%s is not a basic type", bt))
	}
}

func (bt basicType) NewVisitor() string {
	switch bt {
	case "bool", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"rune", "string", "byte", "bytes":
		return fmt.Sprintf("serde.New%sVisitor", templateutils.ToUpperFirst(string(bt)))
	default:
		panic(fmt.Errorf("%s is not a basic type", bt))
	}
}

func (bt basicType) Serializer() string {
	switch bt {
	case "bool", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"rune", "string", "byte", "bytes":
		return fmt.Sprintf("serde.%sSerializer", templateutils.ToUpperFirst(string(bt)))
	default:
		panic(fmt.Errorf("%s is not a basic type", bt))
	}
}

func (bt basicType) Generate() string {
	panic("basic type should not be generated")
}

type mapType struct {
	key   serdeType
	value serdeType
}

func (m mapType) TypeName() string {
	return fmt.Sprintf("map[%s]%s", m.key.Name(), m.value.Name())
}

func (m mapType) Key() serdeType {
	return m.key
}

func (m mapType) Value() serdeType {
	return m.value
}

func (m mapType) Name() string {
	return fmt.Sprintf("%s_%s", m.key.Name(), m.value.Name())
}

func (m mapType) Visitor() string {
	return fmt.Sprintf("serdeMapVisitor_%s", m.Name())
}

func (m mapType) NewVisitor() string {
	return fmt.Sprintf("serdeNewMapVisitor_%s", m.Name())
}

func (m mapType) Serializer() string {
	return fmt.Sprintf("serdeSerializer_%s", m.Name())
}

var serdeMapTmpl = template.Must(template.New("map").Parse(`
type {{ $.Visitor }} struct {
	v *{{ $.TypeName }}

	serde.DummyVisitor
}

func {{ $.NewVisitor }}(v *{{ $.TypeName }}) *{{ $.Visitor }} {
	if *v == nil {
		*v = make({{ $.TypeName }})
	}
	return &{{ $.Visitor }}{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("{{ $.TypeName }}"),
	}
}

func (s *{{ $.Visitor }}) VisitMap(m serde.MapAccess) (err error) {
	var field {{ $.Key.Name }}
	var value {{ $.Value.Name }}
	for {
		ok, err := m.NextKey({{$.Key.NewVisitor}}(&field))
		if !ok {
			break
		}
		if err != nil {
			return err
		}
		err = m.NextValue({{$.Value.NewVisitor}}(&value))
		if err != nil {
			return err
		}
		(*s.v)[field] = value
	}
	return nil
}

type {{ $.Serializer }} {{ $.TypeName }}

func (s {{ $.Serializer }}) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeMap(len(s))
	if err != nil {
		return err
	}

	for k, v := range s {
		err = st.SerializeEntry(
			{{ $.Key.Serializer }}(k),
			{{ $.Value.Serializer }}(v),
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
`))

func (m mapType) Generate() string {
	var buf bytes.Buffer

	err := serdeMapTmpl.Execute(&buf, m)
	if err != nil {
		log.Fatalf(" map %+v execute: %v", m, err)
	}

	return buf.String()
}
