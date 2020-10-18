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

		vs *utils.PoT
		fs *files.PoT
	}
)

func New() *PoT {
	return &PoT {
		Prop: viper.New(),

		vs: utils.New(),
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

	if ok := cfg.vs.Contains(env, data.PropertySfxs); ok == false {
		panic(E.Err(data.ErrPfx, "PropEnvExclude"))
	}

	dir := "./." + env + "." + data.PropertySfx

	ok, _ := cfg.fs.IsExists(dir)
	if ok == false {
		f, err := cfg.fs.Create(dir)
		defer func() {
			f.Close()
		}()

		if err != nil {
			panic(E.Err(data.ErrPfx, "PropEnvFile"))
		}

		f.WriteString(cfg.tpl())
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

func (cfg *PoT) GeT(k string, v interface{}) interface{} {
	if cfg.Prop.IsSet(k) {
		return cfg.Prop.Get(k)
	}

	return v
}

func (cfg *PoT) UsK(k string, v interface{}, opts ...viper.DecoderConfigOption) error {
	return cfg.Prop.UnmarshalKey(k, v, opts ...)
}

func (cfg *PoT) tpl() string {
	return cfg.vs.Sprintf(`## %v ##
PoT:
  Name: "%v"
  Port: "%v"
  Mode: %d
  TimeLocation: "%v"
  Addr: ""
  Hssl:
    Power: %d
    CertFile: ""
    KeysFile: ""

Power:
  Adapter: %d
  Redis: %d

Adapter:
  Mysql:
    Param:
      MaxOpen: 2000
      MaxIdle: 1000
      ShowedSQL: false
      CachedSQL: false

    Conns:
      db_demo:
        Master:
          - Host: "127.0.0.1"
            Port: 3306
            Username: "root"
            Password: "root"
          - Host: "127.0.0.1"
            Port: 3306
            Username: "root"
            Password: "root"

        Slaver:
          - Host: "127.0.0.1"
            Port: 3306
            Username: "root"
            Password: "root"

Redis:
  I:
    Network: "tcp"
    Addr: "127.0.0.1:6397"
    Password: ""
    DB: 0
  II:
    Network: "tcp"
    Addr: "127.0.0.1:6397"
    Password: ""
    DB: 1

RedisCluster:
  Network: "tcp"
  Password: ""
  DB: 0
  AddrCluster:
    - "127.0.0.1:6397"
    - "127.0.0.1:6397"
    - "127.0.0.1:6397"

Logs:
  PoT:
    FileName: "%v"
    ZapCoreLevel: "info"
    Format: "%v"
    Time: "%v"
    Level: "%v"
    Name: "%v"
    Caller: "%v"
    Message: "%v"
    StackTrace: "%v"
    MaxSize: %d
    MaxBackUps: %d
    MaxAge: %d

PoTSelfDefined:
  Test: "Configure Self Defined Success"
`,
data.PoT,
data.PoT,
// - PoT
data.PropertyPort,
data.PropertyMode,
data.PropertyTimeLocation,
data.PropertyHsslPower,
data.PropertyAdapter,
data.PropertyRedis,
// - Log
data.LogFileName,
data.LogFormatConsole,
data.LogTime,
data.LogLevel,
data.LogName,
data.LogCaller,
data.LogMessage,
data.LogStackTrace,
data.LogMaxSize,
data.LogMaxBackups,
data.LogMaxAge,
	)
}
