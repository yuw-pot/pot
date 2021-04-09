// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package queues

import "time"

func Start(queuePoT *PoT) {
	if queuePoT.KeYs == nil || queuePoT.Pool == nil || queuePoT.Channels == nil {
		panic("Queue Parameters is Nil!")
	}

	if queuePoT.Interval == nil {
		queuePoT.Interval = &Interval {
			Value: 3,
			Units: time.Second,
		}
	}

	que := new(queued)
	que.info = queuePoT

	que.do()
}

const (
	MethodRdS string = "rds"
	MethodKfK string = "kafka"
)

var (
	_ queue = new(rds)
	_ queue = new(kfk)
)

type (
	KeYs []interface{}
	Pool struct {
		Method 	string
		Config 	interface{}
	}
	Channels map[interface{}][]Recipient
	Interval struct {
		Value time.Duration
		Units time.Duration
	}


	PoT struct {
		KeYs *KeYs
		Pool *Pool
		Channels *Channels
		Interval *Interval
	}

	queue interface {
		setInfo(info *PoT)
		que()
	}

	queued struct {
		info *PoT
	}
)

func (que *queued) do() {
	var client queue

	switch que.info.Pool.Method {
	case MethodRdS:

		client = new(rds)
		break

	case MethodKfK:

		client = new(kfk)
		break

	default:
		return
	}

	client.setInfo(que.info)
	client.que()
}
