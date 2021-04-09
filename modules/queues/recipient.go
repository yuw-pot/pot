// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package queues

import (
	"github.com/spf13/cast"
)

type (
	Recipient interface {
		Receipted(channel, content interface{})
	}

	recipient struct {

	}
)

func newRecipient() *recipient {
	return &recipient {

	}
}

func (srv *recipient) do(msg string, key interface{}, channels *Channels) {
	p, ok := (*channels)[cast.ToString(key)]
	if ok {
		for _, val := range p {
			val.Receipted(key, msg)
		}
	}
}
