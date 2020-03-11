// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/golang/protobuf/proto"
)

func BuildUrlValues(msg proto.Message) (m url.Values, err error) {
	defer func() {
		if x := recover(); x != nil {
			err = fmt.Errorf("pbgo.BuildUrlValues failed: %v", x)
		}
	}()

	d, err := json.Marshal(msg)
	if err != nil {
		return nil, err
	}

	var mapx map[string]interface{}
	if err := json.Unmarshal(d, &mapx); err != nil {
		return nil, err
	}

	m = pkgUnpackMapXToMapString(mapx)
	return m, nil
}

// X is oneof string/float64/[]interface/map[string]interface{}
func pkgUnpackMapXToMapString(mapx map[string]interface{}) map[string][]string {
	var m = map[string][]string{}
	for k, v := range mapx {
		switch v := v.(type) {
		case string:
			m[k] = append(m[k], v)
		case float64:
			m[k] = append(m[k], fmt.Sprintf("%v", v))
		case []interface{}:
			for i := 0; i < len(v); i++ {
				switch vi := v[i].(type) {
				case string:
					m[k] = append(m[k], vi)
				case float64:
					m[k] = append(m[k], fmt.Sprintf("%v", vi))
				case map[string]interface{}:
					for kk, vv := range pkgUnpackMapXToMapString(vi) {
						m[k+"."+kk] = append(m[k+"."+kk], vv...)
					}
				default:
					// unreachable
				}
			}
		case map[string]interface{}:
			for kk, vv := range pkgUnpackMapXToMapString(v) {
				m[k+"."+kk] = append(m[k+"."+kk], vv...)
			}
		default:
			// unreachable
		}
	}
	return m
}
