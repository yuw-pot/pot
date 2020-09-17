// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapter

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"sync"
	"time"
)

type (
	PoTMysql struct {
		engine *xorm.Engine
		mx *sync.Mutex
		SrcParam *SrcParam
		SrcConns *SrcConns
	}

	SrcParam struct {
		MaxOpen int
		MaxIdle int
		ShowedSQL bool
		CachedSQL bool
	}

	SrcConns struct {
		Host string
		Port int
		Name string
		Username string
		Password string
	}
)

func NewMysql() *PoTMysql {
	mysqlPoT := &PoTMysql {}
	mysqlPoT.mx = &sync.Mutex{}


	return mysqlPoT
}

func (db *PoTMysql) Engine() *xorm.Engine {
	if db.engine == nil {

	}

	return db.engine
}

func (db *PoTMysql) instance() *PoTMysql {
	db.mx.Lock()
	defer db.mx.Unlock()

	if db.engine != nil {
		return db
	}

	if db.SrcConns == nil || db.SrcParam == nil {
		panic(E.Err(data.ErrPfx, "AdapterSrc"))
	}

	driverFormat := "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	driverSource := fmt.Sprintf(
		driverFormat,
		db.SrcConns.Username,
		db.SrcConns.Password,
		db.SrcConns.Host,
		db.SrcConns.Port,
		db.SrcConns.Name,
	)

	engine, err := xorm.NewEngine(adapterMysql, driverSource)
	if err != nil {
		if engine != nil {
			engine.Close()
		}

		panic(E.Err(data.ErrPfx, "AdapterEngineErr"))
	}

	engine.SetMaxOpenConns(db.SrcParam.MaxOpen)
	engine.SetMaxIdleConns(db.SrcParam.MaxIdle)
	engine.SetConnMaxLifetime(time.Second)

	engine.ShowSQL(db.SrcParam.ShowedSQL)
	engine.SetTZDatabase(sysTimeLocation)

	if db.SrcParam.CachedSQL {
		cached := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cached)
	}

	db.engine = engine

	return db
}
