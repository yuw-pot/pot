// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapter

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"sync"
	"time"
	"xorm.io/xorm"
	"xorm.io/xorm/caches"
)

type (
	odbc struct {
		engine *xorm.Engine
		mx *sync.Mutex
		dn interface{}
		sParam *srcParam
		sConns *srcConns
		timeLocation *time.Location
	}
)

func newODBC() *odbc {
	return &odbc {
		mx: &sync.Mutex{},
	}
}

func (db *odbc) xEngine() *xorm.Engine {
	if db.engine == nil {
		panic(E.Err(data.ErrPfx, "AdapterEngineErr"))
	}

	return db.engine
}

func (db *odbc) instance() *odbc {
	db.mx.Lock()
	defer func() { db.mx.Unlock() }()

	if db.engine != nil { return db }

	if db.sConns == nil || db.sParam == nil {
		panic(E.Err(data.ErrPfx, "AdapterSrc"))
	}

	var connsDriver string = "%s:%s@tcp(%s:%d)/%s?charset=utf8"
	var conns string = fmt.Sprintf(
		connsDriver,
		db.sConns.Username,
		db.sConns.Password,
		db.sConns.Host,
		db.sConns.Port,
		db.sConns.Repo,
	)

	engine, err := xorm.NewEngine(cast.ToString(db.dn), conns)
	if err != nil { panic(err) }

	engine.SetMaxOpenConns(db.sParam.MaxOpen) 	// 设置最大打开连接数
	engine.SetMaxIdleConns(db.sParam.MaxIdle) 	// 设置连接池的空闲数大小
	engine.SetConnMaxLifetime(time.Minute) 		// 设置连接的最大生存时间

	engine.ShowSQL(db.sParam.ShowedSQL)
	engine.SetTZDatabase(db.timeLocation)

	if db.sParam.CachedSQL {
		cache := caches.NewLRUCacher(caches.NewMemoryStore(), 1000)
		engine.SetDefaultCacher(cache)
	}

	db.engine = engine
	return db
}
