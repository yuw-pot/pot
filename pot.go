// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pot

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	U "github.com/yuw-pot/pot/modules/utils"
	R "github.com/yuw-pot/pot/routes"

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

func (d *PoT) Run() {
	d.vPoT.Fprintf(gin.DefaultWriter, "[%v] %v\n", data.PoT, version)

	// Disable Console Color
	gin.DisableConsoleColor()

	// Gin Mode Release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	PoTMode := P.PropertyPoT.GeT("PoT.Mode", "1")
	if d.vPoT.Contains(PoTMode, 0,1) == false {
		PoTMode = 1
	}

	d.setMode(r, data.PoTMode[cast.ToInt(PoTMode)])

	R.RPoT.Eng = r
	r = R.RPoT.Made().Eng

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

		d.vPoT.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTPs on %v\n", data.PoT, strPoTPort)

		//   - Run Https Server (SSL)
		rErr = r.RunTLS(strPoTPort, PoTHsslCertFile, PoTHsslKeysFile)

	} else {
		d.vPoT.Fprintf(gin.DefaultWriter, "[%v] Listening and serving HTTP on %v\n", data.PoT, strPoTPort)

		//   - Run Http Server
		rErr = r.Run(strPoTPort)
	}

	if rErr != nil { panic(rErr) }
}

func (d *PoT) PoT() *PoT {
	// Routes Initialized
	R.RPoT = d.PoTRoute

	// Err Modules Initialize
	//   - Combine Error Message of Self Define
	E.EPoT = d.PoTError
	E.EPoT.ErrPoTCombine()

	return d
}


