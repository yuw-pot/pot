// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/exceptions"
	"net/http"
	"strings"
)

var (
	RPoT *PoT = new(PoT)
	RMaP *data.H = &data.H{}
)

type (
	Routes interface {
		Tag() string
		Put(r *gin.Engine, toFunc map[string][]gin.HandlerFunc)
	}

	RouteSrc []Routes
	RouteTpl []interface{}
	RouteArr map[string]map[string][]gin.HandlerFunc

	PoT struct {
		Src *RouteSrc
		Tpl *RouteTpl
		Arr *RouteArr
	}
)

func (route *PoT) Made(r *gin.Engine) {
	var exp *exceptions.PoT = exceptions.New()

	/**
	 * Todo: No Routes To Redirect
	 */
	r.NoRoute(exp.NoRoute)

	/**
	 * Todo: No Method To Redirect
	 */
	r.NoMethod(exp.NoMethod)

	for _, to := range *route.Src {
		if _, ok := (*route.Arr)[to.Tag()]; ok == false {
			continue
		}

		if len((*route.Arr)[to.Tag()]) == 0 {
			continue
		}

		to.Put(r, (*route.Arr)[to.Tag()])
	}
}

func To(g *gin.RouterGroup, toFunc map[string][]gin.HandlerFunc) {
	for r, ctrl := range toFunc {
		x := strings.Split(r, data.RSeP)

		if len(x) != 3 {
			continue
		}

		(*RMaP)[x[0]] = x[2]

		switch strings.ToLower(x[1]) {
		case "get":
			g.GET (x[2], ctrl ...)
			continue

		case "any":
			g.Any (x[2], ctrl ...)
			continue

		case "post":
			g.POST(x[2], ctrl ...)
			continue

		default:
			continue
		}
	}
}

func Cors() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		method := ctx.Request.Method
		origin := ctx.Request.Header.Get("Origin")

		if origin != "" {
			ctx.Header("Access-Control-Allow-Origin", "*")
			ctx.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
			ctx.Header("Access-Control-Allow-Headers", "Content-Type,AccessToken,X-CSRF-Token, Authorization")
			ctx.Header("Access-Control-Allow-Credentials", "true")
			ctx.Set("content-type", "application/json")
		}

		if method == "OPTIONS" {
			ctx.AbortWithStatus(http.StatusNoContent)
		}

		ctx.Next()
	}
}

func LoggerWithFormat() gin.HandlerFunc {
	return gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("[PoT] %v |	%v |	%v |	%v |	%v |	%v(%v)\n",
			param.TimeStamp.Format("2006/01/02 - 15:04:05"),
			param.StatusCode,
			param.Latency,
			param.ClientIP,
			param.Method,
			param.Path,
			param.Request.Proto,
		)
	})
}
