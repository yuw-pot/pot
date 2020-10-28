// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/spf13/pflag"
	"github.com/yuw-pot/pot/data"
	A "github.com/yuw-pot/pot/modules/adapter"
	R "github.com/yuw-pot/pot/modules/cache/redis"
	C "github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	U "github.com/yuw-pot/pot/modules/utils"
)

type autoload struct {
	prop *P.PoT
	vPoT *U.PoT
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

	//   - add JwT Key
	if adPowerPoT.JwT == 1 {
		_ = P.PropertyPoT.UsK("JwT", &C.JPoT)
		if C.JPoT == nil {
			panic(E.Err(data.ErrPfx, "PoTJwTErr"))
		}
	}

	// Initialized Adapter
	if adPowerPoT.Adapter == 1 {
		A.New().Made()
	}

	// Initialized Redis
	if adPowerPoT.Redis == 1 {
		R.New().Made()
	}
}

func ad() *autoload {
	return &autoload {
		prop: P.New(),
		vPoT: U.New(),
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
