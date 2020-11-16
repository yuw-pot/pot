// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package languages

import (
	"github.com/spf13/cast"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

type (
	PoT struct {
		TranslatePoT *TranslatePoT
	}
)

func Translate(key, ln string, replace ... interface{}) string {
	return message.NewPrinter(language.MustParse(ln)).Sprintf(key, replace ...)
}

func New() *PoT {
	return &PoT {}
}

func (ln *PoT) Initialized() *PoT {
	if ln.TranslatePoT != nil {
		for i, translated := range *ln.TranslatePoT {
			if translated != nil {
				for key, val := range translated {
					(*translation)[i][key] = val
				}
			}
		}
	}

	for i, translated := range *translation {
		for key, val := range translated {
			message.SetString(language.MustParse(i), key, cast.ToString(val))
		}
	}

	return ln
}



