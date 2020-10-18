// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	R "github.com/yuw-pot/pot/modules/cache/redis"
	E "github.com/yuw-pot/pot/modules/err"
	U "github.com/yuw-pot/pot/modules/utils"
	"time"
)

type (
	RedisComponent struct {
		vPoT *U.PoT
		client *redis.Client
		ctx context.Context
	}
)

func NewRedis(d string) *RedisComponent {
	r := &RedisComponent {
		vPoT: U.New(),
		ctx: context.Background(),
	}

	r.client, _ = R.Made(d)
	return r
}

func (r *RedisComponent) SeT(k interface{}, v interface{}, d ... time.Duration) error {
	if err := r.check(); err != nil {
		return err
	}

	if k == nil || v == nil {
		return E.Err(data.ErrPfx, "CpCacheSeTParams")
	}

	var expiration time.Duration = 0
	if len(d) == 1 {
		expiration = d[0]
	}

	res := r.client.Set(r.ctx, cast.ToString(k), v, expiration)
	if res.Err() != nil {
		return res.Err()
	}

	return nil
}

func (r *RedisComponent) GeT(d interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	res, err := r.client.Get(r.ctx, cast.ToString(d)).Result()
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (r *RedisComponent) IsExisT(d ... string) {
	r.client.Exists(r.ctx, d ...)
}

func (r *RedisComponent) check() error {
	if r.client == nil {
		return E.Err(data.ErrPfx, "CpCacheClient")
	}

	return nil
}


