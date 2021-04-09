// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package subscriber

func Start(subscriberPoT *PoT) {
	if subscriberPoT.KeYs == nil || subscriberPoT.Pool == nil || subscriberPoT.Channels == nil {
		panic("Subscriber Parameters is Nil!")
	}

	sub := new(subscribed)
	sub.info = subscriberPoT

	go sub.do()
}

const (
	MethodRdS string = "rds"
	MethodKfK string = "kafka"
)

type (
	KeYs []interface{}
	Pool struct {
		Method 	string
		Config 	interface{}
	}
	Channels map[interface{}][]Provider

	subscriber interface {
		setInfo(info *PoT)
		sub()
	}

	PoT struct {
		KeYs *KeYs
		Pool *Pool
		Channels *Channels
	}

	subscribed struct {
		info *PoT
	}
)

var (
	_ subscriber = new(rds)
	_ subscriber = new(kfk)
)

func (sub *subscribed) do() {
	var client subscriber

	switch sub.info.Pool.Method {
	case MethodRdS:

		client = new(rds)
		break

	case MethodKfK:

		client = new(kfk)
		break

	default:
		return
	}

	client.setInfo(sub.info)
	client.sub()
}
