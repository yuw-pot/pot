// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pot

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/properties"
	"github.com/yuw-pot/pot/modules/subscriber"
	"github.com/yuw-pot/pot/modules/utils"
	"github.com/yuw-pot/pot/routes"

	_ "github.com/yuw-pot/pot/autoload"
)

const version string = "v1.0.0"

type (
	PoT struct {
		vPoT *utils.PoT

		PoTRoute *routes.PoT
		PoTError *err.PoT
		PoTSubscriber *subscriber.PoT
	}
)

func New() *PoT {
	return &PoT {
		vPoT: utils.New(),
	}
}

func (d *PoT) Run() {
	d.vPoT.Fprintf(gin.DefaultWriter, "[%v] %v\n", data.PoT, version)

	// Disable Console Color
	gin.DisableConsoleColor()

	// Gin Mode Release
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	PoTMode := properties.PropertyPoT.GeT("PoT.Mode", "1")
	if d.vPoT.Contains(PoTMode, 0,1) == false {
		PoTMode = 1
	}

	d.setMode(r, data.PoTMode[cast.ToInt(PoTMode)])

	routes.RPoT.Made(r)

	// Https Power ON/OFF
	//   - PoT.Hssl
	//     - PoT.Hssl.Power
	//     - PoT.Hssl.CertFile
	//     - PoT.Hssl.KeysFile
	var rErr error
	strPoTPort := ":"+cast.ToString(properties.PropertyPoT.GeT("PoT.Port", data.PropertyPort))

	if properties.PropertyPoT.GeT("PoT.Hssl.Power", 0) == 1 {
		PoTHsslCertFile := cast.ToString(properties.PropertyPoT.GeT("PoT.Hssl.CertFile", ""))
		if PoTHsslCertFile == "" {
			panic(err.Err(data.ErrPfx, "PoTSslCF"))
		}

		PoTHsslKeysFile := cast.ToString(properties.PropertyPoT.GeT("PoT.Hssl.KeysFile", ""))
		if PoTHsslKeysFile == "" {
			panic(err.Err(data.ErrPfx, "PoTSslKF"))
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
	routes.RPoT = d.PoTRoute

	// Err Modules Initialize
	//   - Combine Error Message of Self Define
	err.EPoT = d.PoTError
	err.EPoT.ErrPoTCombine()

	// Subscriber Initialize
	if d.PoTSubscriber != nil {
		subscriber.StarT(d.PoTSubscriber)
	}

	return d
}


