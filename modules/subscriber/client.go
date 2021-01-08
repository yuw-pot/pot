// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package subscriber

import (
	"context"
	"github.com/spf13/cast"
	cache "github.com/yuw-pot/pot/modules/cache/redis"
)

type (
	rds struct {
		info *PoT
	}

	RdSParameters struct {
		Info interface{}
	}

	kfk struct {
		info *PoT
	}

	KfKParameters struct {

	}
)

func (client *rds) setInfo(info *PoT) {
	client.info = info
}

func (client *rds) sub() {
	if len(*client.info.KeYs) == 0 || client.info.Pool.Config == nil {
		return
	}

	keYsToStrSlice := make([]string, len(*client.info.KeYs))
	for i, v := range *client.info.KeYs {
		keYsToStrSlice[i] = cast.ToString(v)
	}

	d := cast.ToString(client.info.Pool.Config.(RdSParameters).Info)
	r, err := cache.Adapter(d)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	receiver := r.Subscribe(ctx, keYsToStrSlice ...)

	_, err = receiver.Receive(ctx)
	if err != nil {
		return
	}

	provide := newProvider()
	for msg := range receiver.Channel() {
		provide.do(msg, client.info.KeYs, client.info.Channels)
	}
}

func (client *kfk) setInfo(info *PoT) {
	client.info = info
}

func (client *kfk) sub() {

}
