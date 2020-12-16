# serde-go

serde-go is a golang port from [serde](https://serde.rs), intended to serialize and deserialize Golang data structures efficiently and generically.

## Tags

Supported tags are listed in [tags](./docs/tags.md)

## Quickstart

```go
// serde: deserialize,serialize
type Example struct {
	vint64   int64
	vmap     map[int]int
	varray   [2]int `serde:"skip"`
	vslice   []int
	vpointer *int
}
```

Use `serde-go` to generate deserialize and serialize for it:

```shell
go run -tags tools github.com/Xuanwo/serde-go/cmd/serde ./...
```

Use a deserializer and serializer to deserialize and serialize:

```go
import (
    "log"
    "testing"

    msgpack "github.com/Xuanwo/serde-msgpack-go"
)

func main() {
	ta := Example{
		A: "xxx",
	}
	content, err := msgpack.SerializeToBytes(&ta)
	if err != nil {
		log.Fatalf("msgpack SerializeToBytes: %v", err)
	}

	x := Example{}
	err = msgpack.DeserializeFromBytes(content, &x)
	if err != nil {
        log.Fatalf("msgpack DeserializeFromBytes: %v", err)
	}
	log.Printf("%#+v", x)
}
```