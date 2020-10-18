// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	PoT	string = "PoT"

	// Mode
	DebugMode	string = "debug"
	ReleaseMode	string = "release"

	// Time Location
	TimeLocation	string = "Asia/Shanghai"

	// Route Separated
	RSeP 		string = "->"

	// Models
	ModONE		string = "ONE"
	ModALL		string = "ALL"
	ModByAsc	string = "ASC"
	ModByEsc	string = "DESC"
)

var (
	PoTMode []interface{} = []interface{}{
		DebugMode,
		ReleaseMode,
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
		Types 		string `json:"types"`// Todo: "ONE"->"FetchOne" "ALL"->"FetchAll"
		Table 		string `json:"table"`
		Field 		string `json:"field"`
		Joins 		[][]interface{} `json:"joins"`
		Limit 		int `json:"limit"`
		Start 		[]int `json:"start"`
		Query 		interface{} `json:"query"`
		QueryArgs 	[]interface{} `json:"query_args"`
		Columns 	[]string `json:"columns"`
		OrderType 	string `json:"order_type"`
		OrderArgs 	[]string `json:"order_args"`
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
