// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	cache "github.com/yuw-pot/pot/modules/cache/redis"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"time"
)

type (
	RedisComponent struct {
		v *utils.PoT
		client *redis.Client
		ctx context.Context
	}
)

func NewRedis(d string) *RedisComponent {
	r := &RedisComponent {
		v: utils.New(),
		ctx: context.Background(),
	}

	r.client, _ = cache.Made(d)
	return r
}

func (r *RedisComponent) Cache() *RedisComponent {
	return r
}

func (r *RedisComponent) SeT(key, val interface{}, d ... time.Duration) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	if key == nil || val == nil {
		return nil, E.Err(data.ErrPfx, "ComponentCacheSeTParams")
	}

	var expiration time.Duration = 0
	if len(d) == 1 {
		expiration = d[0]
	}

	return r.client.Set(r.ctx, cast.ToString(key), val, expiration).Result()
}

func (r *RedisComponent) GeT(d interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.Get(r.ctx, cast.ToString(d)).Result()
}

func (r *RedisComponent) DeL(d interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.Del(r.ctx, cast.ToString(d)).Result()
}

func (r *RedisComponent) IsExisT(d ... string) {
	r.client.Exists(r.ctx, d ...)
}

func (r *RedisComponent) check() error {
	if r.client == nil {
		return E.Err(data.ErrPfx, "ComponentCacheClient")
	}

	return nil
}


