// +build tools

package main

import (
	"flag"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"io/ioutil"
	"log"
	"os"
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
			content, err := ioutil.ReadFile(file)
			if err != nil {
				log.Fatalf("read file %s: %v", file, err)
			}

			f, err := parser.ParseFile(pkg.Fset, file, content, parser.ParseComments)
			if err != nil {
				log.Fatalf("parse file %s: %v", file, err)
			}

			state := newSerdeState()

			for _, v := range f.Decls {
				parse(state, v)
			}

			if !state.NeedGenerate() {
				continue
			}

			generateFile, err := os.OpenFile(
				fmt.Sprintf(formatGeneratedFilename(file)),
				os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
			if err != nil {
				log.Fatalf("open file: %v", err)
			}

			err = packageTmpl.Execute(generateFile, pkg.Name)
			if err != nil {
				log.Fatalf("pacakge tmpl execute: %v", err)
			}

			for _, v := range state.todo {
				_, err = generateFile.WriteString(v.Generate())
				if err != nil {
					log.Fatalf("generate: %v", err)
				}
			}
		}
	}
}

func parse(state *serdeState, decl ast.Decl) {
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
			return false
		}
		st, ok := typeSpec.Type.(*ast.StructType)
		if !ok {
			return true
		}

		s := newSerdeStruct(typeSpec.Name.Name, decl.Doc, st)

		if !s.NeedGenerate() {
			return true
		}

		s.ParseFields(state)
		state.todo = append(state.todo, s)
		return true
	})
}

var packageTmpl = template.Must(template.New("package").Parse(`
package {{ . }}

import (
	"errors"

	"github.com/Xuanwo/serde-go"
)

var _ = errors.New
`))
