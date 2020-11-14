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
	"github.com/yuw-pot/pot/modules/properties"
	"strings"
	"sync"
)

type (
	redisPoT struct {
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

var adapterRedis map[string]*redis.Client

func Made(d string) (*redis.Client, error) {
	d = strings.ToLower(d)

	_, ok := adapterRedis[d]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "")
	}

	return adapterRedis[d], nil
}

func New() *redisPoT {
	client := &redisPoT{}
	client.mx = &sync.Mutex{}
	
	return client.initialized()
}

func (rd *redisPoT) initialized() *redisPoT {
	redisPoT := properties.PropertyPoT.GeT("Redis", nil)
	if redisPoT == nil {
		panic(E.Err(data.ErrPfx, "RedParamsErr"))
	}

	adapterRedis = map[string]*redis.Client{}
	for key, val := range redisPoT.(map[string]interface{}) {
		rd.sParams = &srcParams {
			network:  "",
			addr:     "",
			password: "",
			db:       0,
		}

		var ok bool

		network, ok := val.(map[string]interface{})["network"]
		if ok {
			rd.sParams.network = cast.ToString(network)
		}

		addr, ok := val.(map[string]interface{})["addr"]
		if ok {
			rd.sParams.addr = cast.ToString(addr)
		}

		password, ok := val.(map[string]interface{})["password"]
		if ok {
			rd.sParams.password = cast.ToString(password)
		}

		db, ok := val.(map[string]interface{})["db"]
		if ok {
			rd.sParams.db = cast.ToInt(db)
		}

		adapterRedis[strings.ToLower(key)] = rd.instance()
	}

	return rd
}

func (rd *redisPoT) instance() *redis.Client {
	rd.mx.Lock()
	defer func() { rd.mx.Unlock() }()

	client := redis.NewClient(&redis.Options{
		Network: 	rd.sParams.network,
		Addr: 		rd.sParams.addr,
		Password: 	rd.sParams.password,
		DB: 		rd.sParams.db,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil { panic(err) }

	return client
}