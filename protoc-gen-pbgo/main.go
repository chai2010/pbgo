// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// protoc-gen-pbgo is a plugin for the Protobuf compiler to generate rpc/rest code.
package main

import (
	"strings"

	mainPkg "github.com/chai2010/pbgo/protoc-gen-pbgo/main"
)

func main() {
	mainPkg.Main(func(parameter string) string {
		if !strings.Contains(parameter, "plugins=") {
			return ",plugins=" + pbgoPluginName
		} else {
			return parameter
		}
	})
}
