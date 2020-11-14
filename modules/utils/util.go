// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"io"
	"math/rand"
	"reflect"
	"runtime"
	"strings"
	"time"
)

const char = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type (
	PoT struct {

	}
)

func New() *PoT {
	return &PoT {

	}
}

func (v *PoT) RandString(size int) string {
	rand.NewSource(time.Now().UnixNano())

	var res bytes.Buffer
	for i := 0; i < size; i ++ {
		res.WriteByte(char[rand.Int63() % int64(len(char))])
	}

	return res.String()
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

	for _, val := range d {
		if strings.Contains(cast.ToString(k), cast.ToString(val)) {
			return true
		}
	}

	return false
}

func (v *PoT) ToJson(d interface{}) (interface{}, error) {
	res, err := json.Marshal(d)
	if err != nil { return nil, err }

	return string(res), nil
}

func (v *PoT) ToStruct(d interface{}, res interface{}) error {
	err := json.Unmarshal([]byte(cast.ToString(d)), res)
	if err != nil { return err }

	return nil
}

func (v *PoT) MergeH(d ... *data.H) *data.H {
	h := &data.H{}
	for _, val := range d {
		if val != nil {
			for k, data := range *val {
				(*h)[k] = data
			}
		}
	}

	return h
}

func (v *PoT) GeTUintPtrFuncPC() (interface{}, interface{}, interface{}) {
	funcPCUintPtr := make([]uintptr, 1)
	runtime.Callers(2, funcPCUintPtr)
	funcPCName := runtime.FuncForPC(funcPCUintPtr[0]).Name()
	ptrFuncPCName := strings.Split(funcPCName, "/")

	srvFuncPCName := ptrFuncPCName[len(ptrFuncPCName)-2]
	ctrFuncPCName := strings.Split(ptrFuncPCName[len(ptrFuncPCName)-1], ".")

	return srvFuncPCName, ctrFuncPCName[len(ctrFuncPCName)-2], ctrFuncPCName[len(ctrFuncPCName)-1]
}

func (v *PoT) NumRandom(max int) int {
	return rand.New(rand.NewSource(time.Now().Unix())).Intn(max)
}

func (v *PoT) Fprintf(writer io.Writer, format string, d ... interface{}) {
	fmt.Fprintf(writer, format, d ...)
}

func (v *PoT) Sprintf(format string, d ... interface{}) string {
	return fmt.Sprintf(format, d ...)
}
