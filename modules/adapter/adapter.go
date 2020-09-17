// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapter

import (
	"github.com/go-xorm/xorm"
	"github.com/yuw-pot/pot/data"
	Z "github.com/yuw-pot/pot/modules/zlog"
	"time"
)

const (
	adapterMysql string = "mysql"
	adapterPostgreSQL string = "postgres"
)

type (
	adapter interface {
		Engine() *xorm.Engine
	}

	PoT struct {
		dn string
	}
)

var (
	_ adapter = &PoTMysql{}

	sysTimeLocation *time.Location
)

func init() {
	//_ = P.PropertyPoT.Prop.UnmarshalKey("Logs.Adapter", &zLogPoT)
	//if zLogPoT == nil {
	//	panic(E.Err(data.ErrPfx, "PoTZapLogErr"))
	//}
	//
	//// - Zap Log Install&Made
	//zLog = Z.New(zLogPoT).Made()
}

func New(dn string) *PoT {
	return &PoT {dn:dn}
}

func (adapter *PoT) Engine() *xorm.Engine {
	switch adapter.dn {
	case adapterMysql:
		return nil
	case adapterPostgreSQL:
		return nil
	default:
		panic("")
	}
}




