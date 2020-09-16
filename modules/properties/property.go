// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package properties

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/files"
	"github.com/yuw-pot/pot/modules/utils"
)

var PropertyPoT *PoT

type (
	PoT struct {
		Prop *viper.Viper

		us *utils.Utils
		fs *files.Files
	}
)

func New() *PoT {
	return &PoT {
		Prop: viper.New(),

		us: utils.New(),
		fs: files.New(),
	}
}

func (cfg *PoT) Load() *PoT {
	cfg.Prop.AddConfigPath(".")
	cfg.Prop.SetConfigType(data.PropertySfx)

	env := cfg.Prop.GetString("env")
	if env == "" {
		panic(E.Err(data.ErrPfx, "PropEnvEmpty"))
	}

	if ok := cfg.us.Contains(env, data.PropertySfxs); ok == false {
		panic(E.Err(data.ErrPfx, "PropEnvExclude"))
	}

	dir := "./." + env + "." + data.PropertySfx

	ok, _ := cfg.fs.IsExists(dir)
	if ok == false {
		_, err := cfg.fs.Create(dir)
		if err != nil {
			panic(E.Err(data.ErrPfx, "PropEnvFile"))
		}
	}

	cfg.Prop.SetConfigName("." + env)
	if err := cfg.Prop.ReadInConfig(); err != nil {
		panic(err)
	}

	cfg.Prop.WatchConfig()
	cfg.Prop.OnConfigChange(func (e fsnotify.Event){
		// Todo: do something ...

	})

	return cfg
}

func (cfg *PoT) Get(k string, v interface{}) interface{} {
	if cfg.Prop.IsSet(k) {
		return cfg.Prop.Get(k)
	}

	return v
}
