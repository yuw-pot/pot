// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	// Err Message Keywords
	ErrPfx 		string = PoT
)

var (
	/**
	 * Modules: Errors
	 */
	ErrMsg *ErrH = &ErrH{
		ErrPfx: {
			"ErrDefault":		"Unknown Error",

			"PoTModeErr":		"Error PoT.Mode Parameters (0->debug|1->release) in the file of .("+EnvDEV+"|"+EnvSTG+"|"+EnvPRD+").yaml",
			"PoTZapLogErr":		"Error Logs.PoT Parameters",
			"PoTSslCF":			"Missing SSL Cert File",
			"PoTSslKF":			"Missing SSL Key File",

			"ModParamsErr":		"Error Model Parameter",
			"ModDBTable":		"Error db Table",
			"ModDBSelectErr":	"Error GeT TYPE(ONE|ALL)",

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
		},
	}
)
