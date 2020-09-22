// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapter

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"

	//E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	U "github.com/yuw-pot/pot/modules/utils"
	"strings"
	"time"
)

const (
	Master		string = "master"
	Slaver		string = "slaver"

	Mysql      	string = "mysql"
	PostgreSQL 	string = "postgres"
)

type (
	PoT struct {
		Us *U.PoT
		dn interface{}
		name string
		sConns *srcConns
		sParam *srcParam
		timeLocation *time.Location
	}

	srcParam struct {
		MaxOpen int
		MaxIdle int
		ShowedSQL bool
		CachedSQL bool
	}

	srcConns struct {
		Host string
		Port int
		Username string
		Password string
	}
)

var (
	dModes []interface{} = []interface{}{
		Master,
		Slaver,
	}

	dNames []interface{} = []interface{}{
		Mysql,
		PostgreSQL,
	}

	//  - map[DriverName][dbName][]*xorm.Engine
	adapterMaster map[interface{}]map[interface{}][]*xorm.Engine = map[interface{}]map[interface{}][]*xorm.Engine{
		Mysql:      {},
		PostgreSQL: {},
	}

	adapterSlaver map[interface{}]map[interface{}][]*xorm.Engine = map[interface{}]map[interface{}][]*xorm.Engine{
		Mysql:      {},
		PostgreSQL: {},
	}
)

func Made(driverName string, name string, mode ... interface{}) (*xorm.Engine, error) {
	Us := U.New()
	if ok := Us.Contains(driverName, dNames); ok == false {
		return nil, E.Err(data.ErrPfx, "AdapterMadeDN")
	}

	var modeAdapter interface{}

	if len(mode) == 0 {
		modeAdapter = Master
	} else {
		if ok := Us.Contains(mode[0], dModes); ok {
			modeAdapter = mode[0]
		} else {
			return nil, E.Err(data.ErrPfx, "AdapterMadeMode")
		}
	}

	switch modeAdapter {
	case Master:
		_, okd := adapterMaster[driverName]
		if okd == false {
			return nil, E.Err(data.ErrPfx, "AdapterMadeDN")
		}

		_, okn := adapterMaster[driverName][name]
		if okn == false {
			return nil, E.Err(data.ErrPfx, "AdapterMadeName")
		}

		break

	case Slaver:
		_, okd := adapterSlaver[driverName]
		if okd == false {
			return nil, E.Err(data.ErrPfx, "AdapterMadeDN")
		}

		_, okn := adapterSlaver[driverName][name]
		if okn == false {
			return nil, E.Err(data.ErrPfx, "AdapterMadeName")
		}

		break

	default:
		return nil, E.Err(data.ErrPfx, "AdapterMadeMode")
		break
	}

	var (
		mLen int = len(adapterMaster[driverName][name])
		sLen int = len(adapterSlaver[driverName][name])
	)

	if len(mode) == 0 {
		return adapterMaster[driverName][name][Us.NumRandom(0, mLen)], nil
	}

	var x *xorm.Engine

	switch strings.ToLower(cast.ToString(mode[0])) {
	case "master":
		x =  adapterMaster[driverName][name][Us.NumRandom(0, mLen)]
		break

	case "slaver":
		x =  adapterSlaver[driverName][name][Us.NumRandom(0, sLen)]
		break

	default:
		return nil, E.Err(data.ErrPfx, "AdapterMadeMode")
	}

	return x, nil
}

func New() *PoT {
	return &PoT {
		Us: U.New(),
	}
}

func (adapterPoT *PoT) Made() {
	// Set db Time Location
	//   - GeT Configure
	adapterTimeLocation := P.PropertyPoT.GeT("PoT.TimeLocation", data.TimeLocation)
	d, err := adapterPoT.Us.SetTimeLocation(cast.ToString(adapterTimeLocation))
	if err != nil {
		panic(err)
	}

	adapterPoT.timeLocation = d

	//   - SeT Adapter data
	for _, dn := range dNames {
		var sParam *srcParam
		_ = P.PropertyPoT.UsK(
			fmt.Sprintf("Adapter.%v.Param", cast.ToString(dn)),
			&sParam,
		)

		if sParam == nil {
			continue
		}

		adapterPoT.dn = dn
		adapterPoT.sParam = sParam

		adapterConns := P.PropertyPoT.GeT(
			fmt.Sprintf("Adapter.%v.Conns", cast.ToString(dn)),
			nil,
		)

		if adapterConns == nil {
			continue
		}

		for db, conns := range adapterConns.(map[string]interface{}) {
			dbTag := strings.Split(db, "_")

			if len(dbTag) != 2 {
				continue
			}

			if dbTag[1] == "" {
				continue
			}

			adapterPoT.name = dbTag[1]

			master, okMaster := conns.(map[string]interface{})["master"]
			if okMaster == false || master == nil {
				continue
			}

			adapterMaster[dn][adapterPoT.name] = make([]*xorm.Engine, len(master.([]interface{})))
			adapterMaster[dn][adapterPoT.name] = adapterPoT.check(master.([]interface{}) ...)

			slaver, okSlaver := conns.(map[string]interface{})["slaver"]
			if okSlaver == false || slaver == nil {
				continue
			}

			adapterSlaver[dn][adapterPoT.name] = make([]*xorm.Engine, len(slaver.([]interface{})))
			adapterSlaver[dn][adapterPoT.name] = adapterPoT.check(slaver.([]interface{}) ...)
		}
	}
}

func (adapterPoT *PoT) check(d ... interface{}) []*xorm.Engine {
	var adapters []*xorm.Engine = make([]*xorm.Engine, len(d))

	for k, v := range d {
		adapterPoT.sConns = &srcConns {
			Host:     "",
			Port:     0,
			Username: "",
			Password: "",
		}

		if _, ok := v.(map[interface{}]interface{})["Host"]; ok {
			adapterPoT.sConns.Host = cast.ToString(v.(map[interface{}]interface{})["Host"])
		}

		if _, ok := v.(map[interface{}]interface{})["Port"]; ok {
			adapterPoT.sConns.Port = cast.ToInt(v.(map[interface{}]interface{})["Port"])
		}

		if _, ok := v.(map[interface{}]interface{})["Username"]; ok {
			adapterPoT.sConns.Username = cast.ToString(v.(map[interface{}]interface{})["Username"])
		}

		if _, ok := v.(map[interface{}]interface{})["Password"]; ok {
			adapterPoT.sConns.Password = cast.ToString(v.(map[interface{}]interface{})["Password"])
		}

		adapters[k] = adapterPoT.instance()
	}

	return adapters
}

func (adapterPoT *PoT) instance() *xorm.Engine {
	db := dbNew()

	db.dn = adapterPoT.dn
	db.name = adapterPoT.name
	db.sParam = adapterPoT.sParam
	db.sConns = adapterPoT.sConns
	db.timeLocation = adapterPoT.timeLocation

	db.instance()

	return db.xEngine()
}




