// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	J "github.com/yuw-pot/pot/modules/jwt"
)

func (m *M) JwTMemoryAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		TpL := data.TpLInitialized()
		TpL.Status = data.PoTUnKnown

		token := ctx.Request.Header.Get("token")
		if token == "" {
			TpL.Msg = E.Err(data.ErrPfx, "MWareNoPriority").Error()

			ctx.JSON(data.PoTStatusOK, TpL)
			ctx.Abort()
			return
		}

		JwT, err := J.JPoT.Parse(token)
		if err != nil {
			if err == E.Err(data.ErrPfx, "TokenExpired") {
				TpL.Msg = E.Err(data.ErrPfx, "MWareUnknown").Error()

				ctx.JSON(data.PoTStatusOK, TpL)
				ctx.Abort()
				return
			} else {
				TpL.Msg = err.Error()

				ctx.JSON(data.PoTStatusOK, TpL)
				ctx.Abort()
				return
			}
		}

		JwTRefresh, _ := J.JPoT.Refresh(token)

		ctx.Set("JwT", JwTRefresh)
		ctx.Set("JwTInfo", JwT)
		ctx.Next()
	}
}
