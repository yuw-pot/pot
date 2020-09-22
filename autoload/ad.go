// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/spf13/pflag"
	"github.com/yuw-pot/pot/data"
	A "github.com/yuw-pot/pot/modules/adapter"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	U "github.com/yuw-pot/pot/modules/utils"
)

type autoload struct {
	prop *P.PoT
	Us *U.PoT
}

func init() {
	ad := ad()

	// Initialized Properties
	//   - assign the properties.PropertyPoT
	ad.property()
	if P.PropertyPoT == nil {
		panic(E.Err(data.ErrPfx, "AdPropVar"))
	}

	var adPowerPoT *data.PowerPoT
	_ = P.PropertyPoT.UsK("Power", &adPowerPoT)
	if adPowerPoT == nil {
		panic(E.Err(data.ErrPfx, "PoTPowerErr"))
	}

	// Initialized Adapter
	if adPowerPoT.Adapter == 1 {
		A.New().Made()
	}

	if adPowerPoT.Redis == 1 {
		ad.redis()
	}
}

func ad() *autoload {
	return &autoload {
		prop: P.New(),
		Us: U.New(),
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

func (ad *autoload) redis() {

}
