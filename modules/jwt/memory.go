// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/yuw-pot/pot/data"
	E "github.com/yuw-pot/pot/modules/err"
	"time"
)

type (
	JMemory struct {
		KeY 	string
		Expire 	time.Duration
		Mode 	string
	}

	JMemoryPoT struct {
		Info interface{}
		jwt.StandardClaims
	}
)

var JPoT *JMemory

func (j *JMemory) Produce(jp *JMemoryPoT) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}

	jwt.TimeFunc = time.Now
	expire := j.Expire * getTimeUniT(j.Mode)

	jp.StandardClaims.ExpiresAt = time.Now().Add(expire).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jp)
	return token.SignedString([]byte(j.KeY))
}

func (j *JMemory) Parse(sign string) (*JMemoryPoT, error) {
	token, err := j.parse(sign)
	if err != nil {
		return nil, err
	}

	if jp, ok := token.Claims.(*JMemoryPoT); ok && token.Valid {
		return jp, nil
	}

	return nil, E.Err(data.ErrPfx, "TokenInvalid")
}

func (j *JMemory) Refresh(sign string) (string, error) {
	token, err := j.parse(sign)
	if err != nil {
		return "", err
	}

	if jp, ok := token.Claims.(*JMemoryPoT); ok && token.Valid {
		return j.Produce(jp)
	}

	return "", E.Err(data.ErrPfx, "TokenInvalid")
}

func (j *JMemory) parse(sign string) (*jwt.Token, error)  {
	token, err := jwt.ParseWithClaims(
		sign, &JMemoryPoT{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.KeY), nil
		},
	)

	if err != nil {
		if val, ok := err.(*jwt.ValidationError); ok {
			if val.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, E.Err(data.ErrPfx, "TokenMalformed")
			} else if val.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, E.Err(data.ErrPfx, "TokenExpired")
			} else if val.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, E.Err(data.ErrPfx, "TokenNotValidYet")
			} else {
				return nil, E.Err(data.ErrPfx, "TokenInvalid")
			}
		} else {
			return nil, err
		}
	} else {
		return token, nil
	}
}
