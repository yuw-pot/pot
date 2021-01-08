// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package des

import (
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/crypto"
)

type Component struct {}

func New() *Component {
	return &Component {}
}

func (cry *Component) Encrypt(method, key interface{}) (interface{}, error) {
	encrypt := crypto.New()

	encrypt.Mode = data.ModeToken
	encrypt.D = []interface{}{method, key}

	return encrypt.Made()
}

func (cry *Component) AES(key interface{}) (*crypto.AeSPoT, error) {
	encrypt := crypto.New()

	encrypt.Mode = data.ModeAeS
	encrypt.D = []interface{}{key}

	aes, err := encrypt.Made()

	return aes.(*crypto.AeSPoT), err
}

type RsAPeM struct {
	PubKeY string
	PriKey string
}

func (cry *Component) RSA(rsaPeM *RsAPeM, dir ... interface{}) (*crypto.RsAPoT, error) {
	encrypt := crypto.New()

	encrypt.Mode = data.ModeRsA
	encrypt.D = []interface{}{rsaPeM.PubKeY, rsaPeM.PriKey}

	// self define the directory of rsa public/private pem key
	if len(dir) == 1 {
		if dir[0] != nil && dir[0] != "" {
			encrypt.D = append(encrypt.D, dir[0])
		}
	}

	rsa, err := encrypt.Made()

	return rsa.(*crypto.RsAPoT), err
}