SHELL := /bin/bash

build_internal_tools:
	@pushd internal/cmd/dummy_visitor \
			&& go build -o ../../bin/dummy_visitor . \
			&& popd
	@pushd internal/cmd/integer_visitor \
    		&& go build -o ../../bin/integer_visitor . \
    		&& popd

build: generate tidy

generate: build_internal_tools
	@echo "generate code"
	@go generate ./...
	@go fmt ./...
	@echo "ok"

tidy:
	@go mod tidy && go mod verify