// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package subscriber

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/src/components/des"
	"strings"
)

type Component struct {

}

func New() *Component {
	return &Component {}
}

func (sub *Component) KeY(pfx, key interface{}) (interface{}, error) {
	if pfx == nil || key == nil {
		return nil, E.Err(data.ErrPfx, "ComponentSubscribeKeY")
	}

	inf := strings.Join([]string{cast.ToString(pfx), cast.ToString(key)}, "_")
	cry := des.New()
	res, err := cry.Encrypt(data.MD5, inf)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sub *Component) KeYEncryptByAES(pfx, key, salt interface{}) (interface{}, error) {
	if pfx == nil || key == nil || salt == nil {
		return nil, E.Err(data.ErrPfx, "ComponentSubscribeKeY")
	}

	inf := strings.Join([]string{cast.ToString(pfx), cast.ToString(key)}, "_")
	cry := des.New()
	x, err := cry.AES(salt)
	if err != nil {
		return nil, err
	}

	res, err := x.EncrypT(inf)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (sub *Component) KeYDecryptByAES(cipher, salt interface{}) (interface{}, interface{}, error) {
	if cipher == nil || salt == nil {
		return nil, nil, E.Err(data.ErrPfx, "ComponentSubscribeKeY")
	}

	cry := des.New()
	x, err := cry.AES(salt)
	if err != nil {
		return nil, nil, err
	}

	d, err := x.DecrypT(cipher)
	if err != nil {
		return nil, nil, err
	}

	res := strings.Split(cast.ToString(d), "_")
	if len(res) != 2 {
		return nil, nil, E.Err(data.ErrPfx, "ComponentSubscribeKeY")
	}

	return res[0], res[1], nil
}
