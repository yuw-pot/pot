// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package cache

import (
	"context"
	"time"
)

var (
	_ Component = &RedisComponent{}
	_ Component = &RedisClusterComponent{}
)

type Component interface {
	KeYs(d ... interface{}) (interface{}, error)
	SeT(key, val interface{}, d ... time.Duration) (interface{}, error)
	GeT(d interface{}) (interface{}, error)
	DeL(d interface{}) (interface{}, error)
	IsExisT(d interface{}) (bool, error)
	HSeT(key interface{}, d map[string]interface{}) (interface{}, error)
	HGeT(key, field string) (interface{}, error)
	IsHExisT(key, field string) (bool, error)
	Publish(channel string, message interface{}) (interface{}, error)
	SeTPrefix(pfx string) error
	GeTPrefix() string
	KeYFormaT(key string) string
	GeTx() context.Context
	Cache() interface{}
}
