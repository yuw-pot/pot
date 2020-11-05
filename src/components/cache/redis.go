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
	"github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"strings"
	"time"
)

type (
	RedisComponent struct {
		pfx string
		ctx context.Context
		v *utils.PoT
		client *redis.Client
	}
)

func NewRedis(d string) *RedisComponent {
	r := &RedisComponent {
		pfx: "",
		ctx: context.Background(),
		v: utils.New(),
	}

	r.client, _ = cache.Made(d)
	return r
}

func (r *RedisComponent) KeYs(d ... interface{}) (interface{}, error) {
	var clientPattern string = "*"
	if len(d) == 1 {
		clientPattern = cast.ToString(d[0])
	}

	return r.client.Keys(r.ctx, clientPattern).Result()
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

	return r.client.Set(r.ctx, r.addPrefix(cast.ToString(key)), val, expiration).Result()
}

func (r *RedisComponent) GeT(d interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.Get(r.ctx, r.addPrefix(cast.ToString(d))).Result()
}

func (r *RedisComponent) DeL(d interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.Del(r.ctx, r.addPrefix(cast.ToString(d))).Result()
}

func (r *RedisComponent) IsExisT(d interface{}) (bool, error) {
	err := r.client.Get(r.ctx, r.addPrefix(cast.ToString(d))).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedisComponent) HSeT(key interface{}, d map[string]interface{}) (interface{}, error)  {
	if err := r.check(); err != nil {
		return nil, err
	}

	if key == nil || d == nil {
		return nil, E.Err(data.ErrPfx, "ComponentCacheSeTParams")
	}

	return r.client.HSet(r.ctx, r.addPrefix(cast.ToString(key)), d).Result()
}

func (r *RedisComponent) HGeT(key, field string) (interface{}, error) {
	return r.client.HGet(r.ctx, r.addPrefix(key), field).Result()
}

func (r *RedisComponent) IsHExisT(key, field string) (bool, error) {
	err := r.client.HGet(r.ctx, r.addPrefix(key), field).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedisComponent) SeTPrefix(pfx string) *RedisComponent {
	if pfx == "" {
		return r
	}

	cryptoMd := crypto.New()
	cryptoMd.Mode = data.ModeToken
	cryptoMd.D = []interface{}{data.MD5, pfx}

	cryptoEn, err := cryptoMd.Made()
	if err != nil {
		return r
	}

	r.pfx = cast.ToString(cryptoEn)
	return r
}

func (r *RedisComponent) KeYFormaT(key string) string {
	return r.addPrefix(key)
}

func (r *RedisComponent) Cache() *RedisComponent {
	return r
}

func (r *RedisComponent) GeTPrefix() string {
	return r.pfx
}

func (r *RedisComponent) GeTClienT() *redis.Client {
	return r.client
}

func (r *RedisComponent) GeTx() context.Context {
	return r.ctx
}

func (r *RedisComponent) addPrefix(key string) string {
	if r.pfx == "" {
		return key
	}

	return strings.Join([]string{r.pfx, key}, "_")
}

func (r *RedisComponent) check() error {
	if r.client == nil {
		return E.Err(data.ErrPfx, "ComponentCacheClient")
	}

	return nil
}


