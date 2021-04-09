// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package queues

import (
	"context"
	"github.com/spf13/cast"
	cache "github.com/yuw-pot/pot/modules/cache/redis"
	"time"
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

func (client *rds) que() {
	if len(*client.info.KeYs) == 0 || client.info.Pool.Config == nil {
		return
	}

	d := cast.ToString(client.info.Pool.Config.(RdSParameters).Info)
	r, err := cache.Adapter(d)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	recipient := newRecipient()

	go func() {
		for {
			for _, v := range *client.info.KeYs {
				if num, _ := r.Exists(ctx, cast.ToString(v)).Result(); num == 0 {
					continue
				}

				receiver, err := r.RPop(ctx, cast.ToString(v)).Result()
				if err != nil {
					continue
				}

				recipient.do(receiver, v, client.info.Channels)
			}
			
			time.Sleep((*client).info.Interval.Value * (*client).info.Interval.Units)
		}
	}()
}

func (client *kfk) setInfo(info *PoT) {
	client.info = info
}

func (client *kfk) que() {

}











