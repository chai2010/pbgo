// Copyright 2020 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ignore

package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/chai2010/pbgo"
)

func main() {
	var reply struct {
		CurrentCondition []struct {
			FeelsLikeC       string `json:"FeelsLikeC"`
			FeelsLikeF       string `json:"FeelsLikeF"`
			Cloudcover       string `json:"cloudcover"`
			Humidity         string `json:"humidity"`
			LocalObsDateTime string `json:"localObsDateTime"`
			Observation_time string `json:"observation_time"`
		} `json:"current_condition"`
	}

	err := pbgo.HttpGet("http://wttr.in/wuhan?format=j1", nil, &reply)
	if err != nil {
		log.Fatal(err)
	}

	json, err := json.MarshalIndent(reply, "", "  ")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(json))
}
