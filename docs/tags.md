## Tags

### Container

### Field

#### skip

`skip` is used to skip a field while serializing or deserializing.

```go
type Example struct {
    value `serde:"skip"`
}
```

#### skip_serialize

`skip_serialize` is used to skip a field while serializing.

```go
type Example struct {
    value `serde:"skip_serialize"`
}
```

#### skip_deserialize

`skip_deserialize` is used to skip a field while deserializing.

```go
type Example struct {
    value `serde:"skip_deserialize"`
}
```
