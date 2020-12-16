package test

//go:generate go run -tags tools ../ .

// serde: deserialize,serialize
// serde: default,rename_all=xxx, rename_all_serialize=xxx
type Test struct {
	vint64   int64
	vmap     map[int]int
	varray   [2]int `serde:"skip,rename=xx,"`
	vslice   []int
	vpointer *int
}
