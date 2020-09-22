// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"math/rand"
	"strings"
	"time"
)

type (
	PoT struct {

	}
)

func New() *PoT {
	return &PoT {

	}
}

func (u *PoT) SetTimeLocation(d string) (*time.Location, error) {
	if d == "" {
		d = data.TimeLocation
	}

	return time.LoadLocation(d)
}

func (u *PoT) Contains(k interface{}, d ...interface{}) bool {
	if len(d) < 1 {
		return false
	}

	for _, v := range d {
		if strings.Contains(cast.ToString(k), cast.ToString(v)) {
			return true
		}
	}

	return false
}

func (u *PoT) NumRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
