package main

import (
	"bytes"
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
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
		Mode:  packages.NeedFiles | packages.NeedSyntax | packages.NeedTypes | packages.NeedTypesInfo,
		Tests: true,
	}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		log.Printf("load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf("read file %s: %v", file, err)
			}

			f, err := parser.ParseFile(pkg.Fset, file, content, parser.ParseComments)
			if err != nil {
				log.Fatalf("parse file %s: %v", file, err)
			}

			for _, v := range f.Decls {
				parse(v)
			}
		}
	}
}

func parse(decl ast.Decl) {
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
		typeSpec := decl.Specs[0].(*ast.TypeSpec)
		st, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		generate(typeSpec.Name.Name, decl.Doc, st)
		return true
	})
}

const SerdePrefix = "// serde:"

func generate(name string, comments *ast.CommentGroup, decl *ast.StructType) {
	var structFlag = map[string]bool{
		"Serialize":   false,
		"Deserialize": false,
	}
	for _, comment := range comments.List {
		if !strings.HasPrefix(comment.Text, SerdePrefix) {
			continue
		}
		text := strings.TrimPrefix(comment.Text, SerdePrefix)

		for _, v := range strings.Split(text, ",") {
			v = strings.Trim(v, " ")
			if _, ok := structFlag[v]; ok {
				structFlag[v] = true
			}
		}
	}

	if structFlag["Serialize"] {
		generateSerialize()
	}

	if structFlag["Deserialize"] {
		generateDeserialize(name, decl.Fields.List)
	}
}

func generateSerialize() {}

type structType struct {
	Name   string
	Fields []structFiled
}

type structFiled struct {
	Name string
	Type string
}

func (f *structFiled) Visitor() string {
	switch f.Type {
	case "int", "int8", "int16", "int32", "int64",
		"uint", "uint8", "uint16", "uint32", "uint64",
		"string":
		return fmt.Sprintf("serde.New%sVisitor", string(f.Type[0]-'a'+'A')+f.Type[1:])
	default:
		return fmt.Sprintf("new%sVisitor", f.Type)
	}
}

func generateDeserialize(name string, fields []*ast.Field) {
	var content bytes.Buffer

	st := structType{Name: name}

	for _, v := range fields {
		for _, name := range v.Names {
			st.Fields = append(st.Fields, structFiled{
				Name: name.Name,
				Type: v.Type.(*ast.Ident).Name,
			})
		}
	}

	// Generate struct enums
	err := deTmpl.Execute(&content, st)
	if err != nil {
		log.Fatalf("de tmpl execute: %v", err)
	}

	// Generate struct filed visitor
	// Generate struct value visitor

	log.Print(content.String())
}

var deTmpl = template.Must(template.New("de").Parse(`
type serde{{ $.Name }}Enum = int

const (
{{- range $idx, $field := .Fields }}
	serde{{ $.Name }}Enum{{ $field.Name }} {{ if eq $idx 0 }} serde{{ $.Name }}Enum = itoa + 1 {{ end }}
{{- end }}
)

type serde{{ $.Name }}FieldVisitor struct {
	e serde{{ $.Name }}Enum

	serde.DummyVisitor
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
	return &serde{{ $.Name }}Visitor{v: v}
}

func (s *serde{{ $.Name }}Visitor) VisitMap(m serde.MapAccess) (err error) {
	field := &serde{{ $.Name }}FieldVisitor{}
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
`))
