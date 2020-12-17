package tests

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

// serde: serialize,deserialize
type IntTypeSupport struct {
	v int
}

func TestIntTypeSupportSerialize(t *testing.T) {
	v := IntTypeSupport{
		v: 10,
	}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}

	assert.EqualValues(t, []interface{}{
		testMap(false),
		"v", 10,
		testMap(true),
	}, x)
}

func TestIntTypeSupportDeserialize(t *testing.T) {
	var v IntTypeSupport

	in := []interface{}{
		testMap(false),
		"v", 10,
		testMap(true),
	}
	err := DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize: %v", err)
	}
	assert.EqualValues(t, IntTypeSupport{
		v: 10,
	}, v)

}

// serde: serialize,deserialize
type MapTypeSupport struct {
	v map[int]int
}

func TestMapTypeSupportSerialize(t *testing.T) {
	v := MapTypeSupport{}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"v", testMap(false), testMap(true),
		testMap(true),
	}, x)
}

func TestMapTypeSupportDeserialize(t *testing.T) {
	var v MapTypeSupport

	in := []interface{}{
		testMap(false),
		"v", testMap(false), testMap(true),
		testMap(true),
	}
	err := DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize: %v", err)
	}
	assert.EqualValues(t, MapTypeSupport{
		// FIXME: v should be nil here, but serde-go is not supported for now.
		v: make(map[int]int),
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

// serde: serialize,deserialize
type MapPointerIntTypeSupport struct {
	v map[int]*IntTypeSupport
}

func TestMapPointerIntTypeSupportSerialize(t *testing.T) {
	v := MapPointerIntTypeSupport{
		v: map[int]*IntTypeSupport{
			1: {v: 10},
		},
	}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %#v", x)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"v", testMap(false), 1, testMap(false), "v", 10, testMap(true), testMap(true),
		testMap(true),
	}, x)
}

func TestMapPointerIntTypeSupportDeserialize(t *testing.T) {
	var v MapPointerIntTypeSupport

	in := []interface{}{
		testMap(false),
		"v", testMap(false), 1, testMap(false), "v", 10, testMap(true), testMap(true),
		testMap(true),
	}
	err := DeserializeFromInterfaces(in, &v)
	if err != nil {
		t.Errorf("deserialize: %v", err)
	}
	assert.EqualValues(t, MapPointerIntTypeSupport{
		v: map[int]*IntTypeSupport{
			1: {v: 10},
		},
	}, v)
}

// serde: serialize,deserialize
type ExternalTypeSupport struct {
	v sync.RWMutex
}
