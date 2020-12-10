// +build tools

package main

import (
	"go/parser"
	"go/token"
	"testing"
)

const basicTypeContent = `
package test

// serde: Deserialize
type Test struct {
	X int64
	x map[int]int
}
`

func TestStructType(t *testing.T) {
	state := newSerdeState()
	parseStructs(t, state, basicTypeContent)

	for _, v := range state.todo {
		t.Logf("%s", v.Generate())
	}
}

func parseStructs(t *testing.T, state *serdeState, content string) {
	f, err := parser.ParseFile(token.NewFileSet(), "test.go", content, parser.ParseComments)
	if err != nil {
		t.Errorf("parse file: %v", err)
	}

	for _, v := range f.Decls {
		parse(state, v)
	}
}
