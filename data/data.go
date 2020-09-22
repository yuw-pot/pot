// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

import "github.com/gin-gonic/gin"

const (
	// Time Location
	TimeLocation	string = "Asia/Shanghai"

	// Err Message Keywords
	ErrPfx 		string = "PoT"

	// Property File Suffix
	EnvDEV 		string = "dev"
	EnvSTG 		string = "stg"
	EnvPRD 		string = "prd"
	PropertySfx string = "yaml"

	// Route Separated
	RSeP 		string = "->"

	// Zap Log Constant
	//   - LogFormat
	//   - LogFormatConsole
	//   - EncoderConfigure
	LogFormatJson 		string = "json"
	LogFormatConsole 	string = "console"

	LogTime 			string = "time"
	LogLevel 			string = "level"
	LogName 			string = "logger"
	LogCaller 			string = "caller"
	LogMessage 			string = "msg"
	LogStackTrace 		string = "stacktrace"

	// - Log Size -> 1M
	LogMaxSize 			int = 1
	// - Keep Days
	LogMaxBackups 		int = 5
	// - Keep Files Counts
	LogMaxAge 			int = 7

	// Models
	ModONE		string = "ONE"
	ModALL		string = "ALL"
	ModByAsc	string = "ASC"
	ModByEsc	string = "DESC"
)

var (
	PoTMode []string = []string{
		gin.DebugMode,
		gin.ReleaseMode,
	}

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
		},
	}

	/**
	 * Modules: Property
	 * - dev
	 * - stg
	 * - prd
	 */
	PropertySfxs []interface{} = []interface{}{
		EnvDEV,
		EnvSTG,
		EnvPRD,
	}
)

type (
	/**
	 * Map Struct
	 */

	// Default Type
	H map[string]interface{}

	// Err Type
	ErrH map[string]H

	// Power
	PowerPoT struct {
		Adapter 		int	`yaml:"adapter" json:"adapter"`
		Redis 			int `yaml:"redis" json:"redis"`
	}

	SrvPoT struct {
		Status int
		Msg interface{}
		Response *H
	}

	ModPoT struct {
		Types string `json:"types"`// Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
		Table string `json:"table"`
		Field string `json:"field"`
		Joins [][]interface{} `json:"joins"`
		Limit int `json:"limit"`
		Start []int `json:"start"`
		Query interface{} `json:"query"`
		QueryArgs []interface{} `json:"query_args"`
		Columns []string `json:"columns"`
		OrderType string `json:"order_type"`
		OrderArgs []string `json:"order_args"`
	}

	/**
	 * Modules: Logs
	 */
	ZLogPoT struct {
		FileName 		string `yaml:"filename" json:"filename"`
		ZapCoreLevel 	string `yaml:"zap_core_level" json:"zap_core_level"`
		Format 			string `yaml:"format" json:"format"`
		Time 			string `yaml:"time" json:"time"`
		Level 			string `yaml:"level" json:"level"`
		Name 			string `yaml:"name" json:"name"`
		Caller 			string `yaml:"caller" json:"caller"`
		Message 		string `yaml:"message" json:"message"`
		StackTrace 		string `yaml:"stacktrace" json:"stack_trace"`

		MaxSize 		int `yaml:"max_size" json:"max_size"`
		MaxBackups 		int `yaml:"max_backups" json:"max_backups"`
		MaxAge 			int `yaml:"max_age" json:"max_age"`
	}
)
