// +build tools

package main

import (
	"go/ast"
	"log"
	"strings"
)

func formatGeneratedFilename(s string) string {
	if strings.HasSuffix(s, "_test.go") {
		return strings.ReplaceAll(s, "_test.go", "_serde_generated_test.go")
	}
	return strings.ReplaceAll(s, ".go", "_serde_generated.go")
}

func parseTags(flags map[string]string, s string) {
	// All space in tags could be removed safely.
	s = strings.ReplaceAll(s, " ", "")

	for _, v := range strings.Split(s, ",") {
		vs := strings.Split(v, "=")

		if len(vs) > 0 {
			// Ignore all empty key
			if vs[0] == "" {
				continue
			}
		}

		switch len(vs) {
		case 1:
			flags[vs[0]] = ""
		case 2:
			flags[vs[0]] = vs[1]
		default:
			log.Fatalf("invalid tag format: %s", v)
		}
	}
}

func parseTagsFromStructTag(tag *ast.BasicLit) map[string]string {
	m := make(map[string]string)

	if tag == nil {
		return m
	}

	// `serde:"abc"` => serde:"abc"
	content := strings.Trim(tag.Value, "`")
	// serde:"abc" => "abc"
	content = strings.TrimPrefix(content, "serde:")
	// "abc" => abc
	content = strings.Trim(content, "\"")

	parseTags(m, content)
	return m
}

func parseTagsFromStructComments(comments *ast.CommentGroup) map[string]string {
	serdePrefix := "// serde:"

	m := make(map[string]string)

	if comments == nil {
		return m
	}

	for _, comment := range comments.List {
		if !strings.HasPrefix(comment.Text, serdePrefix) {
			continue
		}
		// "// serde: abc" => "abc"
		text := strings.TrimPrefix(comment.Text, serdePrefix)

		parseTags(m, text)
	}
	return m
}
