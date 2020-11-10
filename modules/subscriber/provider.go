// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package subscriber

import (
	"github.com/go-redis/redis/v8"
	"github.com/yuw-pot/pot/modules/utils"
)

type (
	Provider interface {
		Provided(channel string, content interface{})
	}

	provider struct {
		v *utils.PoT
	}
)

func newProvider() *provider {
	return &provider{
		v: utils.New(),
	}
}

func (srv *provider) do(msg *redis.Message) {
	if ok := srv.v.Contains(msg.Channel, *keys ...); ok == false {
		return
	}

	p, ok := (*channels)[msg.Channel]
	if ok {
		for _, val := range p {
			val.Provided(msg.Channel, msg.Payload)
		}
	}
}
