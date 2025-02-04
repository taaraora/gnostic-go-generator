SHELL = /bin/bash

mkfile_dir := $(shell dirname $(abspath $(lastword $(MAKEFILE_LIST))))

build:	
	go install $(mkfile_dir)
	rm -f $(GOPATH)/bin/gnostic-go-client $(GOPATH)/bin/gnostic-go-server $(GOPATH)/bin/gnostic-go-types
	ln -s $(GOPATH)/bin/gnostic-go-generator $(GOPATH)/bin/gnostic-go-client
	ln -s $(GOPATH)/bin/gnostic-go-generator $(GOPATH)/bin/gnostic-go-server
	ln -s $(GOPATH)/bin/gnostic-go-generator $(GOPATH)/bin/gnostic-go-types

test:
	pushd examples/v2.0/sample && make test && popd
	pushd examples/v2.0/bookstore && make test && popd
	pushd examples/v3.0/bookstore && make test && popd


dev-test: build
	gnostic accounts.yaml --go-generator-out=bookstore