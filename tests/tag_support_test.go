package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// serde: serialize,deserialize
type TagSupport struct {
	skip            int `serde:"skip"`
	skipSerialize   int `serde:"skip_serialize"`
	skipDeserialize int `serde:"skip_deserialize"`
}

func TestTagSupportSerialize(t *testing.T) {
	v := TagSupport{
		skip:            1,
		skipSerialize:   2,
		skipDeserialize: 3,
	}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize %v: %v", x, err)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"skipDeserialize", 3,
		testMap(true),
	}, x)
}

func TestTagSupportDeserialize(t *testing.T) {
	var v TagSupport

	in := []interface{}{
		testMap(false),
		"skipSerialize", 2,
		"skipDeserialize", 3,
		testMap(true),
	}

	err := DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize %v: %v", in, err)
	}
	assert.EqualValues(t, TagSupport{
		skipSerialize: 2,
	}, v)
}
