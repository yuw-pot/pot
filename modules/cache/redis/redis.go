// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package redis

import (
	"github.com/go-redis/redis/v8"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"strings"
	"sync"
	"time"
)

const (
	network string = "tcp"
	addr string = "127.0.0.1:6379"
)

var (
	_ adapter = &redisPoT{}
	_ adapter = &redisClusterPoT{}

	adapterRedis map[string]*redis.Client = map[string]*redis.Client{}
	adapterRedisCluster map[string]*redis.ClusterClient = map[string]*redis.ClusterClient{}
)

type (
	InF struct {
		Conns interface{}
	}

	adapter interface {
		initialized(inf *InF)
	}

	redisClusterPoT struct {
		mx *sync.Mutex
		addrs []string
		username string
		password string
	}

	redisPoT struct {
		mx *sync.Mutex
		sParams *srcParams
	}

	srcParams struct {
		network		string
		addr 		string
		username	string
		password 	string
		db 			int
		poolsize	int
		pooltimeout	time.Duration
	}
)

func Initialized(inf *InF) {
	if len(adapterRedis) == 0 {
		var singletonClient adapter = newSingleton()
		singletonClient.initialized(inf)
	}

	if len(adapterRedisCluster) == 0 {
		var clusterClient adapter = newCluster()
		clusterClient.initialized(inf)
	}
}

func Adapter(d string) (*redis.Client, error) {
	c, ok := adapterRedis[strings.ToLower(d)]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "RedEngineErr")
	}

	return c, nil
}

func AdapterCluster(d string) (*redis.ClusterClient, error) {
	c, ok := adapterRedisCluster[strings.ToLower(d)]
	if ok == false {
		return nil, E.Err(data.ErrPfx, "RedEngineErr")
	}

	return c, nil
}