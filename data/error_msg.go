// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"
)

const (
	// Err Message Keywords
	ErrPfx 		string = PoT

	// Success
	PoTStatusOK			int = http.StatusOK

	// Err Status
	PoTUnKnown			int = -1
	PoTStatusNotFound	int = http.StatusNotFound
)

var (
	// Modules: Errors
	eMsg *ErrH = &ErrH {
		ErrPfx: {
			"ErrDefault":		"Unknown Error",

			"PoTModeErr":		"Error PoT.Mode Parameters (0->debug|1->release) in the file of .("+EnvDEV+"|"+EnvSTG+"|"+EnvPRD+").yaml",
			"PoTZapLogErr":		"Error Logs.PoT Parameters",
			"PoTSslCF":			"Missing SSL Cert File",
			"PoTSslKF":			"Missing SSL Key File",
			"PoTJwTErr":		"Missing JwT Parameters",

			"ModParamsErr":		"Error Model Parameter",
			"ModDBTable":		"Error db Table",
			"ModDBSelectErr":	"Error GeT TYPE(ONE|ALL)",

			"MWareNoPriority":	"No Priority",
			"MWareUnknown":		"UnKnown",

			"RedParamsErr":		"Missing Redis Configure",
			"RedEngineErr":		"Missing Redis Engine",

			"AdPropBind":		"Error Property Bind Env",
			"AdPropVar":		"Empty Property Variables",
			"AdPoTPowerErr":	"Missing PoT.Power Configure",

			"PropEnvEmpty":		"Missing Env Parameters(--env="+EnvDEV+"|"+EnvSTG+"|"+EnvPRD+")",
			"PropEnvExclude":	"Exclude Env Parameters",
			"PropEnvFile":		"Missing Env File",

			"AdapterSrc":		"Missing db Configure (Adapter.*)",
			"AdapterMadeMode":	"Error Engine Mode",
			"AdapterMadeDN":	"Error Engine Driver Name",
			"AdapterMadeName":	"Error Engine db Name",
			"AdapterEngineErr":	"Engine Connect Error",

			"CpCacheClient":	"Missing Client",
			"CpCacheSeTParams":	"Error SeT Parameters",
			"CpCacheSeTExp":	"Error SeT expiration",

			"TokenParamsErr":	"Missing Token Parameters",
			"TokenTypeErr":		"Error Token Type(MD5|SHA1|SHA256)",

			"TokenExpired": 	"Token is expired",
			"TokenNotValidYet": "Token hasn't active yet",
			"TokenMalformed": 	"That's not even a token",
			"TokenInvalid": 	"Token Invalid",
		},
	}
)

func SeTErrMsg(msg *ErrH) {
	if msg != nil {
		for k, v := range *msg {
			if k != ErrPfx {
				(*eMsg)[k] = v
			}
		}
	}
}

func GeTErrMsg(pfx, k string, content ...interface{}) string {
	str := cast.ToString((*eMsg)[ErrPfx]["ErrDefault"])

	s, ok := (*eMsg)[pfx][k]
	if ok {
		str = cast.ToString(s) + ", " + fmt.Sprint(content ...)
	}

	return str
}
