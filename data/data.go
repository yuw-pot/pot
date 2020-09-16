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
)

var (
	PoTMode []string = []string {
		gin.DebugMode,
		gin.ReleaseMode,
	}

	/**
	 * Modules: Errors
	 */
	ErrMsg *ErrH = &ErrH {
		ErrPfx: {
			"ErrDefault":		"Unknown Error",

			"AdPropBind":		"Error Property Bind Env",
			"AdPropVar":		"Empty Property Variables",

			"PropEnvEmpty":		"Missing Env Parameters(--env="+EnvDEV+"|"+EnvSTG+"|"+EnvPRD+")",
			"PropEnvExclude":	"Exclude Env Parameters",
			"PropEnvFile":		"Missing Env File",
		},
	}

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
)
