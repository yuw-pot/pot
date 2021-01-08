// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package autoload

import (
	"github.com/spf13/pflag"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/adapter"
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
	//   - Adapter
	//   - Redis

	ad.property()
	ad.adapter()
	ad.cache()
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

	if properties.PropertyPoT == nil {
		panic(err.Err(data.ErrPfx, "AdPropVar"))
	}
}

func (ad *autoload) adapter() { adapter.Initialized() }

func (ad *autoload) cache() { redis.Initialized() }
