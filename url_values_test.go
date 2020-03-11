// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo_test

import (
	"fmt"

	"github.com/chai2010/pbgo"
	"github.com/chai2010/pbgo/examples/hello.pb"
)

func ExampleBuildUrlValues() {
	var msg = &hello_pb.Message{
		Value: "chai2010",
		Array: []int32{1, 2, 3},
		Dict: map[string]string{
			"aaa": "111",
			"bbb": "222",
		},
		Subfiled: &hello_pb.String{
			Value: "value",
		},
	}

	x, err := pbgo.BuildUrlValues(msg)
	if err != nil {
		panic(err)
	}

	fmt.Println(x.Encode())

	// Output:
	// array=1&array=2&array=3&dict.aaa=111&dict.bbb=222&subfiled.value=value&value=chai2010
}
