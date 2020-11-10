// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	"context"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/cast"
	"github.com/yuw-pot/pot/data"
	"github.com/yuw-pot/pot/modules/crypto"
	E "github.com/yuw-pot/pot/modules/err"
	"github.com/yuw-pot/pot/modules/utils"
	"reflect"
	"strings"
	"time"
)

const (
	accessKeYSeperate 	string = "::"
	accessKeYRandString int = 5

	errCheckClient 		string = "ErrCheckClient"
	errCheckKeY 		string = "ErrCheckKeY"
	errCheckMethod 		string = "ErrCheckMethod"
	errCheckExpire		string = "ErrCheckExpire"
	errCheckInfoEmpty 	string = "ErrCheckInfoEmpty"
	errCheckInfoType 	string = "ErrCheckInfoType"

	RefreshToken 		int = 1
)

type (
	JCachePoT struct {
		// Key
		KeY string
		// Info Struct
		Info interface{}
		// Time UniT (Time.Hour & Time.Minute & Time.Second)
		Method string
		// Time
		Expire time.Duration

		// Components.Client
		client *redis.Client
		// Module.Utils
		v *utils.PoT
	}
)

func NewJCacheAuth(client *redis.Client) *JCachePoT {
	return &JCachePoT {
		client: client,
		v: utils.New(),

		KeY: "",
		Info: nil,
		Method: data.TimeHour,
		Expire: data.JwTExpire,
	}
}

func (j *JCachePoT) Produce() (interface{}, error) {
	if _, err := j.check(
		errCheckClient,
		errCheckKeY,
		errCheckMethod,
		errCheckExpire,
		errCheckInfoEmpty,
		errCheckInfoType,
	); err != nil { return nil, err }

	accessToken, strInfo, err := j.getAccessToken()
	if err != nil {
		return nil, err
	}

	cryptoJwT := crypto.New()
	cryptoJwT.Mode = data.ModeAeS
	cryptoJwT.D = []interface{}{ j.KeY }

	aes, err := cryptoJwT.Made()
	if err != nil { return nil, err }

	aesEncrypt, err := aes.(*crypto.AeSPoT).EncrypT(strInfo)
	if err != nil { return nil, err }

	err = j.setAccessToken(accessToken, aesEncrypt)
	if err != nil { return nil, err }

	return accessToken, nil
}

func (j *JCachePoT) Parse(accessToken interface{}, toStruct interface{}, refresh ... int) (interface{}, error) {
	if _, err := j.check(
		errCheckClient,
		errCheckKeY,
	); err != nil { return nil, err }

	aesEncrypt, err := j.client.Get(
		context.Background(), cast.ToString(accessToken),
	).Result()
	if err != nil { return nil, err }

	cryptoJwT := crypto.New()
	cryptoJwT.Mode = data.ModeAeS
	cryptoJwT.D = []interface{}{ j.KeY }

	aes, _ := cryptoJwT.Made()
	aesDecrypt, err := aes.(*crypto.AeSPoT).DecrypT(aesEncrypt)
	if err != nil { return nil, err }

	err = j.v.ToStruct(aesDecrypt, toStruct)
	if err != nil { return nil, err }

	newAccessToken, err := j.refresh(accessToken, aesEncrypt, refresh ...)
	if err != nil { return nil, err}

	return newAccessToken, nil
}

func (j *JCachePoT) refresh(accessToken, aesEncrypt interface{}, d ... int) (interface{}, error) {
	if _, err := j.check(
		errCheckMethod, errCheckExpire,
	); err != nil { return nil, err }

	if accessToken == nil || aesEncrypt == nil {
		return nil, E.Err(data.ErrPfx, "JwTCacheResult")
	}

	if len(d) != 0 && d[0] == RefreshToken {
		return j.accessTokenEdit(accessToken, aesEncrypt)
	} else {
		return j.accessTokenCopy(accessToken, aesEncrypt)
	}
}

