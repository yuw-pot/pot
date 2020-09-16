// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/spf13/cast"
	"strings"
)

type (
	Utils struct {

	}
)

func New() *Utils {
	return &Utils {

	}
}

func (u *Utils) Contains(k string, d ...interface{}) bool {
	if len(d) < 1 {
		return false
	}

	for _, v := range d {
		if strings.Contains(k, cast.ToString(v)) {
			return true
		}
	}

	return false
}

//func (u *Utils) Merge(d ...*data.H) (src *data.H) {
//	if len(d) < 1 {
//		return
//	}
//
//
//
//	return
//}
