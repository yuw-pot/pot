// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package adapter

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/properties"
	"github.com/yuw-pot/pot/modules/utils"
	"strings"
	"time"
	"xorm.io/xorm"
)

const (
	Master		string = "master"
	Slaver		string = "slaver"

	Mysql      	string = "mysql"
	PostgreSQL 	string = "postgres"
)

type (
	PoT struct {
		sConns 			*srcConns
		sParam 			*srcParam
		timeLocation 	*time.Location
		u 				*utils.PoT

		driver 			interface{}
	}

	srcParam struct {
		MaxOpen 	int
		MaxIdle 	int
		ShowedSQL 	bool
		CachedSQL 	bool
	}

	srcConns struct {
		Repo 		string
		Host 		string
		Port 		int
		Username 	string
		Password 	string
	}
)

var (
	dModes []interface{} = []interface{}{ Master, Slaver }
	dNames []interface{} = []interface{}{ Mysql, PostgreSQL }

	//  - map[DriverName][Label][]*xorm.Engine
	adapterMaster map[interface{}]map[interface{}]*xorm.Engine = map[interface{}]map[interface{}]*xorm.Engine{
		Mysql:      {},
		PostgreSQL: {},
	}

	adapterSlaver map[interface{}]map[interface{}][]*xorm.Engine = map[interface{}]map[interface{}][]*xorm.Engine{
		Mysql:      {},
		PostgreSQL: {},
	}
)

func Initialized() {
	newAdapter().initialized()
}

func Conns(driver, label string) (map[interface{}][]*xorm.Engine, error) {
	var ok bool

	_, ok = adapterMaster[driver]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "AdapterMode")
	}

	_, ok = adapterSlaver[driver]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "AdapterMode")
	}

	_, ok = adapterMaster[driver][strings.ToLower(label)]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "AdapterModeName")
	}

	_, ok = adapterSlaver[driver][strings.ToLower(label)]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "AdapterModeName")
	}

	var conns map[interface{}][]*xorm.Engine = map[interface{}][]*xorm.Engine{}

	conns[Master] = make([]*xorm.Engine, 1)
	conns[Master][0] = adapterMaster[driver][strings.ToLower(label)]

	conns[Slaver] = make([]*xorm.Engine, len(adapterSlaver[driver][strings.ToLower(label)]))
	conns[Slaver] = adapterSlaver[driver][strings.ToLower(label)]

	return conns, nil
}

func newAdapter() *PoT {
	return &PoT{ u:utils.New() }
}

func (adapterPoT *PoT) initialized() *PoT {
	// Set db Time Location
	//   - GeT Configure
	adapterTimeLocation := properties.PropertyPoT.GeT("PoT.TimeLocation", data.TimeLocation)
	d, err := adapterPoT.u.SetTimeLocation(cast.ToString(adapterTimeLocation))
	if err != nil { panic(err) }

	adapterPoT.timeLocation = d

	//   - SeT Adapter data
	for _, dn := range dNames {
		var sParam *srcParam
		_ = properties.PropertyPoT.UsK(
			fmt.Sprintf("Adapter.%v.Param", cast.ToString(dn)),
			&sParam,
		)

		if sParam == nil {
			continue
		}

		adapterPoT.driver = dn
		adapterPoT.sParam = sParam

		adapterConns := properties.PropertyPoT.GeT(
			fmt.Sprintf("Adapter.%v.Conns", cast.ToString(dn)),nil,
		)

		if adapterConns == nil {
			continue
		}

		for label, conns := range adapterConns.(map[string]interface{}) {
			var connsTo map[string]interface{} = conns.(map[string]interface{})

			master, okMaster := connsTo["master"]
			if okMaster == false || master == nil {
				content := fmt.Sprintf("Adapter.%v.Conns.%v.Master", dn, label)
				panic(E.Err(data.ErrPfx, "AdapterConfigErr", content))
			}

			slaver, okSlaver := connsTo["slaver"]
			if okSlaver == false || slaver == nil {
				content := fmt.Sprintf("Adapter.%v.Conns.%v.Slaver", dn, label)
				panic(E.Err(data.ErrPfx, "AdapterConfigErr", content))
			}

			masterTo := adapterPoT.u.ToMapInterface(master.(map[string]interface{}))
			masterEngine := adapterPoT.validate(masterTo)
			if len(masterEngine) != 1 {
				content := fmt.Sprintf("Adapter.%v.Conns.%v.Master", dn, label)
				panic(E.Err(data.ErrPfx, "AdapterConfigErr", content))
			}

			adapterMaster[dn][label] = masterEngine[0]

			adapterSlaver[dn][label] = make([]*xorm.Engine, len(slaver.([]interface{})))
			adapterSlaver[dn][label] = adapterPoT.validate(slaver.([]interface{}) ...)
			if len(adapterSlaver[dn][label]) == 0 {
				content := fmt.Sprintf("Adapter.%v.Conns.%v.Slaver", dn, label)
				panic(E.Err(data.ErrPfx, "AdapterConfigErr", content))
			}
		}
	}

	return adapterPoT
}

func (adapterPoT *PoT) validate(d ... interface{}) []*xorm.Engine {
	var adapters []*xorm.Engine = make([]*xorm.Engine, len(d))

	for key, val := range d {
		adapterPoT.sConns = &srcConns {
			Repo:     "",
			Host:     "",
			Port:     3306,
			Username: "",
			Password: "",
		}

		var ok bool
		var v map[interface{}]interface{} = val.(map[interface{}]interface{})
		
		repo, ok := v["repo"]
		if ok == false {
			continue
		}

		host, ok := v["host"]
		if ok == false {
			continue
		}

		port, ok := v["port"]
		if ok == false {
			continue
		}

		username, ok := v["username"]
		if ok == false {
			continue
		}

		password, ok := v["password"]
		if ok == false {
			continue
		}

		adapterPoT.sConns.Repo = cast.ToString(repo)
		adapterPoT.sConns.Host = cast.ToString(host)
		adapterPoT.sConns.Port = cast.ToInt(port)
		adapterPoT.sConns.Username = cast.ToString(username)
		adapterPoT.sConns.Password = cast.ToString(password)

		adapters[key] = adapterPoT.instance()
	}

	return adapters
}

func (adapterPoT *PoT) instance() *xorm.Engine {
	objODBC := newODBC()
	objODBC.dn 				= adapterPoT.driver
	objODBC.sParam			= adapterPoT.sParam
	objODBC.sConns			= adapterPoT.sConns
	objODBC.timeLocation	= adapterPoT.timeLocation

	return objODBC.instance().xEngine()
}