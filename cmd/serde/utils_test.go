// +build tools

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseTags(t *testing.T) {
	cases := []struct {
		name     string
		text     string
		expected map[string]string
	}{
		{
			"simple",
			"a",
			map[string]string{
				"a": "",
			},
		},
		{
			"complex",
			"a,b=c",
			map[string]string{
				"a": "",
				"b": "c",
			},
		},
		{
			"has space",
			"a , b = c",
			map[string]string{
				"a": "",
				"b": "c",
			},
		},
	}

	for _, v := range cases {
		t.Run(v.name, func(t *testing.T) {
			m := make(map[string]string)

			parseTags(m, v.text)

			assert.EqualValues(t, v.expected, m)
		})
	}
}
