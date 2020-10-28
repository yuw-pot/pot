// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package languages

import (
	U "github.com/yuw-pot/pot/modules/utils"
	"golang.org/x/text/language"
)

var (
	EN = language.English.String()
	CN = language.English.String()
)

type (
	PoT struct {
		vs *U.PoT
	}
)

func New() *PoT {
	return &PoT {
		vs: U.New(),
	}
}

