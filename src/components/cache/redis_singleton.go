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
	"strings"
	"time"
)

type RedisComponent struct {
	pfx string
	ctx context.Context
	client *redis.Client
}

func NewRedis(d string) *RedisComponent {
	r := &RedisComponent {
		pfx: "",
		ctx: context.Background(),
	}

	r.client, _ = cache.Adapter(d)
	return r
}

func (r *RedisComponent) GeTClient() *redis.Client {
	return r.client
}

func (r *RedisComponent) KeYs(d ... interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

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
	if err := r.check(); err != nil {
		return false, err
	}

	err := r.client.Get(r.ctx, r.addPrefix(cast.ToString(d))).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedisComponent) HSeT(key interface{}, d map[string]interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	if key == nil || d == nil {
		return nil, E.Err(data.ErrPfx, "ComponentCacheSeTParams")
	}

	return r.client.HSet(r.ctx, r.addPrefix(cast.ToString(key)), d).Result()
}

func (r *RedisComponent) HGeT(key, field string) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.HGet(r.ctx, r.addPrefix(key), field).Result()
}

func (r *RedisComponent) HDeL(key string, val ... string) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.HDel(r.ctx, key, val ...).Result()
}

func (r *RedisComponent) IsHExisT(key, field string) (bool, error) {
	if err := r.check(); err != nil {
		return false, err
	}

	err := r.client.HGet(r.ctx, r.addPrefix(key), field).Err()
	if err != nil {
		return false, err
	}

	return true, nil
}

func (r *RedisComponent) LPush(key string, val ... interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.LPush(r.ctx, key, val ...).Result()
}

func (r *RedisComponent) RPop(key string) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.RPop(r.ctx, key).Result()
}

func (r *RedisComponent) Publish(channel string, message interface{}) (interface{}, error) {
	if err := r.check(); err != nil {
		return nil, err
	}

	return r.client.Publish(r.ctx, channel, message).Result()
}

func (r *RedisComponent) SeTPrefix(pfx string) error {
	if pfx == "" { return E.Err(data.ErrPfx, data.PoTErrorNull) }

	cryptoMd := crypto.New()
	cryptoMd.Mode = data.ModeToken
	cryptoMd.D = []interface{}{data.MD5, pfx}

	cryptoEn, err := cryptoMd.Made()
	if err != nil { return err }

	r.pfx = cast.ToString(cryptoEn)
	return nil
}

func (r *RedisComponent) KeYFormaT(key string) string {
	return r.addPrefix(key)
}

func (r *RedisComponent) Cache() interface{} {
	return r
}

func (r *RedisComponent) GeTPrefix() string {
	return r.pfx
}

func (r *RedisComponent) GeTx() context.Context {
	return r.ctx
}

func (r *RedisComponent) addPrefix(key string) string {
	if r.pfx == "" { return key }
	return strings.Join([]string{r.pfx, key}, "_")
}

func (r *RedisComponent) check() error {
	if r.client == nil { return E.Err(data.ErrPfx, "ComponentCacheClient") }
	return nil
}


