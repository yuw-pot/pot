// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/exceptions"
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
