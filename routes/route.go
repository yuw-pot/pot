// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/crypto"
	"github.com/yuw-pot/pot/modules/exceptions"
)

var (
	RPoT *PoT = new(PoT)
	rMaP *data.H = &data.H{}
)

type (
	Routes interface {
		Tag() string
		PuT(r *gin.Engine, toFunc map[*KeY][]gin.HandlerFunc)
	}

	RouteSrc []Routes
	RouteArr map[string]map[*KeY][]gin.HandlerFunc

	PoT struct {
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

		to.PuT(r, (*route.Arr)[to.Tag()])
	}
}

func To(ctx *gin.RouterGroup, toFunc map[*KeY][]gin.HandlerFunc) {
	for x, ctrl := range toFunc {
		cry := crypto.New()

		cry.Mode = data.ModeToken
		cry.D = []interface{}{data.MD5, x.Service + x.Controller + x.Action}

		k, err := cry.Made()
		if err != nil {
			panic(err)
		}

		(*rMaP)[cast.ToString(k)] = x.Path

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
