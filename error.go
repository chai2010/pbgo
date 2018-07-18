// Copyright 2018 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pbgo

import (
	"fmt"
	"net/http"
)

type Error interface {
	HttpStatus() int
	Text() string
	error

	private()
}

type errHttpError struct {
	code int
	text string
}

func NewError(httpStatus int, text string) Error {
	if httpStatus == http.StatusOK {
		return nil
	}
	return &errHttpError{
		code: httpStatus,
		text: text,
	}
}

func NewErrorf(httpStatus int, fromat string, args ...interface{}) Error {
	if httpStatus == http.StatusOK {
		return nil
	}
	return &errHttpError{
		code: httpStatus,
		text: fmt.Sprintf(fromat, args...),
	}
}

func (p *errHttpError) Error() string {
	return fmt.Sprintf("pbgo error: httpStatus = %d desc = %s", p.code, p.text)
}
func (p *errHttpError) HttpStatus() int {
	return p.code
}
func (p *errHttpError) Text() string {
	return p.text
}

func (p *errHttpError) private() {}
