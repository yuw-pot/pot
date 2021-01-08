// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package rbac

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/casbin"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
)

type Component struct {
	u *utils.PoT
	enforcer *casbin.PoT
}

func New() *Component {
	return &Component {
		u: utils.New(),
		enforcer: casbin.New(),
	}
}

func (rbac *Component) Initialized(mod, adapter string, info interface{}) error {
	var ok bool
	ok = rbac.u.Contains(mod, casbin.RBAC, casbin.RBACMultiple, casbin.RBACMultipleRole)
	if ok == false {
		return E.Err(data.ErrPfx, "ComponentCasbinParamMod")
	}

	ok = rbac.u.Contains(adapter, casbin.RBACAdapterRedis)
	if ok == false {
		return E.Err(data.ErrPfx, "ComponentCasbinParamAdapter")
	}

	rbac.enforcer.Mod = mod
	rbac.enforcer.Adapter = adapter
	rbac.enforcer.AdapterInfo = info

	return rbac.enforcer.Enforcer()
}

func (rbac *Component) AddPolicy(d ... interface{}) error {
	return rbac.enforcer.Add(d ...)
}

func (rbac *Component) GeT() *casbin.PoT {
	return rbac.enforcer
}


