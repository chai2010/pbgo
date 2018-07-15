# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOPATH:=$(shell go env GOPATH)

default:
	protoc -I. --go_out=$(GOPATH)/src pbgo.proto
	go vet && go test

vgo_test:
	vgo test

clean:
	-rm *.pb.go
