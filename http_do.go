// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func HttpGet(urlpath string, input, output interface{}) error {
	return httpGet(urlpath, input, output)
}

func HttpPost(urlpath string, input, output interface{}) error {
	return httpDo("POST", urlpath, input, output)
}

func HttpDo(method, urlpath string, input, output interface{}) error {
	if method == "GET" {
		return httpGet(urlpath, input, output)
	} else {
		return httpDo(method, urlpath, input, output)
	}
}

func NewHttpRequest(method, urlpath string, input interface{}) (*http.Request, error) {
	url, err := url.Parse(urlpath)
	if err != nil {
		return nil, err
	}

	var body io.Reader
	if method == "GET" {
		if input != nil {
			urlValues, err := BuildUrlValues(input)
			if err != nil {
				return nil, err
			}
			for k, v := range url.Query() {
				urlValues[k] = v
			}
			url.RawQuery = urlValues.Encode()
		}
	} else {
		reqBody, err := json.Marshal(input)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequest(method, url.String(), body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	return req, nil
}

func httpGet(urlpath string, input, output interface{}) error {
	url, err := url.Parse(urlpath)
	if err != nil {
		return err
	}

	if input != nil {
		urlValues, err := BuildUrlValues(input)
		if err != nil {
			return err
		}
		for k, v := range url.Query() {
			urlValues[k] = v
		}
		url.RawQuery = urlValues.Encode()
	}

	r, err := http.Get(url.String())
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

func httpDo(method, urlpath string, input, output interface{}) error {
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
