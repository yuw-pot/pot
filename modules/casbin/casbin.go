// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package casbin

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/redis-adapter/v2"
	"github.com/casbin/xorm-adapter/v2"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/properties"
	"strings"
	"xorm.io/xorm"
)

const (
	RBAC 				string = "rbac"
	RBACMultiple 		string = "rbac_multiple"
	RBACMultipleRole	string = "rbac_multiple_role"

	RBACAdapterXorm		string = "xorm"
	RBACAdapterRedis	string = "redis"
)

type (
	PoT struct {
		Mod string
		Adapter string
		AdapterInfo interface{}

		enforcer *casbin.Enforcer
	}

	AdapterXorm struct {
		Engine *xorm.Engine
		Table string
	}

	AdapterRedis struct {
		Tag string
	}
)

func New() *PoT {
	return &PoT {
		Mod: RBAC,
		Adapter: RBACAdapterRedis,
		AdapterInfo: nil,
	}
}

func (rbac *PoT) Enforcer() error {
	casbinModel, err := rbac.casbinModelTpL()
	if err != nil {
		return err
	}

	casbinAdapter := rbac.casbinAdapterTpL()
	if casbinAdapter == nil {
		return E.Err(data.ErrPfx, data.PoTErrorNull)
	}

	rbac.enforcer, err = casbin.NewEnforcer(casbinModel, casbinAdapter)
	if err != nil {
		return err
	}

	if err := rbac.enforcer.LoadPolicy(); err != nil {
		return err
	}

	return nil
}

func (rbac *PoT) Check(d ... interface{}) (bool, error) {
	if err := rbac.checkEnforcer(); err != nil {
		return false, err
	}

	return rbac.enforcer.Enforce(d ...)
}

func (rbac *PoT) Add(d ... interface{}) error {
	if err := rbac.checkEnforcer(); err != nil { return err }

	ok, err := rbac.enforcer.AddPolicy(d ...)
	if err != nil { return err }
	if ok == false { return E.Err(data.ErrPfx, data.PoTErrorNull) }

	err = rbac.enforcer.SavePolicy()
	if err != nil { return err }

	return nil
}

func (rbac *PoT) casbinAdapterTpL() interface{} {
	if rbac.AdapterInfo == nil {
		return nil
	}

	var rbacAdapter interface{}

	switch rbac.Adapter {

	case RBACAdapterXorm:

		var (
			err error
			xormAdapter *AdapterXorm = rbac.AdapterInfo.(*AdapterXorm)
		)

		if xormAdapter.Engine == nil {
			return nil
		}

		if xormAdapter.Table == "" {
			rbacAdapter, err = xormadapter.NewAdapterByEngine(xormAdapter.Engine)
			if err != nil { return nil }
		} else {
			rbacAdapter, err = xormadapter.NewAdapterByEngineWithTableName(xormAdapter.Engine, xormAdapter.Table, )
			if err != nil { return nil }
		}

		return rbacAdapter

	case RBACAdapterRedis:

		var redisAdapter *AdapterRedis = rbac.AdapterInfo.(*AdapterRedis)

		if redisAdapter.Tag == "" {
			return nil
		}

		adapterKeY := strings.Join([]string{"Redis", redisAdapter.Tag}, ".")
		adapterVaL := properties.PropertyPoT.GeT(adapterKeY, map[string]interface{}{}).(map[string]interface{})

		var ok bool
		addr, ok := adapterVaL["addr"]
		if ok == false {
			return nil
		}

		network, ok := adapterVaL["network"]
		if ok == false {
			return nil
		}

		pwd, ok := adapterVaL["password"]
		if ok == false {
			return nil
		}

		rbacAdapter = redisadapter.NewAdapterWithPassword(cast.ToString(network), cast.ToString(addr), cast.ToString(pwd))
		return rbacAdapter

	default:
		return nil
	}
}

func (rbac *PoT) casbinModelTpL() (model.Model, error) {
	var rbacTxT string

	switch rbac.Mod {
	case RBAC:
		rbacTxT = `[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && r.obj == p.obj && r.act == p.act
`
	case RBACMultiple:
		rbacTxT = `
[request_definition]
r = sub, dom, obj, act

[policy_definition]
p = sub, dom, obj, act

[role_definition]
g = _, _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub, r.dom) && r.dom == p.dom && r.obj == p.obj && r.act == p.obj
`

	case RBACMultipleRole:
		rbacTxT = `
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _
g2 = _, _

[policy_effect]
e = some(where (p.eft == allow))

[matchers]
m = g(r.sub, p.sub) && g2(r.obj, p.obj) && r.act == p.act
`

	default:
		return nil, E.Err(data.ErrPfx, data.PoTErrorNull)
	}

	return model.NewModelFromString(rbacTxT)
}

func (rbac *PoT) checkEnforcer() error {
	if rbac.enforcer == nil {
		return E.Err(data.ErrPfx, "CasbinEnforcer")
	}

	return nil
}

