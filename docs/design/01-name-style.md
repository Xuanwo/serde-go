# Name Style

serde should follow a comment name styles, include:

- struct fields enum type
- struct fields enum value
- struct field visitor type
- new struct filed vistior type
- struct visitor
- new struct visitor
- map/slice vistior
- ...

## Rules

- Make sure those types will not conflict with user types

So let's add a `serde` prefix for all internal types, make sure they are unique and not exported.

- Make sure all types are easy to recognize

```go
type Test struct {
	B string
	b string
}
```

We will use `_` to distinguish serde part and user part, and use Pascal style in the part

so we will got a following types:

```go
serdeStruct_Test
serdeStructEnum_Test
serdeStructEnum_Test_B
serdeNewStruct_Test
serdeNewStrcutFiled_Test
```

- Always keep original struct name and type name

```go
type Test struct {
	B string
	b string
}

type test struct {
	B string
	b string
}
```

we should keep original name so that we will not mess up `B` and `b`:

```go
serdeStruct_Test
serdeStruct_test
serdeStructEnum_Test_B
serdeStructEnum_Test_b
```

- Use map/slice's type as their name:
    - `map[A]B` => `serdeMap_A_B`
    - `[]A` => `serdeSlice_A`

```go
type Test struct {
	B map[int]int
	b string
}
```

We will get these types:

```go
serdeMap_int_int
```

---

- [Serde Name Style](https://github.com/Xuanwo/serde-go/issues/4)
