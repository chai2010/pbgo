// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func JSON(w http.ResponseWriter, code int, obj interface{}) {
	s, err := json.Marshal(obj)
	if err != nil {
		log.Panicln(err)
	}

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")

	fmt.Fprintln(w, string(s))
}
