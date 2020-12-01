// +build tools

package main

import (
	"go/parser"
	"go/token"
	"log"
	"testing"
)

const content = `
package test

// serde: Deserialize
type Test struct {
	A,D string 
	B string
	C int32
	Value int64
	q uint8
}
`

func TestParse(t *testing.T) {
	f, err := parser.ParseFile(token.NewFileSet(), "test.go", content, parser.ParseComments)
	if err != nil {
		log.Fatalf("parse file: %v", err)
	}

	for _, v := range f.Decls {
		parse(v)
	}
}
