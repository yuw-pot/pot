// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/spf13/pflag"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
)

type autoload struct {
	prop *P.PoT
}

func init() {
	ad := ad()

	// Initialized Properties
	// - assign the properties.PropertyPoT
	ad.property()
	if P.PropertyPoT == nil {
		panic(E.Err(data.ErrPfx, "AdPropVar"))
	}
}

func ad() *autoload {
	return &autoload {
		prop: P.New(),
	}
}

func (ad *autoload) property() {
	pflag.String("env", "", "environment configure")
	pflag.Parse()

	if err := ad.prop.Prop.BindPFlags(pflag.CommandLine); err != nil {
		panic(E.Err(data.ErrPfx, "AdPropBind"))
	}

	P.PropertyPoT = ad.prop.Load()
}

func (ad *autoload) adapter() {

}
