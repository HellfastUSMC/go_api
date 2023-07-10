.PHONY: test
test:
	go test -v -race -timeout 30s ./ ...

.PHONY: build
build:
	go build -v ./cmd/apiserver

.DEFAULT_GOAL := build