// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/spf13/pflag"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/adapter"
	"github.com/yuw-pot/pot/modules/auth"
	"github.com/yuw-pot/pot/modules/cache/redis"
	"github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/properties"
)

type autoload struct {
	prop *properties.PoT
}

func init() {
	ad := ad()

	// Initialized Properties
	//   - assign the properties.PropertyPoT
	ad.property()

	if properties.PropertyPoT == nil {
		panic(err.Err(data.ErrPfx, "AdPropVar"))
	}

	var adPowerPoT *data.PowerPoT
	_ = properties.PropertyPoT.UsK("Power", &adPowerPoT)
	if adPowerPoT == nil {
		panic(err.Err(data.ErrPfx, "PoTPowerErr"))
	}

	//   - add JwT Key
	if adPowerPoT.JwT == 1 {
		_ = properties.PropertyPoT.UsK("JwT", &auth.JPoT)
		if auth.JPoT == nil {
			panic(err.Err(data.ErrPfx, "PoTJwTErr"))
		}
	}

	// Initialized Adapter
	if adPowerPoT.Adapter == 1 {
		adapter.New().Made()
	}

	// Initialized Redis
	if adPowerPoT.Redis == 1 {
		redis.New().Made()
	}
}

func ad() *autoload {
	return &autoload {
		prop: properties.New(),
	}
}

func (ad *autoload) property() {
	pflag.String("env", "", "environment configure")
	pflag.Parse()

	if errProperties := ad.prop.Prop.BindPFlags(pflag.CommandLine); errProperties != nil {
		panic(err.Err(data.ErrPfx, "AdPropBind"))
	}

	properties.PropertyPoT = ad.prop.Load()
}
