// +build tools

package main

import (
	"go/parser"
	"go/token"
	"io/ioutil"
	"testing"
)

func TestStructType(t *testing.T) {
	testContent, _ := ioutil.ReadFile("testdata/test.go")

	state := newSerdeState()
	parseStructs(t, state, string(testContent))

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
