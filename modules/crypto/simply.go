// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/hex"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
)

type (
	SimplyPoT struct {

	}
)

func newSimply() *SimplyPoT {
	return &SimplyPoT {

	}
}

func (simply *SimplyPoT) made(d ... interface{}) (interface{}, error) {
	if len(d) <= 1 {
		return nil, E.Err(data.ErrPfx, "TokenParamsErr")
	}

	if d[1] == "" {
		return nil, E.Err(data.ErrPfx, "TokenParamsErr")
	}

	switch d[0] {
	case data.MD5:
		return simply.md5(d ...)

	case data.SHA1:
		return simply.sha1(d ...)

	case data.SHA256:
		return simply.sha256(d ...)

	default:
		return nil, E.Err(data.ErrPfx, "TokenTypeErr")
	}
}

func (simply *SimplyPoT) md5(d ... interface{}) (interface{}, error) {
	res := md5.New()
	res.Write([]byte(cast.ToString(d[1])))

	return hex.EncodeToString(res.Sum([]byte(""))), nil
}

func (simply *SimplyPoT) sha1(d ... interface{}) (interface{}, error)  {
	res := sha1.New()
	res.Write([]byte(cast.ToString(d[1])))

	return hex.EncodeToString(res.Sum([]byte(""))), nil
}

func (simply *SimplyPoT) sha256(d ... interface{}) (interface{}, error) {
	res := sha256.New()
	res.Write([]byte(cast.ToString(d[1])))

	return hex.EncodeToString(res.Sum([]byte(""))), nil
}
