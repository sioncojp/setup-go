REVISION := $(shell git describe --always)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS	:= -ldflags="-X \"main.Revision=$(REVISION)\" -X \"main.BuildDate=${DATE}\""

.PHONY: build deps deps/updateclean run help

name		:= setup-go
darwin_name	:= $(name)-darwin-amd64

build: ## go build
	go build -o bin/$(name) $(LDFLAGS) *.go

clean: ## remove bin/*
	rm -f bin/*

run: ## go run
	go run main.go

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "\033[36m%-22s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST)
