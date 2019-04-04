// Copyright 2019 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/golang/protobuf/proto"
)

func BindQuery(r *http.Request, req proto.Message) error {
	return PopulateQueryParametersEx(req, r.URL.Query(), true)
}

func BindBody(r *http.Request, req proto.Message) error {
	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil && err != io.EOF {
		return err
	}
	return nil
}

func BindNamedParams(
	r *http.Request, req proto.Message,
	params interface{ ByName(name string) string },
	name ...string,
) error {
	for _, fieldPath := range name {
		err := PopulateFieldFromPath(req,
			fieldPath, params.ByName(fieldPath),
		)
		if err != nil {
			return err
		}
	}
	return nil
}
