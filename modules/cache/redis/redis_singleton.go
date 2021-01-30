// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/modules/properties"
	"strings"
	"sync"
)

func newSingleton() *redisPoT {
	return &redisPoT{ mx: &sync.Mutex{} }
}

func (rd *redisPoT) initialized(inf *InF) {
	var rds interface{}

	if inf.Conns != nil && len(inf.Conns.(map[string]interface{})) != 0 {
		rds = inf.Conns
	} else {
		rds = properties.PropertyPoT.GeT("Redis", nil)

	}

	adapterRedis = map[string]*redis.Client{}
	for key, val := range rds.(map[string]interface{}) {
		rd.sParams = &srcParams {
			network:  network,
			addr:     addr,
			username: "",
			password: "",
			db:       0,
		}

		var ok bool
		var v map[interface{}]interface{} = val.(map[interface{}]interface{})

		network, ok := v["network"]
		if ok {
			rd.sParams.network = cast.ToString(network)
		}

		addr, ok := v["addr"]
		if ok {
			rd.sParams.addr = cast.ToString(addr)
		}

		username, ok := v["username"]
		if ok {
			rd.sParams.username = cast.ToString(username)
		}

		password, ok := v["password"]
		if ok {
			rd.sParams.password = cast.ToString(password)
		}

		db, ok := v["db"]
		if ok {
			rd.sParams.db = cast.ToInt(db)
		}

		poolSize, ok := v["poolsize"]
		if ok {
			rd.sParams.poolsize = cast.ToInt(poolSize)
		} else {
			rd.sParams.poolsize = 10
		}

		adapterRedis[strings.ToLower(key)] = rd.instance()
	}
}

func (rd *redisPoT) instance() *redis.Client {
	rd.mx.Lock()
	defer func() { rd.mx.Unlock() }()

	client := redis.NewClient(&redis.Options{
		Network: 	rd.sParams.network,
		Addr: 		rd.sParams.addr,
		Username: 	rd.sParams.username,
		Password: 	rd.sParams.password,
		DB: 		rd.sParams.db,
		PoolSize:	rd.sParams.poolsize,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil { panic(err) }

	return client
}