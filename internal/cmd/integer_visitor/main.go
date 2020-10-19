package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

type Integer struct {
	name string
	Size int
}

func (i *Integer) Type() string {
	return i.name
}

func (i *Integer) Name() string {
	return string(i.name[0]-'a'+'A') + i.name[1:]
}

func (i *Integer) IsUint() bool {
	return strings.HasPrefix(i.name, "u")
}

var Integers = []Integer{
	{"int8", 1},
	{"uint8", 2},
	{"int16", 3},
	{"uint16", 4},
	{"int32", 5},
	{"int", 6},
	{"uint32", 7},
	{"uint", 8},
	{"int64", 9},
	{"uint64", 10},
}

func main() {
	f, err := os.OpenFile("de_primitive_integer.go", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0600)
	if err != nil {
		log.Fatalf("open file: %v", err)
	}

	err = tmpl.Execute(f, Integers)
	if err != nil {
		log.Fatalf("execute template: %v", err)
	}
}

var tmpl = template.Must(template.New("integer_visitor").Parse(`
package serde

import (
	"errors"
)

{{ range $_, $v := . }}
{{ $type := $v.Type }}
{{ $name := $v.Name }}
type {{$name}}Visitor struct {
	v *{{$type}}
}

func New{{$name}}Visitor(v *{{$type}}) {{$name}}Visitor {
	return {{$name}}Visitor{v: v}
}

func (vi {{$name}}Visitor) String() string {
	return "{{$name}}"
}

func (vi {{$name}}Visitor) VisitInt8(v int8) (err error) {
	{{- if $v.IsUint }}
	if v < Min{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitUint8(v uint8) (err error) {
	{{- if lt $v.Size 2 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitInt16(v int16) (err error) {
	{{- if lt $v.Size 3 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	{{- if or (lt $v.Size 3) ($v.IsUint) }}
	if v < Min{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitUint16(v uint16) (err error) {
	{{- if lt $v.Size 4 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitInt32(v int32) (err error) {
	{{- if lt $v.Size 5 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	{{- if or (lt $v.Size 5) ($v.IsUint) }}
	if v < Min{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitInt(v int) (err error) {
	{{- if lt $v.Size 6 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	{{- if or (lt $v.Size 6) ($v.IsUint) }}
	if v < Min{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitUint32(v uint32) (err error) {
	{{- if eq $v.Size 6 }}
	if UintSize == 32 && v > MaxInt32 {
		return errors.New("overflow")
	}
	{{- else if lt $v.Size 7 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitUint(v uint) (err error) {
	{{- if lt $v.Size 8 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitInt64(v int64) (err error) {
	{{- if eq $v.Size 8 }}
	if UintSize == 32 && v > MaxUint32 {
		return errors.New("overflow")
	}
	{{- else if lt $v.Size 9 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	{{- if or (lt $v.Size 9) ($v.IsUint) }}
	if v < Min{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}

func (vi {{$name}}Visitor) VisitUint64(v uint64) (err error) {
	{{- if lt $v.Size 10 }}
	if v > Max{{$name}} {
		return errors.New("overflow")
	}
	{{- end }}
	*vi.v = {{$type}}(v)
	return nil
}
{{ end }}
`))
