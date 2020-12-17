package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// serde: serialize,deserialize
type TypeSupport struct {
	vint int
	vmap map[string]string
}

func TestTypeSupportSerialize(t *testing.T) {
	v := TypeSupport{}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"vint", 0,
		"vmap", testMap(false), testMap(true),
		testMap(true),
	}, x)

	v = TypeSupport{
		vint: 10,
		vmap: map[string]string{
			"map_key_a": "map_value_a",
		},
	}

	x, err = SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}

	assert.EqualValues(t, []interface{}{
		testMap(false),
		"vint", 10,
		"vmap", testMap(false), "map_key_a", "map_value_a", testMap(true),
		testMap(true),
	}, x)
}

func TestTypeSupportDeserialize(t *testing.T) {
	var v TypeSupport

	in := []interface{}{
		testMap(false),
		"vint", 0,
		"vmap", testMap(false), testMap(true),
		testMap(true),
	}
	err := DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize: %v", err)
	}
	assert.EqualValues(t, TypeSupport{
		// FIXME: vmap should be nil here, but serde-go is not supported for now.
		vmap: make(map[string]string),
	}, v)

	in = []interface{}{
		testMap(false),
		"vint", 10,
		"vmap", testMap(false), "map_key_a", "map_value_a", testMap(true),
		testMap(true),
	}
	err = DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize: %v", err)
	}
	assert.EqualValues(t, TypeSupport{
		vint: 10,
		vmap: map[string]string{
			"map_key_a": "map_value_a",
		},
	}, v)

}

// serde: serialize,deserialize
type MapPointerTypeSupport struct {
	v map[int]*int
}

func TestMapPointerTypeSupportSerialize(t *testing.T) {
	vint := 10
	v := MapPointerTypeSupport{
		v: map[int]*int{
			1: &vint,
		},
	}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"v", testMap(false), 1, 10, testMap(true),
		testMap(true),
	}, x)
}

func TestMapPointerTypeSupportDeserialize(t *testing.T) {
	var v MapPointerTypeSupport

	in := []interface{}{
		testMap(false),
		"v", testMap(false), 1, 10, testMap(true),
		testMap(true),
	}
	err := DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize: %v", err)
	}
	vint := 10
	assert.EqualValues(t, MapPointerTypeSupport{
		v: map[int]*int{
			1: &vint,
		},
	}, v)
}
