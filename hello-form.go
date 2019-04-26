// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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
