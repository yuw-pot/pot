// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package subscriber

import (
	"github.com/spf13/cast"
)

func StarT(subscriberPoT *PoT) {
	if subscriberPoT.KeYs == nil || subscriberPoT.Pool == nil || subscriberPoT.Channels == nil {
		return
	}

	keys = subscriberPoT.KeYs
	pool = subscriberPoT.Pool
	channels = subscriberPoT.Channels

	go new(subscribed).do()
}

const (
	MethodRdS string = "rds"
	MethodKfK string = "kfk"
)

type (
	subscriber interface {
		setPool(pool *Pool)
		sub(channels ... string)
	}

	KeYs []interface{}
	Pool []interface{}
	Channels map[string][]Provider

	PoT struct {
		KeYs *KeYs
		Channels *Channels
		Pool *Pool
	}

	subscribed struct {}
)

var (
	_ subscriber = new(rds)
	_ subscriber = new(kfk)

	keys *KeYs
	pool *Pool
	channels *Channels
)

func (sub *subscribed) do() {
	if len(*pool) <= 1 {
		return
	}

	var client subscriber

	switch (*pool)[0] {
	case MethodRdS:

		client = new(rds)
		break

	case MethodKfK:

		client = new(kfk)
		break

	default:
		return
	}

	keYsToStrSlice := make([]string, len(*keys))
	for i, val := range *keys {
		keYsToStrSlice[i] = cast.ToString(val)
	}

	client.setPool(pool)
	client.sub(keYsToStrSlice ...)
}
