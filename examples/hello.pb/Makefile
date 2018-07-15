# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOPATH:=$(shell go env GOPATH)

default:
	go install github.com/chai2010/pbgo/protoc-gen-pbgo
	protoc -I=. -I=$(GOPATH)/src --pbgo_out=plugins=netrpc:. hello.proto

clean:
	-rm *.pb.go
