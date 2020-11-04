// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package routes

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/crypto"
	"github.com/yuw-pot/pot/modules/err"
)

const (
	PoTMethodAnY	 = "ANY"
	PoTMethodGeT  	 = "GET"
	PoTMethodHeaD    = "HEAD"
	PoTMethodPosT    = "POST"
	PoTMethodPuT     = "PUT"
	PoTMethodPatch   = "PATCH" // RFC 5789
	PoTMethodDelete  = "DELETE"
	PoTMethodConnect = "CONNECT"
	PoTMethodOptions = "OPTIONS"
	PoTMethodTrace   = "TRACE"
)

func GeTPath(service, controller, action string) (interface{}, error) {
	cry := crypto.New()

	cry.Mode = data.ModeToken
	cry.D = []interface{}{data.MD5, service + controller + action}

	k, errCrypto := cry.Made()
	if errCrypto != nil {
		return nil, errCrypto
	}

	_, ok := (*rMaP)[cast.ToString(k)]
	if ok == false {
		return nil, err.Err(data.ErrPfx, "ErrDefault")
	}

	return (*rMaP)[cast.ToString(k)], nil
}
