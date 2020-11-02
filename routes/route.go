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
)

var (
	RPoT *PoT = new(PoT)
	rMaP *data.H = &data.H{}
)

type (
	Routes interface {
		Tag() string
		PuT(r *PoT, toFunc map[*KeY][]interface{})
	}

	RouteSrc []Routes
	RouteArr map[string]map[*KeY][]interface{}

	PoT struct {
		Eng *gin.Engine
		Src *RouteSrc
		Arr *RouteArr
	}

	KeY struct {
		Service 	string
		Controller 	string
		Action 		string
		Mode 		string
		Path 		string
	}
)

func (route *PoT) Made() *PoT {
	var exp *exceptions.PoT = exceptions.New()

	/**
	 * Todo: No Routes To Redirect
	 */
	route.Eng.NoRoute(exp.NoRoute)

	/**
	 * Todo: No Method To Redirect
	 */
	route.Eng.NoMethod(exp.NoMethod)

	for _, to := range *route.Src {
		if _, ok := (*route.Arr)[to.Tag()]; ok == false {
			continue
		}

		if len((*route.Arr)[to.Tag()]) == 0 {
			continue
		}

		to.PuT(route, (*route.Arr)[to.Tag()])
	}

	return route
}

func To(ctx *gin.RouterGroup, toFunc map[*KeY][]interface{}) {
	for x, ctrlHandlerFunc := range toFunc {
		(*rMaP)[x.Service + x.Controller + x.Action] = x.Path

		ctrl := make([]gin.HandlerFunc, len(ctrlHandlerFunc))
		for k, val := range ctrlHandlerFunc {
			ctrl[k] = val.(gin.HandlerFunc)
		}

		switch x.Mode {
		case PoTMethodAnY:
			ctx.Any (x.Path, ctrl ...)
			continue

		case PoTMethodGeT:
			ctx.GET (x.Path, ctrl ...)
			continue

		case PoTMethodPosT:
			ctx.POST(x.Path, ctrl ...)
			continue

		case PoTMethodDelete:
			ctx.DELETE(x.Path, ctrl ...)
			continue

		case PoTMethodPuT:
			ctx.PUT(x.Path, ctrl ...)
			continue

		case PoTMethodHeaD:
			ctx.HEAD(x.Path, ctrl ...)
			continue

		case PoTMethodPatch:
			ctx.PATCH(x.Path, ctrl ...)
			continue

		case PoTMethodOptions:
			ctx.OPTIONS(x.Path, ctrl ...)
			continue

		default:
			continue
		}
	}
}

func GeTPath(service string, controller string, action string) interface{} {
	_, ok := (*rMaP)[service + controller + action]
	if ok == false {
		return nil
	}

	return (*rMaP)[service + controller + action]
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
