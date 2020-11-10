// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	"github.com/yuw-pot/pot/data"
	"time"
)

func getTimeUniT(method string) time.Duration {
	switch method {
	case data.TimeHour:
		return time.Hour

	case data.TimeMinute:
		return time.Minute

	case data.TimeSecond:
		return time.Second

	default:
		return time.Hour
	}
}
