// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package err

import (
	"errors"
	"fmt"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"runtime"
)

var EPoT *PoT = new(PoT)

type (
	PoT struct {
		ErrMsg *data.ErrH
	}
)

func (err *PoT) ErrPoTCombine() {
	if err.ErrMsg != nil {
		for k, v := range *err.ErrMsg {
			if k != data.ErrPfx {
				(*data.ErrMsg)[k] = v
			}
		}
	}
}

func Err(pfx string, key string, content ...string) error {
	str := cast.ToString((*data.ErrMsg)["PoT"]["ErrDefault"])

	s, ok := (*data.ErrMsg)[pfx][key]
	if ok {
		str = cast.ToString(s)
	}

	return errors.New(str)
}

func Position() interface{} {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}
