// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/chai2010/pbgo"
	"github.com/chai2010/pbgo/examples/form.pb"
	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	// http://localhost:8080
	router.GET("/", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		w.Header().Add("Content-Type", "text/html; charset=utf-8")

		fmt.Fprint(w, `
			<form method="POST" action="/api/comment">
				<input type="text" name="user" value="user name">
				<input type="email" name="email" value="123@gmail.com">
				<input name="url" value="http://chai2010.cn"></textarea>
				<textarea name="text" value="text">www</textarea>
				<input type="text" name="next" value="/thanks">
				<button type="submit">Send Test</button>
			</form>
		`)
	})

	router.GET("/thanks", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		fmt.Fprintln(w, "Thanks")
	})

	// curl -d "user=chai&email=11@qq.com&text=hello&next=http://github.com" localhost:8080/api/comment
	router.POST("/api/comment", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		var form form_pb.Comment
		if err := pbgo.BindForm(r, &form); err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			log.Println("err:", err)
			return
		}

		if form.Next != "" {
			http.Redirect(w, r, form.Next, http.StatusFound)
		}

		log.Println("form:", form)
	})

	log.Fatal(http.ListenAndServe(":8080", router))
}
