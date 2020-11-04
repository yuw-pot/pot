// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package models

import (
	"github.com/go-xorm/xorm"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"strings"
)

type Models struct {
	engine *xorm.Engine
}

func New(engine *xorm.Engine) *Models {
	return &Models {
		engine: engine,
	}
}

func (m *Models) Engine() *xorm.Engine {
	return m.engine
}

func (m *Models) Insert(d interface{}) (i int64, err error) {
	db := m.engine.NewSession()

	defer func() {
		db.Close()
	}()

	return db.Insert(d)
}

func (m *Models) Total(d interface{}) (i int64, err error) {
	db := m.engine.NewSession()

	defer func() {
		db.Close()
	}()

	return db.Count(d)
}

func (m *Models) Update(mPoT *data.ModPoT, d interface{}) (i int64, err error) {
	if mPoT.Query == nil || mPoT.QueryArgs == nil {
		return 0, E.Err(data.ErrPfx, "ModParamsErr")
	}

	db := m.engine.NewSession()

	defer func() {
		db.Close()
	}()

	return db.Where(mPoT.Query, mPoT.QueryArgs ...).Update(d)
}

func (m *Models) Delete(mPoT *data.ModPoT, d interface{}) (i int64, err error) {
	if mPoT.Query == nil || mPoT.QueryArgs == nil {
		return 0, E.Err(data.ErrPfx, "ModParamsErr")
	}

	db := m.engine.NewSession()

	defer func() {
		db.Close()
	}()

	return db.Where(mPoT.Query, mPoT.QueryArgs ...).Delete(d)
}

func (m *Models) GeT(mPoT *data.ModPoT, d interface{}) (bool, error) {
	db := m.engine.NewSession()

	defer func() {
		db.Close()
	}()

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