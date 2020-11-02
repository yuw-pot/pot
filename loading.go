// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pot

import (
	"fmt"
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	P "github.com/yuw-pot/pot/modules/properties"
	Z "github.com/yuw-pot/pot/modules/zlog"
	"time"
)

func (d *PoT) setMode(r *gin.Engine, mode interface{}) {
	switch mode {
	case data.ConsoleMode:
		// Mode: Debug
		r.Use(gin.Recovery())
		r.Use(d.loadLoggerWithFormat())

		return

	case data.ReleaseMode:
		// Mode Release

		// GeT Log Configure
		//   - data.ZLogPoT construct
		var zLogPoT *data.ZLogPoT = nil
		_ = P.PropertyPoT.UsK("Logs.PoT", &zLogPoT)

		if zLogPoT == nil {
			zLogPoT = &data.ZLogPoT{}
		}

		// - Zap Log New
		zLog := Z.New(zLogPoT)
		// - Zap Log Made
		zLogMade := zLog.Made()

		// Add a ginzap middleware, which:
		//   - Logs all requests, like a combined access and error log.
		//   - Logs to stdout.
		//   - RFC3339 with UTC time format.
		r.Use(ginzap.Ginzap(zLogMade, time.RFC3339, true))

		// Logs all panic to error log
		//   - stack means whether output the stack info.
		r.Use(ginzap.RecoveryWithZap(zLogMade, true))

		return
	}
}

func (d *PoT) loadLoggerWithFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[PoT] %v |	%v |	%v |	%v |	%v |	%v(%v)\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
		)
	})
}