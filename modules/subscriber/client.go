// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package subscriber

import (
	"context"
	"github.com/spf13/cast"
	cache "github.com/yuw-pot/pot/modules/cache/redis"
)

type (
	rds struct {
		pool *Pool
	}

	kfk struct {
		pool *Pool
	}
)

func (client *rds) setPool(pool *Pool) {
	client.pool = pool
}

func (client *rds) sub(channels ... string) {
	r, err := cache.Made(cast.ToString((*client.pool)[1]))
	if err != nil { panic(err) }

	ctx := context.Background()
	receiver := r.Subscribe(ctx, channels ...)
	_, err = receiver.Receive(ctx)
	if err != nil {
		return
	}

	provider := newProvider()
	for msg := range receiver.Channel() {
		provider.do(msg)
	}
}

func (client *kfk) setPool(pool *Pool) {
	client.pool = pool
}

func (client *kfk) sub(channels ... string) {

}
