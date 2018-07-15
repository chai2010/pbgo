# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

GOPATH:=$(shell go env GOPATH)

default:
	protoc -I=. -I=$(GOPATH)/src --go_out=. hello.proto

clean:
	-rm *.pb.go
