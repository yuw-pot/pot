// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package request

import (
	"bytes"
	"context"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"io/ioutil"
	"net/http"
)

type (
	Component struct {
		u *utils.PoT
	}

	InF struct {
		Api 	interface{}
		Method 	interface{}
	}
)

func New() *Component {
	return &Component {
		u: utils.New(),
	}
}

func (req *Component) Send(d []byte, inf *InF) ([]byte, error) {
	if inf.Method == "" {
		inf.Method = data.ReqMethodPOST
	}

	if ok := req.u.Contains(inf.Method, data.ReqMethods ...); ok == false {
		return nil, E.Err(data.ErrPfx, "ComponentRequestMethod")
	}

	if inf.Api == nil {
		return nil, E.Err(data.ErrPfx, "ComponentRequestURL")
	}

	buf := bytes.NewBuffer(d)
	res, err := http.NewRequest(cast.ToString(inf.Method), cast.ToString(inf.Api), buf)
	if err != nil {
		return nil, err
	}

	res.Header.Set("Content-Type", "application/json;charset=UTF-8")

	obj := http.Client{}
	client, err := obj.Do(res.WithContext(context.TODO()))
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadAll(client.Body)
	if err != nil {
		return nil, err
	}

	return result, nil
}
