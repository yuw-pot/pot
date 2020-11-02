// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package middleware

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/libs"
	"github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
)

func (m *M) JwTAuth() *libs.PoT {
	return m.lib.SeT(func(p *libs.PoT) {
		ctx := p.Lib()

		token := ctx.Request.Header.Get("token")
		if token == "" {
			ctx.JSON(data.PoTStatusOK, &data.SrvPoT{
				Status:   data.PoTUnKnown,
				Msg:      E.Err(data.ErrPfx, "MWareNoPriority").Error(),
				Response: nil,
			})
			ctx.Abort()
			return
		}

		JwT, err := crypto.JPoT.Parse(token)
		if err != nil {
			if err == E.Err(data.ErrPfx, "TokenExpired") {
				ctx.JSON(data.PoTStatusOK, &data.SrvPoT{
					Status:   data.PoTUnKnown,
					Msg:      E.Err(data.ErrPfx, "MWareUnknown").Error(),
					Response: nil,
				})
				ctx.Abort()
				return
			} else {
				ctx.JSON(data.PoTStatusOK, &data.SrvPoT{
					Status:   data.PoTUnKnown,
					Msg:      err.Error(),
					Response: nil,
				})
				ctx.Abort()
				return
			}
		}

		JwTRefresh, _ := crypto.JPoT.Refresh(token)

		ctx.Set("JwT", JwTRefresh)
		ctx.Set("JwTInfo", JwT)
		ctx.Next()
	})
}
