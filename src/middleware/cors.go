// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
)

func (m *M) Cors() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")

		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token,Authorization")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			ctx.AbortWithStatus(data.PoTStatusNoContent)
		}

		ctx.Next()
	}
}