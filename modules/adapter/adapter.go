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
		v 				*utils.PoT
		sConns 			*srcConns
		sParam 			*srcParam
		timeLocation 	*time.Location

		driver 			interface{}
	}

	srcParam struct {
		Policy 		string
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

func New() *PoT {
	adapter := &PoT{}
	adapter.v = utils.New()

	return adapter.initialized()
}

func (adapterPoT *PoT) initialized() *PoT {
	// Set db Time Location
	//   - GeT Configure
	adapterTimeLocation := properties.PropertyPoT.GeT("PoT.TimeLocation", data.TimeLocation)
	d, err := adapterPoT.v.SetTimeLocation(cast.ToString(adapterTimeLocation))
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

			adapterMaster[dn][label] = adapterPoT.validate(master.(map[string]interface{}))
			if adapterMaster[dn][label] == nil {
				content := fmt.Sprintf("Adapter.%v.Conns.%v.Master", dn, label)
				panic(E.Err(data.ErrPfx, "AdapterConfigErr", content))
			}

			adapterSlaver[dn][label] = make([]*xorm.Engine, len(slaver.([]interface{})))
			adapterSlaver[dn][label] = adapterPoT.validateGroup(slaver.([]interface{}) ...)
		}
	}

	return adapterPoT
}

func (adapterPoT *PoT) validate(d map[string]interface{}) *xorm.Engine {
	adapterPoT.sConns = &srcConns {
		Repo:     "",
		Host:     "",
		Port:     0,
		Username: "",
		Password: "",
	}

	repo, ok := d["repo"]
	if ok == false { return nil }

	host, ok := d["host"]
	if ok == false { return nil }

	port, ok := d["port"]
	if ok == false { return nil }

	username, ok := d["username"]
	if ok == false { return nil }

	password, ok := d["password"]
	if ok == false { return nil }

	adapterPoT.sConns.Repo = cast.ToString(repo)
	adapterPoT.sConns.Host = cast.ToString(host)
	adapterPoT.sConns.Port = cast.ToInt(port)
	adapterPoT.sConns.Username = cast.ToString(username)
	adapterPoT.sConns.Password = cast.ToString(password)

	return adapterPoT.instance()
}

func (adapterPoT *PoT) validateGroup(d ... interface{}) []*xorm.Engine {
	var adapters []*xorm.Engine = make([]*xorm.Engine, len(d))

	for key, val := range d {
		adapterPoT.sConns = &srcConns {
			Repo:     "",
			Host:     "",
			Port:     0,
			Username: "",
			Password: "",
		}
		
		repo, ok := val.(map[interface{}]interface{})["Repo"]
		if ok == false {
			continue
		}

		host, ok := val.(map[interface{}]interface{})["Host"]
		if ok == false {
			continue
		}

		port, ok := val.(map[interface{}]interface{})["Port"]
		if ok == false {
			continue
		}

		username, ok := val.(map[interface{}]interface{})["Username"]
		if ok == false {
			continue
		}

		password, ok := val.(map[interface{}]interface{})["Password"]
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