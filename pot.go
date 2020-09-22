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
	Z "github.com/yuw-pot/pot/modules/zlog"
	R "github.com/yuw-pot/pot/routes"
	"time"

	_ "github.com/yuw-pot/pot/autoload"
)

const (
	Version string = "v1.0.0"
	ginPort string = "8888"
)

type (
	PoT struct {
		PoTRoute *R.PoT
		PoTError *E.PoT
	}
)

func New() *PoT {
	return &PoT {}
}

func (engine *PoT) Run() {
	// Disable Console Color
	gin.DisableConsoleColor()

	// Set Mode by YamlConfigure
	//   - 0->debug
	//   - 1->release

	codePoTMode := cast.ToInt(P.PropertyPoT.GeT("PoT.Mode", 1))
	if codePoTMode != 0 && codePoTMode != 1 {
		panic(E.Err(data.ErrPfx, "PoTModeErr"))
	}

	gin.SetMode(data.PoTMode[codePoTMode])

	// GeT Log Configure
	// - data.ZLogPoT construct
	var zLogPoT *data.ZLogPoT = nil
	_ = P.PropertyPoT.UsK("Logs.PoT", &zLogPoT)
	if zLogPoT == nil {
		panic(E.Err(data.ErrPfx, "PoTZapLogErr"))
	}

	// - Zap Log New
	zLog := Z.New(zLogPoT)
	// - Zap Log Made
	zLogMade := zLog.Made()

	r := gin.New()

	// Add a ginzap middleware, which:
	//   - Logs all requests, like a combined access and error log.
	//   - Logs to stdout.
	//   - RFC3339 with UTC time format.
	r.Use(ginzap.Ginzap(zLogMade, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zLogMade, true))

	R.RPoT.Made(r)

	// Https Power ON/OFF
	//   - PoT.Hssl
	//     - PoT.Hssl.Power
	//     - PoT.Hssl.CertFile
	//     - PoT.Hssl.KeysFile
	var rErr error
	strPoTPort := ":"+cast.ToString(P.PropertyPoT.GeT("PoT.Port", ginPort))

	if P.PropertyPoT.GeT("PoT.Hssl.Power", 0) == 1 {
		PoTHsslCertFile := cast.ToString(P.PropertyPoT.GeT("PoT.Hssl.CertFile", ""))
		if PoTHsslCertFile == "" {
			panic(E.Err(data.ErrPfx, "PoTSslCF"))
		}

		PoTHsslKeysFile := cast.ToString(P.PropertyPoT.GeT("PoT.Hssl.KeysFile", ""))
		if PoTHsslKeysFile == "" {
			panic(E.Err(data.ErrPfx, "PoTSslKF"))
		}

		//   - Run Https Server (SSL)
		rErr = r.RunTLS(strPoTPort, PoTHsslCertFile, PoTHsslKeysFile)

	} else {
		//   - Run Http Server
		rErr = r.Run(strPoTPort)
	}

	if rErr != nil {
		zLogMade.Fatal("PoT Server Start Failure", zLog.ZErr(rErr))
	}
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


