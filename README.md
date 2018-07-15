# pbgo: mini rpc/rest framework based on Protobuf

[![Build Status](https://travis-ci.org/chai2010/pbgo.svg)](https://travis-ci.org/chai2010/pbgo)
[![Go Report Card](https://goreportcard.com/badge/github.com/chai2010/pbgo)](https://goreportcard.com/report/github.com/chai2010/pbgo)
[![GoDoc](https://godoc.org/github.com/chai2010/pbgo?status.svg)](https://godoc.org/github.com/chai2010/pbgo)

## Install

1. `go get github.com/chai2010/pbgo`
1. `go get github.com/chai2010/pbgo/protoc-gen-pbgo`
1. `go run hello.go`

## Example (net/rpc)

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


## Example (rest)

```go
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/chai2010/pbgo/examples/hello.pb"
)

type HelloService struct{}

func (p *HelloService) Hello(request *hello_pb.String, reply *hello_pb.String) error {
	reply.Value = "hello:" + request.GetValue()
	return nil
}

func main() {
	router := hello_pb.HelloServiceHandler(new(HelloService))
	log.Fatal(http.ListenAndServe(":8080", router))
}
```

## BUGS

Report bugs to <chaishushan@gmail.com>.

Thanks!
