# Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

default:
	protoc -I=. --go_out=. comment.proto
	go fmt && go vet && go test

protoc:
	go get github.com/golang/protobuf/protoc-gen-go
	go install github.com/golang/protobuf/protoc-gen-go

clean:
	-rm *.pb.go
