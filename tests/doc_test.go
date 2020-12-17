/*
This package is used for serde-go tests.

We implement Serializer and Deserializer which store and parse interfaces directly to
verify whether serde-go works as expected.
*/
package tests

//go:generate go run -tags tools ../cmd/serde ./...
