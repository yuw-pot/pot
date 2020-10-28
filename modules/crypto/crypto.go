// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package crypto

import "github.com/yuw-pot/pot/data"

type (
	crypto interface {
		made(d ... interface{}) (interface{}, error)
	}

	PoT struct {
		Mode interface{}
		D []interface{}
	}
)

var (
	_ crypto = &RsAPoT{}
	_ crypto = &SimplyPoT{}
)

func New() *PoT {
	return &PoT {}
}

func (cryptoPoT *PoT) Made() (interface{}, error) {
	switch cryptoPoT.Mode {
	case data.ModeToken:
		var res crypto = newToken()
		return res.made(cryptoPoT.D ...)

	case data.ModeRsA:
		var res crypto = newRsA()
		return res.made(cryptoPoT.D ...)

	default:
		return "", nil
	}
}
