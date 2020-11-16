// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package languages

import (
	"github.com/yuw-pot/pot/data"
	"golang.org/x/text/language"
)

type TranslatePoT map[string]data.H

var (
	EN = language.English.String()
	CN = language.Chinese.String()
	CA = language.CanadianFrench.String()

	translation *TranslatePoT = &TranslatePoT {
		EN: {
			"tag_i": 	"i test",
			"tag_ii": 	"ii test",
			"tag_iii":	"iii %v test %v",
		},
		CN: {
			"tag_i": 	"i 测试",
			"tag_ii": 	"ii 测试",
			"tag_iii":	"iii %v 测试 %v",
		},
		CA: {

		},
	}
)
