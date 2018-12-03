package main

import (
	"html/template"
	"log"
	"os"
)

// Inventory ...for go-template
type Inventory struct {
	RepoName reponame
}

// Get ... get all template name
func (t *Templates) Get() map[string]string {
	return map[string]string{
		".gitignore": gitignoreTemplate,
		"Makefile":   makefileTemplate,
		"build.sh":   buildShellTemplate,
	}
}

// Create ...create template
func (t *Templates) Create(rn reponame) {
	templates := t.Get()
	in := Inventory{RepoName: rn}
	for k, v := range templates {
		// check Template exist
		if FileExist(k) {
		  continue
		}

		func() {
			f, err := os.Create(k)
			if err != nil {
				log.Printf("[WARN] cannot create %s template: %s\n", k, err)
				return
			}
			defer f.Close()

			tpl := template.Must(template.New(k).Parse(v))
			err = tpl.Execute(f, in)
			if err != nil {
				return
			}
			log.Printf("[INFO] create template: %s\n", k)
		}()
	}
}

const gitignoreTemplate = `_vendor
.vendor
vendor
_vendor-*
*.pid
*.swp
.DS_Store
tmp/*
bin/

### Other ###
.idea/

### .gitkeep ###
!.gitkeep
`

const makefileTemplate = `REVISION := $(shell git describe --always)
DATE := $(shell date +%Y-%m-%dT%H:%M:%S%z)
LDFLAGS	:= -ldflags="-X \"main.Revision=$(REVISION)\" -X \"main.BuildDate=${DATE}\""

.PHONY: build-cross dist build deps deps/update clean run help

name		:= {{.RepoName}}
linux_name	:= $(name)-linux-amd64
darwin_name	:= $(name)-darwin-amd64

dist: build-docker ## create .tar.gz linux & darwin to /bin
	cd bin && tar zcvf $(linux_name).tar.gz $(linux_name) && rm -f $(linux_name)
	cd bin && tar zcvf $(darwin_name).tar.gz $(darwin_name) && rm -f $(darwin_name)

build-cross: ## create to build for linux & darwin to bin/
	GOOS=linux GOARCH=amd64 go build -o bin/$(linux_name) $(LDFLAGS) cmd/$(name)/*.go
	GOOS=darwin GOARCH=amd64 go build -o bin/$(darwin_name) $(LDFLAGS) cmd/$(name)/*.go

build: ## go build
	go build -o bin/$(name) $(LDFLAGS) cmd/$(name)/*.go

build-docker: ## go build on Docker
	@docker run --rm -v "$(PWD)":/go/src/github.com/sioncojp/$(name) -w /go/src/github.com/sioncojp/$(name) golang:latest bash build.sh

test: ## go test
	go test -v $$(go list ./... | grep -v /vendor/)

dep: ## dep ensure
	dep ensure

dep/init: ## dep init
	dep init

dep/update: ## dep update
	dep ensure -update

clean: ## remove bin/*
	rm -f bin/*

run: ## go run
	go run cmd/$(name)/main.go -c examples/config.toml

lint: ## go lint ignore vendor
	golint $(go list ./... | grep -v /vendor/)

help:
	@awk -F ':|##' '/^[^\t].+?:.*?##/ { printf "\033[36m%-22s\033[0m %s\n", $$1, $$NF }' $(MAKEFILE_LIST)
`

const buildShellTemplate = `#!/bin/bash -
for GOOS in darwin linux; do
    GOOS=$GOOS GOARCH=amd64 go build -o bin/{{.RepoName}}-$GOOS-amd64 cmd/{{.RepoName}}/*.go
done
`
