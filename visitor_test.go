package serde

import (
	"testing"
)

func TestNewIntVisitor(t *testing.T) {
	x := 10
	vv := NewIntVisitor(&x)
	err := vv.VisitInt(100)
	if err != nil {
		t.Error(err)
	}
	println(x)
}

func BenchmarkNewIntVisitor(b *testing.B) {
	x := 10
	for i := 0; i < b.N; i++ {
		vv := NewIntVisitor(&x)
		_ = vv.VisitInt(100)
	}
}