func (j *JCachePoT) accessTokenCopy(accessToken, aesEncrypt interface{}) (interface{}, error) {
	err := j.setAccessToken(accessToken, aesEncrypt)
	if err != nil { return nil, err }

	return accessToken, nil
}

func (j *JCachePoT) accessTokenEdit(accessToken, aesEncrypt interface{}) (interface{}, error) {
	cryptoJwT := crypto.New()
	cryptoJwT.Mode = data.ModeAeS
	cryptoJwT.D = []interface{}{ j.KeY }

	aes, _ := cryptoJwT.Made()
	aesDecrypt, err := aes.(*crypto.AeSPoT).DecrypT(aesEncrypt)
	if err != nil { return nil, err }

	strRand := j.v.RandString(accessKeYRandString)
	accessKeY := strings.Join([]string {
		j.KeY, cast.ToString(aesDecrypt), strRand,
	}, accessKeYSeperate)

	cryptoJwT.Mode = data.ModeToken
	cryptoJwT.D = []interface{}{data.SHA256, accessKeY}

	newAccessToken, err := cryptoJwT.Made()
	if err != nil { return nil, err }

	err = j.setAccessToken(newAccessToken, aesEncrypt)
	if err != nil { return nil, err }

	return newAccessToken, nil
}

func (j *JCachePoT) IsAccessTokenExisT(accessToken interface{}) bool {
	if _, err := j.check( errCheckClient ); err != nil { return false }

	err := j.client.Get(
		context.Background(), cast.ToString(accessToken),
	).Err()
	if err != nil { return false }

	return true
}

func (j *JCachePoT) getAccessToken() (interface{}, interface{}, error) {
	strInfo, err := j.v.ToJson(j.Info)
	if err != nil { return nil, nil, err }

	strRand := j.v.RandString(accessKeYRandString)
	accessKeY := strings.Join([]string {
		j.KeY, cast.ToString(strInfo), strRand,
	}, accessKeYSeperate)

	cryptoJwT := crypto.New()
	cryptoJwT.Mode = data.ModeToken
	cryptoJwT.D = []interface{}{data.SHA256, accessKeY}

	accessToken, err := cryptoJwT.Made()
	if err != nil { return nil, nil, err }

	return accessToken, strInfo, nil
}

func (j *JCachePoT) setAccessToken(accessToken, aesEncrypt interface{}) error {
	expire := j.Expire * getTimeUniT(j.Method)
	_, err := j.client.Set(
		context.Background(), cast.ToString(accessToken),
		aesEncrypt, expire,
	).Result()
	if err != nil { return err }

	return nil
}

func (j *JCachePoT) check(d ... interface{}) (bool, error) {
	if ok := j.v.Contains(errCheckClient, d ...); ok {
		if j.client == nil {
			return false, E.Err(data.ErrPfx, "JwTCacheErr")
		}
	}

	if ok := j.v.Contains(errCheckKeY, d ...); ok {
		if j.KeY == "" {
			return false, E.Err(data.ErrPfx, "JwTKeYEmpty")
		}
	}

	if ok := j.v.Contains(errCheckMethod, d ...); ok {
		if ok := j.v.Contains(j.Method, data.TimeHour, data.TimeMinute, data.TimeSecond); ok == false {
			return false, E.Err(data.ErrPfx, "JwTKeYErr")
		}
	}

	if ok := j.v.Contains(errCheckExpire, d ...); ok {
		if j.Expire <= 0 {
			return false, E.Err(data.ErrPfx, "JwTExpireErr")
		}
	}

	if ok := j.v.Contains(errCheckInfoEmpty, d ...); ok {
		if j.Info == nil {
			return false, E.Err(data.ErrPfx, "JwTInfoEmpty")
		}
	}

	if ok := j.v.Contains(errCheckInfoType, d ...); ok {
		if reflect.TypeOf(j.Info).Kind() != reflect.Ptr {
			return false, E.Err(data.ErrPfx, "JwTInfoType")
		}
	}

	return true, nil
}