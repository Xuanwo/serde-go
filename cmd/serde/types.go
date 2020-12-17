// +build tools

package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"log"
	"strconv"
	"text/template"

	"github.com/Xuanwo/templateutils"
)

type serdeType interface {
	// Name is the display name, follows serde name style.
	Name() string
	// TypeName is the type name.
	TypeName() string

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

func newSerdeState(generated map[string]struct{}) *serdeState {
	return &serdeState{
		todo:      make([]serdeType, 0),
		generated: generated,
	}
}

type serdeStruct struct {
	structType

	Fields []structField
	Flags  map[string]string

	decl *ast.StructType
}

func newSerdeStruct(name string, comments *ast.CommentGroup, decl *ast.StructType) serdeStruct {
	return serdeStruct{
		structType: structType(name),
		decl:       decl,
		Flags:      parseTagsFromStructComments(comments),
	}
}

func (s serdeStruct) IsSerialize() bool {
	_, hasSerialize := s.Flags[TagSerialize]
	return hasSerialize
}

func (s serdeStruct) IsDeserialize() bool {
	_, hasDeserialize := s.Flags[TagDeserialize]
	return hasDeserialize
}

func (s serdeStruct) NeedGenerate() bool {
	_, hasDeserialize := s.Flags[TagDeserialize]
	_, hasSerialize := s.Flags[TagSerialize]
	return hasDeserialize || hasSerialize
}

func (s *serdeStruct) ParseFields(state *serdeState) {
	for _, v := range s.decl.Fields.List {
		for _, name := range v.Names {
			st := parseSerdeType(state, v.Type)

			s.Fields = append(s.Fields, structField{
				Name:      name.Name,
				Flags:     parseTagsFromStructTag(v.Tag),
				serdeType: st,
			})
		}
	}
}

var serdeStructTmpl = template.Must(template.New("struct").Funcs(templateutils.FuncMap()).Parse(`
type serdeStructEnum_{{ $.Name }} = int

const (
	serdeStructEnumSkip_{{ $.Name }} serdeStructEnum_{{ $.Name }} = 0
{{- range $idx, $field := .Fields }}
	{{ if not (and $field.IsSkipDeserialize $field.IsSkipSerialize) }}
	serdeStructEnum_{{ $.Name }}_{{ $field.Name }}  serdeStructEnum_{{ $.Name }} = {{ add $idx 1 }}
	{{ end }}
{{- end }}
)

{{ if $.IsDeserialize }}
type {{ $.FieldVisitor }} struct {
	e *serdeStructEnum_{{ $.Name }}

	serde.DummyVisitor
}

func serdeNewStructFieldVisitor_{{ $.Name }}(e *serdeStructEnum_{{ $.Name }}) {{ $.FieldVisitor }} {
	return serdeStructFieldVisitor_{{ $.Name }}{
		e: e,
		DummyVisitor: serde.NewDummyVisitor("{{ $.Name }} Field"),
	}
}

func (s {{ $.FieldVisitor }}) VisitString(v string) (err error) {
	switch v {
{{- range $idx, $field := .Fields }}
	{{ if not $field.IsSkipDeserialize }}
	case "{{ $field.Name }}":
		*s.e = serdeStructEnum_{{ $.Name }}_{{ $field.Name }}
	{{ else }}
	case "{{ $field.Name }}":
		*s.e = serdeStructEnumSkip_{{ $.Name }}
	{{ end }}
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

func serdeNewStructVisitor_{{ $.Name }}(v *{{ $.Name }}) {{ $.Visitor }} {
	return {{ $.Visitor }}{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("{{ $.Name }}"),
	}
}

func (s {{ $.Visitor }}) VisitMap(m serde.MapAccess) (err error) {
	var fieldValue serdeStructEnum_{{ $.Name }}
	field := serdeNewStructFieldVisitor_{{ $.Name }}(&fieldValue)
	for {
		ok, err := m.NextKey(field)
		if !ok {
			break
		}
		if err != nil {
			return err
		}

		var v serde.Visitor
		switch *field.e {
		case serdeStructEnumSkip_{{ $.Name }}:
			v = serde.SkipVisitor{}
{{- range $idx, $field := .Fields }}
	{{ if not $field.IsSkipDeserialize }}
		case serdeStructEnum_{{ $.Name }}_{{ $field.Name }}:
			v = {{ $field.NewVisitor }}(&s.v.{{ $field.Name }})
	{{ end }}
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

{{ if $.IsSerialize }}
func (s {{ $.Name }}) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeStruct("{{ $.Name }}", {{ $.Fields | len }})
	if err != nil {
		return err
	}

{{- range $idx, $field := .Fields }}
	{{ if not $field.IsSkipSerialize }}
	err = st.SerializeField(
		serde.StringSerializer("{{ $field.Name }}"),
		{{ $field.Serializer }}(s.{{ $field.Name }}),
	)
	if err != nil {
		return
	}
	{{ end }}
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

func parseSerdeType(state *serdeState, t ast.Expr) (st serdeType) {
	switch ty := t.(type) {
	case *ast.Ident:
		if isBasicType(ty.Name) {
			return basicType(ty.Name)
		} else {
			return structType(ty.Name)
		}
	case *ast.MapType:
		st = mapType{
			key:   parseSerdeType(state, ty.Key),
			value: parseSerdeType(state, ty.Value),
		}
		state.AppendTodo(st)
		return st
	case *ast.ArrayType:
		st := sliceType{
			length:  0,
			element: parseSerdeType(state, ty.Elt),
		}

		if l, ok := ty.Len.(*ast.BasicLit); ok {
			stl, err := strconv.ParseInt(l.Value, 10, 64)
			if err != nil {
				panic(fmt.Errorf("invalid array length: %v", l.Value))
			}

			st.length = int(stl)
		}
		state.AppendTodo(st)
		return st
	case *ast.StarExpr:
		st = pointerType{internal: parseSerdeType(state, ty.X)}
		state.AppendTodo(st)
		return st
	case *ast.FuncType, *ast.ChanType:
		// Ignore golang runtime types.
		return nil
	default:
		log.Panicf("Expr %#+v is not supported for now", ty)
		return nil
	}
}

type structField struct {
	Name  string
	Flags map[string]string
	serdeType
}

func (s structField) IsSkipSerialize() bool {
	_, ok := s.Flags[TagSkipSerialize]
	if ok {
		return true
	}
	_, ok = s.Flags[TagSkip]
	return ok
}

func (s structField) IsSkipDeserialize() bool {
	_, ok := s.Flags[TagSkipDeserialize]
	if ok {
		return true
	}
	_, ok = s.Flags[TagSkip]
	return ok
}

type structType string

func (s structType) Name() string {
	return string(s)
}

func (s structType) TypeName() string {
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

func (bt basicType) TypeName() string {
	return bt.Name()
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

func (m mapType) Key() serdeType {
	return m.key
}

func (m mapType) Value() serdeType {
	return m.value
}

func (m mapType) Name() string {
	return fmt.Sprintf("%s_%s", m.key.Name(), m.value.Name())
}

func (m mapType) TypeName() string {
	return fmt.Sprintf("map[%s]%s", m.key.TypeName(), m.value.TypeName())
}

func (m mapType) Visitor() string {
	return fmt.Sprintf("serdeMapVisitor_%s", m.Name())
}

func (m mapType) NewVisitor() string {
	return fmt.Sprintf("serdeNewMapVisitor_%s", m.Name())
}

func (m mapType) Serializer() string {
	return fmt.Sprintf("serdeMapSerializer_%s", m.Name())
}

var serdeMapTmpl = template.Must(template.New("map").Parse(`
type {{ $.Visitor }} struct {
	v *{{ $.TypeName }}

	serde.DummyVisitor
}

func {{ $.NewVisitor }}(v *{{ $.TypeName }}) {{ $.Visitor }} {
	if *v == nil {
		*v = make({{ $.TypeName }})
	}
	return {{ $.Visitor }}{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("{{ $.TypeName }}"),
	}
}

func (s {{ $.Visitor }}) VisitMap(m serde.MapAccess) (err error) {
	var field {{ $.Key.TypeName }}
	var value {{ $.Value.TypeName }}
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
		log.Fatalf("map %+v execute: %v", m, err)
	}

	return buf.String()
}

type sliceType struct {
	length  int
	element serdeType
}

func (s sliceType) Name() string {
	if s.length == 0 {
		return fmt.Sprintf("%s", s.element.Name())
	}
	return fmt.Sprintf("%d_%s", s.length, s.element.Name())
}

func (s sliceType) TypeName() string {
	if s.length != 0 {
		return fmt.Sprintf("[%d]%s", s.length, s.element.TypeName())
	}
	return fmt.Sprintf("[]%s", s.element.TypeName())
}

func (s sliceType) Length() int {
	return s.length
}

func (s sliceType) Element() serdeType {
	return s.element
}

func (s sliceType) Visitor() string {
	return fmt.Sprintf("serdeSliceVisitor_%s", s.Name())
}

func (s sliceType) NewVisitor() string {
	return fmt.Sprintf("serdeNewSliceVisitor_%s", s.Name())
}

func (s sliceType) Serializer() string {
	return fmt.Sprintf("serdeSliceSerializer_%s", s.Name())
}

var serdeSliceTmpl = template.Must(template.New("slice").Parse(`
type {{ $.Visitor }} struct {
	v *{{ $.TypeName }}

	serde.DummyVisitor
}

func {{ $.NewVisitor }}(v *{{ $.TypeName }}) {{ $.Visitor }} {
	return {{ $.Visitor }}{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("{{ $.TypeName }}"),
	}
}

func (s {{ $.Visitor }}) VisitSlice(m serde.SliceAccess) (err error) {
	var value {{ $.Element.Name }}

	{{- if ne $.Length 0 }}
	i := 0
	{{- end }}
	for {
		ok, err := m.NextElement({{$.Element.NewVisitor}}(&value))
		if !ok {
			break
		}
		if err != nil {
			return err
		}
		{{ if eq $.Length 0 }}
		*s.v = append(*s.v, value)
		{{ else }}
		(*s.v)[i] = value
		i += 1
		{{ end }}
	}
	return nil
}

type {{ $.Serializer }} {{ $.TypeName }}

func (s {{ $.Serializer }}) Serialize(ser serde.Serializer) (err error) {
	st, err := ser.SerializeSlice(len(s))
	if err != nil {
		return err
	}

	for _, v := range s {
		err = st.SerializeElement({{ $.Element.Serializer }}(v))
		if err != nil {
			return
		}
	}

	err = st.EndSlice()
	if err != nil {
		return
	}
	return nil
}
`))

func (s sliceType) Generate() string {
	var buf bytes.Buffer

	err := serdeSliceTmpl.Execute(&buf, s)
	if err != nil {
		log.Fatalf("slice %+v execute: %v", s, err)
	}

	return buf.String()
}

type pointerType struct {
	internal serdeType
}

func (p pointerType) Name() string {
	// Add prefix "_" to represents pointer "*"
	return fmt.Sprintf("_%s", p.internal.Name())
}

func (p pointerType) TypeName() string {
	return fmt.Sprintf("*%s", p.internal.TypeName())
}

func (p pointerType) Internal() serdeType {
	return p.internal
}

func (p pointerType) Visitor() string {
	return fmt.Sprintf("serdePointerVisitor_%s", p.Name())
}

func (p pointerType) NewVisitor() string {
	return fmt.Sprintf("serdeNewPointerVisitor_%s", p.Name())
}

func (p pointerType) Serializer() string {
	return fmt.Sprintf("serdePointerSerializer_%s", p.Name())
}

var serdePointerTmpl = template.Must(template.New("pointer").Parse(`
type {{ $.Visitor }} struct {
	{{ $.Internal.Visitor }}
}

func {{ $.NewVisitor }}(v *{{ $.TypeName }}) {{ $.Visitor }} {
	// FIXME: nil is not handled correctly
	var tv {{ $.Internal.TypeName }}
	*v = &tv
	return {{ $.Visitor }}{ {{ $.Internal.NewVisitor }}(*v) }
}

func {{ $.Serializer }}(v {{ $.TypeName }}) serde.Serializable {
	if v == nil {
		return serde.NilSerializer{}
	}
	return {{ $.Internal.Serializer }}(*v)
}
`))

func (p pointerType) Generate() string {
	var buf bytes.Buffer

	err := serdePointerTmpl.Execute(&buf, p)
	if err != nil {
		log.Fatalf("pointer %+v execute: %v", p, err)
	}

	return buf.String()
}
