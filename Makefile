# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOPATH:=$(shell go env GOPATH)

default:
	mkdir -p ./_output
	go run makestatic.go
	protoc -I. --go_out=./_output pbgo.proto
	cp ./_output/github.com/chai2010/pbgo/pbgo.pb.go .
	go vet && go test

	make examples

.PHONY: examples
examples:
	cd ./examples/form.pb && make
	cd ./examples/hello.pb && make

clean:
	-rm *.pb.go
