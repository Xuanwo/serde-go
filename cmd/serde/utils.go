package main

import (
	"strings"
)

func formatGeneratedFilename(s string) string {
	if strings.HasSuffix(s, "_test.go") {
		return strings.ReplaceAll(s, "_test.go", "_serde_generated_test.go")
	}
	return strings.ReplaceAll(s, ".go", "_serde_generated.go")
}
