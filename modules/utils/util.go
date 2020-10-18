// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"io"
	"math/rand"
	"reflect"
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

func (v *PoT) IsInT(d interface{}) bool {
	if reflect.ValueOf(d).Kind() == reflect.Int {
		return true
	}

	return false
}

func (v *PoT) IsInT64(d interface{}) bool {
	if reflect.ValueOf(d).Kind() == reflect.Int64 {
		return true
	}

	return false
}

func (v *PoT) IsString(d interface{}) bool {
	if reflect.ValueOf(d).Kind() == reflect.String {
		return true
	}

	return false
}

func (v *PoT) ReflectContains(d ... interface{}) bool {
	if len(d) != 2 {
		return false
	}

	val := reflect.ValueOf(d[1])
	switch reflect.TypeOf(d[1]).Kind() {
	case reflect.Slice, reflect.Array:
		var i int
		for i = 0; i < val.Len(); i++ {
			if val.Index(i).Interface() == d[0] {
				return true
			}
		}

	case reflect.Map:
		if val.MapIndex(reflect.ValueOf(d[0])).IsValid() {
			return true
		}
	}

	return false
}

func (v *PoT) SetTimeLocation(d string) (*time.Location, error) {
	if d == "" {
		d = data.TimeLocation
	}

	return time.LoadLocation(d)
}

func (v *PoT) Contains(k interface{}, d ...interface{}) bool {
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

func (v *PoT) NumRandom(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

func (v *PoT) Fprintf(writer io.Writer, format string, d ... interface{}) {
	fmt.Fprintf(writer, format, d ...)
}

func (v *PoT) Sprintf(format string, d ... interface{}) string {
	return fmt.Sprintf(format, d ...)
}
