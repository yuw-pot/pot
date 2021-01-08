// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

import (
	"fmt"
	"github.com/spf13/cast"
)

const (
	// Err Message Keywords
	ErrPfx string = PoT
	PoTSuccessOK string = "SuccessOK"
	PoTErrorNull string = "ErrDefault"
)

var (
	// Modules: Errors
	eMsg *ErrH = &ErrH {
		ErrPfx: {
			PoTSuccessOK:		"Success",
			PoTErrorNull:		"Unknown Error",

			"PoTModeErr":		"Error PoT.Mode Parameters (0->debug|1->release) in the file of .("+EnvDEV+"|"+EnvSTG+"|"+EnvPRD+").yaml",
			"PoTZapLogErr":		"Error Logs.PoT Parameters",
			"PoTSslCF":			"Missing SSL Cert File",
			"PoTSslKF":			"Missing SSL Key File",
			"PoTJwTErr":		"Missing JwT Parameters",
			"PoTRouteErr":		"Missing Route Configuration",

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

			"TokenParamsErr":	"Missing Token Parameters",
			"TokenTypeErr":		"Error Token Type(MD5|SHA1|SHA256)",

			"TokenExpired": 	"Token is expired",
			"TokenNotValidYet": "Token hasn't active yet",
			"TokenMalformed": 	"That's not even a token",
			"TokenInvalid": 	"Token Invalid",

			"CryptoParamsErr":	"Params Empty",
			"CryptoAeSKeYErr":	"AES KeY Type Error",
			"CryptoRsAPubKeY":	"Missing RsA Public KeY",
			"CryptoRsAPriKeY":	"Missing RsA Private KeY",

			"JwTKeYErr":		"Err Key",
			"JwTKeYEmpty":		"Empty Key",
			"JwTCacheErr":		"Err Cache Client",
			"JwTExpireErr":		"JwT Expire is 0",
			"JwTInfoType":		"JwT Info is not struct",
			"JwTInfoEmpty":		"JwT Info is empty",
			"JwTCacheResult":	"JwT Token & Encrypt is empty",

			"FilePathExist":	"File Path Exist",

			"CasbinEnforcer":	"Casbin Initialize Error",

			"ComponentCacheClient":			"Missing Client",
			"ComponentCacheSeTParams":		"Error SeT Parameters",
			"ComponentCasbinParamMod":		"Error Param Casbin Mod",
			"ComponentCasbinParamAdapter":	"Error Param Casbin Adapter",

			"ComponentRequestMethod":	"Error Request Method",
			"ComponentRequestURL":		"Empty Request URL",

			"ComponentSubscribeKeY": 	"Configure Subscribe KeY Empty",

			"AdapterSrc":				"Missing db Configure (Adapter.*)",
			"AdapterConfigErr":			"Error db Configure",
			"AdapterMode":				"Error Engine Mode",
			"AdapterModeDN":			"Error Engine Driver Name",
			"AdapterModeName":			"Error Engine db Tag",
			"AdapterEngineErr":			"Engine Connect Error",
			"AdapterEngineGroupErr":	"Engine Group Connect Error",
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
		str = cast.ToString(s)
		if len(content) != 0 {
			str = str + ", " + fmt.Sprint(content ...)
		}
	}

	return str
}
