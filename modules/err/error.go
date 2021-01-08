// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package err

import (
	"errors"
	"fmt"
	"github.com/yuw-pot/pot/data"
	"runtime"
)

type (
	PoT struct {
		ErrMsg *data.ErrH
	}
)

func (err *PoT) Initialized() {
	data.SeTErrMsg(err.ErrMsg)
}

func Err(pfx string, key string, content ...interface{}) error {
	return errors.New(data.GeTErrMsg(pfx, key, content ...))
}

func Position() interface{} {
	_, file, line, _ := runtime.Caller(1)
	return fmt.Sprintf("%v:%v", file, line)
}
