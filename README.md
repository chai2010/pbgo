# pbgo: mini rpc/rest framework based on Protobuf

[![Build Status](https://travis-ci.org/chai2010/pbgo.svg)](https://travis-ci.org/chai2010/pbgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/pbgo)](https://goreportcard.com/report/github.com/chai2010/pbgo)
[![GoDoc](https://godoc.org/github.com/chai2010/pbgo?status.svg)](https://godoc.org/github.com/chai2010/pbgo)

## Install

1. install `protoc` at first: http://github.com/google/protobuf/releases
1. `go get github.com/chai2010/pbgo`
1. `go get github.com/chai2010/pbgo/protoc-gen-pbgo`
1. `go run hello.go`

## Example (net/rpc)

create proto file:

```protobuf
syntax = "proto3";
package hello_pb;

message String {
	string value = 1;
}
message Message {
	string value = 1;
}

service HelloService {
	rpc Hello (String) returns (String);
	rpc Echo (Message) returns (Message);
}
```

generate rpc code:

```
$ protoc -I=. -I=$(GOPATH)/src --pbgo_out=. hello.proto
```

use generate code:

```go
package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/chai2010/pbgo/examples/hello.pb"
)

type HelloService struct{}

func (p *HelloService) Hello(request *hello_pb.String, reply *hello_pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
func (p *HelloService) Echo(request *hello_pb.Message, reply *hello_pb.Message) error {
	*reply = *request
	return nil
}

func startRpcServer() {
	hello_pb.RegisterHelloService(rpc.DefaultServer, new(HelloService))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("ListenTCP error:", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error:", err)
		}

		go rpc.ServeConn(conn)
	}
}

func tryRpcClient() {
	client, err := hello_pb.DialHelloService("tcp", "localhost:1234")
	if err != nil {
		log.Fatal(err)
	}

	reply, err := client.Hello(&hello_pb.String{Value: "gopher"})
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(reply.GetValue())
}

func main() {
	go startRpcServer()
	tryRpcClient()
}
```


## Example (rest api)

create proto file:

```protobuf
syntax = "proto3";
package hello_pb;

import "github.com/chai2010/pbgo/pbgo.proto";

message String {
	string value = 1;
}

message Message {
	string value = 1;
	repeated int32 array = 2;
	map<string,string> dict = 3;
	String subfiled = 4;
}

message StaticFile {
	string content_type = 1;
	bytes content_body = 2;
}

service HelloService {
	rpc Hello (String) returns (String) {
		option (pbgo.rest_api) = {
			get: "/hello/:value"
			post: "/hello"

			additional_bindings {
				method: "DELETE"; url: "/hello"
			}
			additional_bindings {
				method: "PATCH"; url: "/hello"
			}
		};
	}
	rpc Echo (Message) returns (Message) {
		option (pbgo.rest_api) = {
			get: "/echo/:subfiled.value"
		};
	}
	rpc Static(String) returns (StaticFile) {
		option (pbgo.rest_api) = {
			additional_bindings {
				method: "GET";
				url: "/static/:value";
				content_type: ":content_type";
				content_body: ":content_body"
			}
		};
	}
}
```

generate rpc/rest code:

```
$ protoc -I=. -I=$(GOPATH)/src --pbgo_out=. hello.proto
```

use generate code:

```go
package main

import (
	"io/ioutil"
	"log"
	"mime"
	"net/http"
	"time"

	"github.com/chai2010/pbgo/examples/hello.pb"
)

type HelloService struct{}

func (p *HelloService) Hello(request *hello_pb.String, reply *hello_pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}
func (p *HelloService) Echo(request *hello_pb.Message, reply *hello_pb.Message) error {
	*reply = *request
	return nil
}
func (p *HelloService) Static(request *hello_pb.String, reply *hello_pb.StaticFile) error {
	data, err := ioutil.ReadFile("./testdata/" + request.Value)
	if err != nil {
		return err
	}

	reply.ContentType = mime.TypeByExtension(request.Value)
	reply.ContentBody = data
	return nil
}

func someMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, r *http.Request) {
		timeStart := time.Now()
		defer func() {
			timeElapsed := time.Since(timeStart)
			log.Println(r.Method, r.URL, timeElapsed)
		}()

		next.ServeHTTP(wr, r)
	})
}

func main() {
	router := hello_pb.HelloServiceHandler(new(HelloService))
	log.Fatal(http.ListenAndServe(":8080", someMiddleware(router)))
}
```

```
$ curl localhost:8080/hello/gopher
{"value":"hello:gopher"}
$ curl localhost:8080/hello/gopher?value=vgo
{"value":"hello:vgo"}
$ curl localhost:8080/hello -X POST --data '{"value":"cgo"}'
{"value":"hello:cgo"}

$ curl localhost:8080/echo/gopher
{"subfiled":{"value":"gopher"}}
$ curl "localhost:8080/echo/gopher?array=123&array=456"
{"array":[123,456],"subfiled":{"value":"gopher"}}
$ curl "localhost:8080/echo/gopher?dict%5Babc%5D=123"
{"dict":{"abc":"123"},"subfiled":{"value":"gopher"}}

$ curl localhost:8080/static/gopher.png
$ curl localhost:8080/static/hello.txt
```

## Form Example

Create [example/form_pb/comment.proto](example/form_pb/comment.proto):

```proto
syntax = "proto3";
package form_pb;

message Comment {
	string user = 1;
	string email = 2;
	string url = 3;
	string text = 4;
	string next = 5; // redirect
}
```

generate proto code:

```
$ protoc -I=. --go_out=. comment.proto
```

create web server:

```go
package main

import (
	"github.com/chai2010/pbgo"
	"github.com/chai2010/pbgo/examples/form_pb"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// curl -d "user=chai&email=11@qq.com&text=hello&next=http://github.com" localhost:8080/api/comment-form
	router.POST("/api/comment-form", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var form form_pb.Comment
		if err := pbgo.BindForm(r, &form); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println(err)
			return
		}

		s, _ := json.Marshal(form)
		fmt.Fprintln(w, string(s))

		if form.Next != "" {
			http.Redirect(w, r, form.Next, http.StatusFound)
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
```

```
$ curl -d "user=chai&email=11@qq.com&text=hello&next=http://github.com" localhost:8080/api/comment-form
{"user":"chai","email":"11@qq.com","text":"hello","next":"http://github.com"}
```

## gRPC & Rest Example

See https://github.com/chai2010/pbgo-grpc

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
