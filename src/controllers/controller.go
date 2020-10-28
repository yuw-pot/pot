// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/yuw-pot/pot/libs"
)

type Controller struct {
	lib *libs.PoT
}

func New() *Controller {
	return &Controller {
		lib: new(libs.PoT),
	}
}

func (c *Controller) PoT(HandlerFunc libs.HandlerFuncPoT) gin.HandlerFunc {
	return c.lib.SeT(HandlerFunc).PoT()
}
