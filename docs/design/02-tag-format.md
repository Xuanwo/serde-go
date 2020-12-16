# Tag Format

The document will describe the tag format in `serde-go`.

## Goal

It's obvious that `serde-go` will have a lot of tags, so the tag format should be:

- Unified
- Understandable
- Extensible

## Proposal

`serde-go` will adopt the tag format following:

```go
// serde: deserialize,serialize
// serde: default,rename_all=xxx,rename_all_serialize=xxx
type Example struct {
	value  [2]int `serde:"skip,rename=xx"`
}
```

Tags have two categories: `container` and `field`

- `container` applies for structs or other container type, show up as comments: `serde: xxx`
- `field` applies for structs fields, show up as struct tags: `serde:"xxx"`

The specific format will be:

- Split via `,`: `skip,rename` (tailing `,` will be ignored)
- Only support bool and string type: `skip,rename=abc`, string value will be explicit instead of surrounding by `"`
- All available key will be in snake case: `rename_all,skip_if`
- Specially, we will support use `_deserialize` or `_serialize` as suffix of a key to represent this key only applied to `deserialize` or `serialize`: `rename_deserialize=A,rename_serialize=B`

## Relational

### serde-rs

The original version of `serde-go` is written in rust, so it adopts the rust's [Attributes](https://doc.rust-lang.org/book/attributes.html):

```rust
#[derive(Serialize, Deserialize)]
#[serde(deny_unknown_fields)]  // <-- this is a container attribute
struct S {
    #[serde(default)]  // <-- this is a field attribute
    f: i32,
}

#[derive(Serialize, Deserialize)]
#[serde(rename = "e")]  // <-- this is also a container attribute
enum E {
    #[serde(rename = "a")]  // <-- this is a variant attribute
    A(String),
}
```

It supports following styles:

- bool type: `#[serde(default)]`
- string type: `#[serde(rename = "a")]`
- struct type: `#[serde(rename_all(serialize = "...", deserialize = "..."))]`

### go/json

The stdlib in go support following tags:

```go
// Field appears in JSON as key "myName".
Field int `json:"myName"`

// Field appears in JSON as key "myName" and
// the field is omitted from the object if its value is empty,
// as defined above.
Field int `json:"myName,omitempty"`

// Field appears in JSON as key "Field" (the default), but
// the field is skipped if empty.
// Note the leading comma.
Field int `json:",omitempty"`

// Field is ignored by this package.
Field int `json:"-"`

// Field appears in JSON as key "-".
Field int `json:"-,"`
```

The first key always means the name for the field.

### vmihailenco/msgpack

[msgpack](https://github.com/vmihailenco/msgpack) supports following tags:

```go
msgpack:"my_field_name"
msgpack:"alias:another_name"
msgpack:",omitempty"

type ItemOmitEmpty struct {
	_msgpack struct{} `msgpack:",omitempty"`
	Foo      string
	Bar      string
}
```

Almost the same with go/json.
