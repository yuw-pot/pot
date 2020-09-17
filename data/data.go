// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

import "github.com/gin-gonic/gin"

const (
	/**
	 * Err Message Keywords
	 */
	ErrPfx string = "PoT"

	/**
	 * Property File Suffix
	 */
	EnvDEV string = "dev"
	EnvSTG string = "stg"
	EnvPRD string = "prd"
	PropertySfx string = "yaml"

	/**
	 * Route Separated
	 */
	RSeP string = "->"

	/**
	 * Zap Log Constant
	 * - LogFormat
	 * - LogFormatConsole
	 * - EncoderConfigure
	 */
	LogFormatJson string = "json"
	LogFormatConsole string = "console"

	LogTime string = "time"
	LogLevel string = "level"
	LogName string = "logger"
	LogCaller string = "caller"
	LogMessage string = "msg"
	LogStackTrace string = "stacktrace"

	// - Log Size -> 1M
	LogMaxSize int = 1
	// - Keep Days
	LogMaxBackups int = 5
	// - Keep Files Counts
	LogMaxAge int = 7
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

			"AdPropBind":		"Error Property Bind Env",
			"AdPropVar":		"Empty Property Variables",

			"PropEnvEmpty":		"Missing Env Parameters(--env="+EnvDEV+"|"+EnvSTG+"|"+EnvPRD+")",
			"PropEnvExclude":	"Exclude Env Parameters",
			"PropEnvFile":		"Missing Env File",

			"AdapterSrc":		"Missing db Configure (Adapter.*)",
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
		Adapter int
		Redis int
	}

	/**
	 * Modules: Logs
	 */
	ZLogPoT struct {
		FileName string `yaml:"filename"`
		ZapCoreLevel string `yaml:"zap_core_level"`
		Format string `yaml:"format"`
		Time string `yaml:"time"`
		Level string `yaml:"level"`
		Name string `yaml:"name"`
		Caller string `yaml:"caller"`
		Message string `yaml:"message"`
		StackTrace string `yaml:"stacktrace"`
		MaxSize int `yaml:"max_size"`
		MaxBackups int `yaml:"max_backups"`
		MaxAge int `yaml:"max_age"`
	}
)
