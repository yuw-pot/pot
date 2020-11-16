// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"database/sql"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/adapter"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"strings"
	"xorm.io/xorm"
)

type Models struct {
	v *utils.PoT
	group map[interface{}][]*xorm.Engine
}

func New(group map[interface{}][]*xorm.Engine) *Models {
	return &Models {
		v: utils.New(),
		group: group,
	}
}

func (m *Models) GeTEngine(method string) (*xorm.Engine, error) {
	if ok := m.v.Contains(method, adapter.Master, adapter.Slaver); ok == false {
		return nil, E.Err(data.ErrPfx, "AdapterModeName")
	}

	var adapterLength int = 0

	switch method {
	case adapter.Master:

		adapterLength = len(m.group[adapter.Master])

		_, ok := m.group[adapter.Master]
		if ok == false || adapterLength != 1 {
			return nil, E.Err(data.ErrPfx, "AdapterEngineGroupErr")
		}

		return m.group[adapter.Master][0], nil

	case adapter.Slaver:

		adapterLength = len(m.group[adapter.Slaver])

		_, ok := m.group[adapter.Slaver]
		if ok == false || adapterLength <= 0 {
			return nil, E.Err(data.ErrPfx, "AdapterEngineGroupErr")
		}

		if adapterLength == 1 {
			return m.group[adapter.Slaver][0], nil
		} else {
			return m.group[adapter.Slaver][m.v.NumRandom(adapterLength)], nil
		}

	default:
		return nil, E.Err(data.ErrPfx, "AdapterModeName")
	}
}

func (m *Models) Transaction(operation func(db *xorm.Session)) error {
	conn, err := m.GeTEngine(adapter.Master)
	if err != nil { return err  }

	db := conn.NewSession()
	defer func() { db.Close() }()

	err = db.Begin()
	if err != nil { return err }

	operation(db)

	err = db.Commit()
	if err != nil {
		_ = db.Rollback()
		return err
	}

	return nil
}

func (m *Models) Exec(sqlOrArgs ...interface{}) (sql.Result, error) {
	conn, err := m.GeTEngine(adapter.Master)
	if err != nil { return nil, err }

	return conn.Exec(sqlOrArgs ...)
}

func (m *Models) Query(sqlOrArgs ...interface{}) ([]map[string][]byte, error) {
	conn, err := m.GeTEngine(adapter.Slaver)
	if err != nil { return nil, err }

	return conn.Query(sqlOrArgs ...)
}

func (m *Models) Insert(d interface{}) (i int64, err error) {
	conn, err := m.GeTEngine(adapter.Master)
	if err != nil { return 0, err }

	db := conn.NewSession()
	defer func() { db.Close() }()

	i, err = db.Insert(d)
	//conn.ClearCache()

	return
}

func (m *Models) Update(mPoT *data.ModPoT, d interface{}) (i int64, err error) {
	conn, err := m.GeTEngine(adapter.Master)
	if err != nil { return 0, err }

	db := conn.NewSession()
	defer func() { db.Close() }()

	return db.Where(mPoT.Query, mPoT.QueryArgs ...).Update(d)
}

func (m *Models) Delete(mPoT *data.ModPoT, d interface{}) (i int64, err error) {
	if mPoT.Query == nil || mPoT.QueryArgs == nil {
		return 0, E.Err(data.ErrPfx, "ModParamsErr")
	}

	conn, err := m.GeTEngine(adapter.Master)
	if err != nil { return 0, err }

	db := conn.NewSession()
	defer func() { db.Close() }()

	return db.Where(mPoT.Query, mPoT.QueryArgs ...).Delete(d)
}

func (m *Models) Total(d interface{}) (i int64, err error) {
	conn, err := m.GeTEngine(adapter.Slaver)
	if err != nil { return 0, err }

	db := conn.NewSession()
	defer func() { db.Close() }()

	return db.Count(d)
}

func (m *Models) GeT(mPoT *data.ModPoT, d interface{}) (bool, error) {
	conn, err := m.GeTEngine(adapter.Slaver)
	if err != nil { return false, err }

	db := conn.NewSession()
	defer func() { db.Close() }()

	if mPoT.Table != "" && mPoT.Field != nil {
		if ok, _ := db.IsTableExist(mPoT.Table); ok == false {
			return false, E.Err(data.ErrPfx, "ModDBTable")
		}

		db = db.Table(mPoT.Table).Select(strings.Join(mPoT.Field,","))
	} else {
		if mPoT.Columns != nil {
			db = db.Cols(mPoT.Columns ...)
		} else {
			db = db.Cols()
		}
	}

	// Add Join Table
	//   - INNER
	//   - LEFT
	//   - RIGHT
	if mPoT.Joins != nil {
		for _, join := range mPoT.Joins {
			db = db.Join(join.JoinOperator, join.TableName, join.Condition)
		}
	}

	// Add Condition
	if mPoT.Query != nil && mPoT.QueryArgs != nil {
		db = db.Where(mPoT.Query, mPoT.QueryArgs ...)
	}

	switch mPoT.Types {

	case data.ModONE:
		return db.Get(d)

	case data.ModALL:

		// Add OrderBy
		//   - ASC
		//   - DESC
		if len(mPoT.OrderArgs) > 0 {
			if mPoT.OrderType == data.ModByAsc {
				db = db.Asc(mPoT.OrderArgs ...)
			}

			if mPoT.OrderType == data.ModByEsc {
				db = db.Desc(mPoT.OrderArgs ...)
			}
		}

		//   - Limit
		//   - Start
		if len(mPoT.Start) > 0 && mPoT.Limit != 0 {
			db = db.Limit(mPoT.Limit, mPoT.Start ...)
		}

		return true, db.Find(d)

	default:
		return false, E.Err(data.ErrPfx, "ModDBSelectErr")
	}
}