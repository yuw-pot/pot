// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapter

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"sync"
	"time"
)

type (
	dbPoT struct {
		engine *xorm.Engine
		mx *sync.Mutex
		dn interface{}
		name string
		sParam *srcParam
		sConns *srcConns
		timeLocation *time.Location
	}
)

func dbNew() *dbPoT {
	return &dbPoT {
		mx: &sync.Mutex{},
	}
}

func (db *dbPoT) xEngine() *xorm.Engine {
	if db.engine == nil {
		panic(E.Err(data.ErrPfx, "AdapterEngineErr"))
	}

	return db.engine
}

func (db *dbPoT) instance() {
	db.mx.Lock()

	defer func() {
		db.mx.Unlock()
	}()

	if db.engine != nil {
		return
	}

	if db.sConns == nil || db.sParam == nil {
		panic(E.Err(data.ErrPfx, "AdapterSrc"))
	}

	driverFormat := "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	driverSource := fmt.Sprintf(
		driverFormat,
		db.sConns.Username,
		db.sConns.Password,
		db.sConns.Host,
		db.sConns.Port,
		db.name,
	)

	engine, err := xorm.NewEngine(cast.ToString(db.dn), driverSource)
	if err != nil {
		if engine != nil {
			engine.Close()
		}

		panic(err)
	}

	engine.SetMaxOpenConns(db.sParam.MaxOpen)
	engine.SetMaxIdleConns(db.sParam.MaxIdle)
	engine.SetConnMaxLifetime(time.Second)

	engine.ShowSQL(db.sParam.ShowedSQL)
	engine.SetTZDatabase(db.timeLocation)

	if db.sParam.CachedSQL {
		cached := xorm.NewLRUCacher(xorm.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cached)
	}

	db.engine = engine
	return
}
