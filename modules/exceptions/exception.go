// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package exceptions

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/err"
	"net/http"
)

type Exceptions struct {}

func New() *Exceptions {
	return &Exceptions {}
}

func (exp *Exceptions) NoRoute(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusNotFound, err.Err(data.ErrPfx, "ErrDefault"))
	return
}

func (exp *Exceptions) NoMethod(ctx *gin.Context) {
	ctx.AbortWithError(http.StatusNotFound, err.Err(data.ErrPfx, "ErrDefault"))
	return
}

