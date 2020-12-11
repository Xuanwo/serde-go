// +build tools

package main

import (
	"go/parser"
	"go/token"
	"testing"
)

const testContent = `
package test

// serde: Deserialize,Serialize
type Test struct {
	vint64 int64
	vmap map[int]int
	varray [2]int
	vslice []int
}
`

func TestStructType(t *testing.T) {
	state := newSerdeState()
	parseStructs(t, state, testContent)

	for _, v := range state.todo {
		t.Logf("%s", v.Generate())
	}
}

func parseStructs(t *testing.T, state *serdeState, content string) {
	f, err := parser.ParseFile(token.NewFileSet(), "test.go", content, parser.ParseComments)
	if err != nil {
		t.Errorf("parse file: %v", err)
		return
	}

	for _, v := range f.Decls {
		parse(state, v)
	}
}
