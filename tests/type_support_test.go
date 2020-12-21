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
type PointerIntTypeSupport struct {
	v *int
}

func TestPointerIntTypeSupportSerialize(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		v := PointerIntTypeSupport{}

		x, err := SerializeToInterfaces(&v)
		if err != nil {
			t.Errorf("serialize: %v", x)
		}

		assert.EqualValues(t, []interface{}{
			testMap(false),
			"v", nil,
			testMap(true),
		}, x)
	})

	t.Run("valid value", func(t *testing.T) {
		tv := 10
		v := PointerIntTypeSupport{
			v: &tv,
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
	})
}

func TestPointerIntTypeSupportDeserialize(t *testing.T) {
	t.Run("nil value", func(t *testing.T) {
		var v PointerIntTypeSupport

		in := []interface{}{
			testMap(false),
			"v", nil,
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, PointerIntTypeSupport{}, v)
	})

	t.Run("valid value", func(t *testing.T) {
		var v PointerIntTypeSupport

		in := []interface{}{
			testMap(false),
			"v", 10,
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		tv := 10
		assert.EqualValues(t, PointerIntTypeSupport{
			v: &tv,
		}, v)
	})
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

	t.Run("empty map", func(t *testing.T) {
		in := []interface{}{
			testMap(false),
			"v", testMap(false), testMap(true),
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, MapTypeSupport{}, v)
	})

	t.Run("non empty map", func(t *testing.T) {
		in := []interface{}{
			testMap(false),
			"v", testMap(false), 1, 2, testMap(true),
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, MapTypeSupport{
			v: map[int]int{1: 2},
		}, v)
	})
}

// serde: serialize,deserialize
type SliceTypeSupport struct {
	v []int
}

func TestSliceTypeSupportSerialize(t *testing.T) {
	v := SliceTypeSupport{}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"v", testSlice(false), testSlice(true),
		testMap(true),
	}, x)
}

func TestSliceTypeSupportDeserialize(t *testing.T) {
	var v SliceTypeSupport

	t.Run("empty slice", func(t *testing.T) {
		in := []interface{}{
			testMap(false),
			"v", testSlice(false), testSlice(true),
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, SliceTypeSupport{}, v)
	})

	t.Run("non empty slice", func(t *testing.T) {
		in := []interface{}{
			testMap(false),
			"v", testSlice(false), 1, 2, testSlice(true),
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, SliceTypeSupport{
			v: []int{1, 2},
		}, v)
	})
}

// serde: serialize,deserialize
type ArrayTypeSupport struct {
	v [2]int
}

func TestArrayTypeSupportSerialize(t *testing.T) {
	v := ArrayTypeSupport{}

	x, err := SerializeToInterfaces(&v)
	if err != nil {
		t.Errorf("serialize: %v", x)
	}
	assert.EqualValues(t, []interface{}{
		testMap(false),
		"v", testSlice(false), 0, 0, testSlice(true),
		testMap(true),
	}, x)
}

func TestArrayTypeSupportDeserialize(t *testing.T) {
	var v ArrayTypeSupport

	t.Run("empty array", func(t *testing.T) {
		in := []interface{}{
			testMap(false),
			"v", testSlice(false), 0, 0, testSlice(true),
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, ArrayTypeSupport{}, v)
	})

	t.Run("non empty array", func(t *testing.T) {
		in := []interface{}{
			testMap(false),
			"v", testSlice(false), 1, 2, testSlice(true),
			testMap(true),
		}
		err := DeserializeFromInterfaces(in, &v)
		if err != nil {
			t.Errorf("deserialize: %v", err)
		}
		assert.EqualValues(t, ArrayTypeSupport{
			v: [2]int{1, 2},
		}, v)
	})
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
