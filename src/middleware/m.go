// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/yuw-pot/pot/libs"
	"github.com/yuw-pot/pot/modules/utils"
)

type m struct {
	lib *libs.PoT
	v *utils.PoT
}

func New() *m {
	return &m {
		lib: new(libs.PoT),
		v: utils.New(),
	}
}
