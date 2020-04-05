// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"log"

	"github.com/chai2010/pbgo/examples/hello.pb"
)

func main() {
	c := hello_pb.NewHelloServiceHttpClient("http://127.0.0.1:8080")

	reply, err := c.Echo(&hello_pb.Message{
		Value:    "chai2010",
		Array:    []int32{1, 2, 3},
		Subfiled: &hello_pb.String{Value: "xx"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// {chai2010 [1 2 3] map[] value:"xx"  {} [] 0}
	fmt.Println(*reply)
}
