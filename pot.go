// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pot

import (
	"github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	U "github.com/yuw-pot/pot/modules/utils"
	Z "github.com/yuw-pot/pot/modules/zlog"
	R "github.com/yuw-pot/pot/routes"
	"time"

	_ "github.com/yuw-pot/pot/autoload"
)

const version string = "v1.0.0"

type (
	PoT struct {
		vPoT *U.PoT

		PoTRoute *R.PoT
		PoTError *E.PoT
	}
)

func New() *PoT {
	return &PoT {
		vPoT: U.New(),
	}
}

func (engine *PoT) Run() {
	engine.vPoT.Fprintf(gin.DefaultWriter, "[%v], %v\n", data.PoT, version)

	// Disable Console Color
	gin.DisableConsoleColor()

	// Gin Mode Release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	PoTMode := P.PropertyPoT.GeT("PoT.Mode", "1")
	if engine.vPoT.Contains(PoTMode, 0,1) == false {
		PoTMode = 1
	}

	engine.setMode(r, data.PoTMode[cast.ToInt(PoTMode)])

	R.RPoT.Made(r)

	// Https Power ON/OFF
	//   - PoT.Hssl
	//     - PoT.Hssl.Power
	//     - PoT.Hssl.CertFile
	//     - PoT.Hssl.KeysFile
	var rErr error
	strPoTPort := ":"+cast.ToString(P.PropertyPoT.GeT("PoT.Port", data.PropertyPort))

	if P.PropertyPoT.GeT("PoT.Hssl.Power", 0) == 1 {
		PoTHsslCertFile := cast.ToString(P.PropertyPoT.GeT("PoT.Hssl.CertFile", ""))
		if PoTHsslCertFile == "" {
			panic(E.Err(data.ErrPfx, "PoTSslCF"))
		}

		PoTHsslKeysFile := cast.ToString(P.PropertyPoT.GeT("PoT.Hssl.KeysFile", ""))
		if PoTHsslKeysFile == "" {
			panic(E.Err(data.ErrPfx, "PoTSslKF"))
		}

		engine.vPoT.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTPs on %v\n", data.PoT, strPoTPort)

		//   - Run Https Server (SSL)
		rErr = r.RunTLS(strPoTPort, PoTHsslCertFile, PoTHsslKeysFile)

	} else {
		engine.vPoT.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTP on %v\n", data.PoT, strPoTPort)

		//   - Run Http Server
		rErr = r.Run(strPoTPort)
	}

	if rErr != nil { panic(rErr) }
}

func (engine *PoT) PoT() *PoT {
	// Routes Initialized
	R.RPoT = engine.PoTRoute

	// Err Modules Initialize
	//   - Combine Error Message of Self Define
	E.EPoT = engine.PoTError
	E.EPoT.ErrPoTCombine()

	return engine
}

func (engine *PoT) setMode(r *gin.Engine, mode interface{}) {
	switch cast.ToString(mode) {
	case data.DebugMode:
		// Mode: Debug
		r.Use(gin.Recovery())
		r.Use(R.LoggerWithFormat())

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


