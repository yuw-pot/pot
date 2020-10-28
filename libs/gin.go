// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package libs

import "github.com/gin-gonic/gin"

type (
	PoT struct {
		lib HandlerFuncPoT
		ctx *gin.Context
	}

	HandlerFuncPoT func(ctx *PoT)
)

func (g *PoT) SeT(Handler HandlerFuncPoT) *PoT {
	g.lib = Handler
	return g
}

func (g *PoT) Lib() *gin.Context {
	return g.ctx
}

func (g *PoT) PoT() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		g.ctx = ctx
		g.lib(g)
	}
}
