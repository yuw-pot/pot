// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package pot

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	R "github.com/yuw-pot/pot/routes"

	_ "github.com/yuw-pot/pot/autoload"
)

const (
	Version string = "v1.0.0"

	ginMode string = gin.DebugMode
	ginPort string = "8888"
)

type (
	PoT struct {
		PoTRoute *R.PoT
		PoTError *E.PoT
	}
)

func init() {}

func New() *PoT {
	return &PoT {}
}

func (engine *PoT) Run() {
	gin.DisableConsoleColor()

	/* Todo Set Mode by YamlConfigure */
	gin.SetMode(data.PoTMode[cast.ToInt(P.PropertyPoT.Get("PoT.Mode", 1))])
	fmt.Println(P.PropertyPoT.Get("PoT.Mode", 1))
	fmt.Println(data.PoTMode[cast.ToInt(P.PropertyPoT.Get("PoT.Mode", 1))])

	r := gin.New()
	r.Use(gin.Recovery(), R.LoggerWithFormat())

	R.RPoT.Made(r)
	r.Run(":" + cast.ToString(P.PropertyPoT.Get("PoT.Port", ginPort)))
}

func (engine *PoT) PoT() *PoT {
	/* Todo: Routes Initialized */
	R.RPoT = engine.PoTRoute

	/* Todo: Err Modules Initialized */
	E.EPoT = engine.PoTError
	E.EPoT.ErrPoTCombine()

	return engine
}


