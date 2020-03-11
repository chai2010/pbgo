// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func HttpDo(method, urlpath string, input, output proto.Message) error {
	if method == "GET" {
		return httpGet(urlpath, input, output)
	} else {
		return httpDo(method, urlpath, input, output)
	}
}

func httpGet(urlpath string, input, output proto.Message) error {
	urlValues, err := BuildUrlValues(input)
	if err != nil {
		return err
	}

	r, err := http.Get(urlpath + "?" + urlValues.Encode())
	if err != nil {
		return err
	}
	defer r.Body.Close()

	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyData, output)
	if err != nil {
		return err
	}

	return nil
}

func httpDo(method, urlpath string, input, output proto.Message) error {
	reqBody, err := json.Marshal(input)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(method, urlpath, bytes.NewReader(reqBody))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	r, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	bodyData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bodyData, output)
	if err != nil {
		return err
	}

	return nil
}
