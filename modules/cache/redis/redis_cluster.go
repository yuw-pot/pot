// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/modules/properties"
	"sync"
)

func newCluster() *redisClusterPoT {
	return &redisClusterPoT{ mx: &sync.Mutex{} }
}

func (rd *redisClusterPoT) initialized(inf *InF) {
	cfg := properties.PropertyPoT.GeT("RedisCluster", nil)
	if cfg == nil {
		return
	}

	if inf != nil {}

	for key, val := range cfg.(map[string]interface{}) {
		rd.addrs = []string{}
		rd.username = ""
		rd.password = ""

		var ok bool
		var v map[string]interface{} = val.(map[string]interface{})

		addrs, ok := v["addrs"]
		if ok {
			rd.addrs = cast.ToStringSlice(addrs)
		} else {
			continue
		}

		username, ok := v["username"]
		if ok { rd.username = cast.ToString(username) }

		password, ok := v["password"]
		if ok { rd.password = cast.ToString(password) }

		adapterRedisCluster[key] = rd.instance()
	}
}

func (rd *redisClusterPoT) instance() *redis.ClusterClient {
	rd.mx.Lock()
	defer func() { rd.mx.Unlock() }()

	client := redis.NewClusterClient(&redis.ClusterOptions{
		Addrs:		rd.addrs,
		Username:	rd.username,
		Password:	rd.password,
	})

	_, err := client.Ping(context.Background()).Result()
	if err != nil { panic(err) }

	return client
}
