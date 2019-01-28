# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOPATH:=$(shell go env GOPATH)

default:
	go fmt github.com/chai2010/pbgo/protoc-gen-pbgo/...
	go install github.com/chai2010/pbgo/protoc-gen-pbgo
	protoc -I=. -I=../../ --pbgo_out=. hello.proto
	#protoc -I=. -I=../../ --pbgo_out=plugins=grpc+pbgo:. hello.proto
	go fmt && go vet && go test

clean:
	-rm *.pb.go
