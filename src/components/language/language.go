// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package language

import (
	"github.com/yuw-pot/pot/modules/languages"
)

type Component struct {
	Language string
}

func New() *Component {
	return &Component {
		Language: languages.EN,
	}
}

func (ln *Component) T(key string, replace ... interface{}) string {
	return languages.Translate(key, ln.Language, replace ...)
}


