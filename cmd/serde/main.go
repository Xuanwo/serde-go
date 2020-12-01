// +build tools

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"golang.org/x/tools/go/packages"
)

func main() {
	flag.Parse()

	cfg := &packages.Config{
		Mode:  packages.NeedName | packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Tests: true,
	}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		log.Fatalf("load: %v\n", err)
	}
	// Errors could be stored in pkg.Errors, but we can ignore them fow now.

	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			log.Printf("read file: %s", file)
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf("read file %s: %v", file, err)
			}

			f, err := parser.ParseFile(pkg.Fset, file, content, parser.ParseComments)
			if err != nil {
				log.Fatalf("parse file %s: %v", file, err)
			}

			var sts []structType

			for _, v := range f.Decls {
				sts = append(sts, parse(v)...)
			}

			var generateFile io.Writer

			if len(sts) == 0 {
				continue
			}

			generateFile, err = os.OpenFile(
				fmt.Sprintf(formatGeneratedFilename(file)),
				os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("open file: %v", err)
			}

			err = packageTmpl.Execute(generateFile, pkg.Name)
			if err != nil {
				log.Fatalf("pacakge tmpl execute: %v", err)
			}

			for _, v := range sts {
				generate(generateFile, v)
			}
		}
	}
}

type structType struct {
	Name   string
	Fields []structFiled
	Flags  map[string]bool

	comments *ast.CommentGroup
	decl     *ast.StructType
}

func (s *structType) Parse() error {
	// Parse struct Flags
	s.Flags = map[string]bool{}
	for _, comment := range s.comments.List {
		if !strings.HasPrefix(comment.Text, SerdePrefix) {
			continue
		}
		text := strings.TrimPrefix(comment.Text, SerdePrefix)

		for _, v := range strings.Split(text, ",") {
			s.Flags[strings.Trim(v, " ")] = true
		}
	}

	// Parse fields.
	for _, v := range s.decl.Fields.List {
		for _, name := range v.Names {
			f := structFiled{Name: name.Name}

			switch ty := v.Type.(type) {
			case *ast.Ident:
				f.Type = ty.Name
			default:
				log.Printf("struct %s field %s %+v is not supported for now", s.Name, f.Name, ty)
				continue
			}

			s.Fields = append(s.Fields, f)
		}
	}

	return nil
}

type structFiled struct {
	Name string
	Type string
}

func parse(decl ast.Decl) []structType {
	var sts []structType

	ast.Inspect(decl, func(node ast.Node) bool {
		decl, ok := node.(*ast.GenDecl)
		if !ok {
			return true
		}
		// Only support type here.
		if decl.Tok != token.TYPE {
			return true
		}
		// Ignore types without comments.
		if decl.Doc == nil {
			return true
		}
		// Only handle struct type.
		typeSpec, ok := decl.Specs[0].(*ast.TypeSpec)
		if !ok {
			log.Printf("%+v is not supported for now", decl.Specs[0])
			return false
		}
		st, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		s := structType{
			Name:     typeSpec.Name.Name,
			Fields:   nil,
			comments: decl.Doc,
			decl:     st,
		}
		err := s.Parse()
		if err != nil {
			log.Fatalf("struct %v parse: %v", s.Name, err)
		}

		if len(s.Flags) > 0 {
			sts = append(sts, s)
		}
		return true
	})

	return sts
}

const SerdePrefix = "// serde:"

func generate(w io.Writer, st structType) {
	if st.Flags["Serialize"] {
		log.Printf("generate serialize for %s", st.Name)
		generateSerialize()
	}

	if st.Flags["Deserialize"] {
		log.Printf("generate deserialize for %s", st.Name)
		generateDeserialize(w, st)
	}
}

func generateSerialize() {}

func (f *structFiled) Visitor() string {
	switch f.Type {
	case "bool", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"rune", "string", "byte", "bytes":
		return fmt.Sprintf("serde.New%sVisitor", string(f.Type[0]-'a'+'A')+f.Type[1:])
	default:
		return fmt.Sprintf("new%sVisitor", f.Type)
	}
}

func (f *structFiled) Serializer() string {
	switch f.Type {
	case "bool", "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"float32", "float64", "complex64", "complex128",
		"rune", "string", "byte", "bytes":
		return fmt.Sprintf("serde.%sSerializer", string(f.Type[0]-'a'+'A')+f.Type[1:])
	default:
		return fmt.Sprintf("new%sSerializer", f.Type)
	}
}

func generateDeserialize(w io.Writer, st structType) {
	err := deTmpl.Execute(w, st)
	if err != nil {
		log.Fatalf("de tmpl execute: %v", err)
	}
}

var packageTmpl = template.Must(template.New("package").Parse(`
package {{ . }}

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)
`))

var deTmpl = template.Must(template.New("de").Parse(`
type serde{{ $.Name }}Enum = int

const (
{{- range $idx, $field := .Fields }}
	serde{{ $.Name }}Enum{{ $field.Name }} {{ if eq $idx 0 }} serde{{ $.Name }}Enum = iota + 1 {{ end }}
{{- end }}
)

type serde{{ $.Name }}FieldVisitor struct {
	e serde{{ $.Name }}Enum

	serde.DummyVisitor
}

func newSerde{{ $.Name }}FieldVisitor() *serde{{ $.Name }}FieldVisitor {
	return &serde{{ $.Name }}FieldVisitor{
		DummyVisitor: serde.NewDummyVisitor("{{ $.Name }} Field"),
	}
}

func (s *serde{{ $.Name }}FieldVisitor) VisitString(v string) (err error) {
	switch v {
{{- range $idx, $field := .Fields }}
	case "{{ $field.Name }}":
		s.e = serde{{ $.Name }}Enum{{ $field.Name }}
{{- end }}
	default:
		return errors.New("invalid field")
	}
	return nil
}

type serde{{ $.Name }}Visitor struct {
	v *{{ $.Name }}

	serde.DummyVisitor
}

func new{{ $.Name }}Visitor(v *{{ $.Name }}) *serde{{ $.Name }}Visitor {
	return &serde{{ $.Name }}Visitor{
		v: v,
		DummyVisitor: serde.NewDummyVisitor("{{ $.Name }}"),
	}
}

func (s *serde{{ $.Name }}Visitor) VisitMap(m serde.MapAccess) (err error) {
	field := newSerde{{ $.Name }}FieldVisitor()
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
		case serde{{ $.Name }}Enum{{ $field.Name }}:
			v = {{ $field.Visitor }}(&s.v.{{ $field.Name }})
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
	return de.DeserializeStruct("{{ $.Name }}", nil, new{{ $.Name }}Visitor(s))
}

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
`))
