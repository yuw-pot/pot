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
	tokenPoT struct {

	}
)

func newToken() *tokenPoT {
	return &tokenPoT {

	}
}

func (token *tokenPoT) made(d ... interface{}) (interface{}, error) {
	if len(d) <= 1 {
		return nil, E.Err(data.ErrPfx, "TokenParamsErr")
	}

	if d[1] == "" {
		return nil, E.Err(data.ErrPfx, "TokenParamsErr")
	}

	switch d[0] {
	case data.MD5:
		return token.md5(d ...)

	case data.SHA1:
		return token.sha1(d ...)

	case data.SHA256:
		return token.sha256(d ...)

	default:
		return nil, E.Err(data.ErrPfx, "TokenTypeErr")
	}
}

func (token *tokenPoT) md5(d ... interface{}) (interface{}, error) {
	res := md5.New()
	res.Write([]byte(cast.ToString(d[1])))

	return hex.EncodeToString(res.Sum([]byte(""))), nil
}

func (token *tokenPoT) sha1(d ... interface{}) (interface{}, error)  {
	res := sha1.New()
	res.Write([]byte(cast.ToString(d[1])))

	return hex.EncodeToString(res.Sum([]byte(""))), nil
}

func (token *tokenPoT) sha256(d ... interface{}) (interface{}, error) {
	res := sha256.New()
	res.Write([]byte(cast.ToString(d[1])))

	return hex.EncodeToString(res.Sum([]byte(""))), nil
}
