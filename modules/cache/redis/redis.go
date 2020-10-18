// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	P "github.com/yuw-pot/pot/modules/properties"
	"strings"
	"sync"
)

type (
	rPoT struct {
		rd *redis.Client
		mx *sync.Mutex
		sParams *srcParams
	}

	srcParams struct {
		network		string
		addr 		string
		password 	string
		db 			int
	}
)

var (
	adapterRedis map[string]*redis.Client
)

func Made(d string) (*redis.Client, error) {
	d = strings.ToLower(d)

	_, ok := adapterRedis[d]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "")
	}

	return adapterRedis[d], nil
}

func New() *rPoT {
	return &rPoT {
		mx: &sync.Mutex{},
	}
}

func (r *rPoT) Made() {
	redisPoT := P.PropertyPoT.GeT("Redis", nil)
	if redisPoT == nil {
		panic(E.Err(data.ErrPfx, "RedParamsErr"))
	}

	adapterRedis = map[string]*redis.Client{}
	for k, v := range redisPoT.(map[string]interface{}) {
		r.sParams = &srcParams {
			network:  "",
			addr:     "",
			password: "",
			db:       0,
		}

		if _, ok := v.(map[string]interface{})["Network"]; ok {
			r.sParams.network = cast.ToString(v.(map[string]interface{})["Network"])
		}

		if _, ok := v.(map[string]interface{})["Addr"]; ok {
			r.sParams.addr = cast.ToString(v.(map[string]interface{})["Addr"])
		}

		if _, ok := v.(map[string]interface{})["Password"]; ok {
			r.sParams.password = cast.ToString(v.(map[string]interface{})["Password"])
		}

		if _, ok := v.(map[string]interface{})["DB"]; ok {
			r.sParams.db = cast.ToInt(v.(map[string]interface{})["DB"])
		}

		adapterRedis[strings.ToLower(k)] = r.instance()
	}
}

func (r *rPoT) instance() *redis.Client {
	r.mx.Lock()

	defer func() {
		r.mx.Unlock()
	}()

	client := redis.NewClient(&redis.Options{
		Network: 	r.sParams.network,
		Addr: 		r.sParams.addr,
		Password: 	r.sParams.password,
		DB: 		r.sParams.db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}

	return client
}