// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	PoT	string = "PoT"

	// PoT Mode
	ReleaseMode	string = "release"
	ConsoleMode string = "console"

	// Time Location
	TimeLocation	string = "Asia/Shanghai"
)

var (
	PoTMode []interface{} = []interface{}{
		ConsoleMode,
		ReleaseMode,
	}
)

type (
	// Default Type
	H map[string]interface{}

	// Err Type
	ErrH map[string]H

	// Power
	PowerPoT struct {
		JwT 	int `yaml:"jwt" json:"jwt"`
		Adapter int	`yaml:"adapter" json:"adapter"`
		Redis 	int `yaml:"redis" json:"redis"`
	}

	TpL struct {
		Status int
		Msg interface{}
		Response *H
	}

	SrvTpL struct {
		Status int
		Msg interface{}
		Data *H
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

func SrvTpLInitialized() *SrvTpL {
	return &SrvTpL {
		Status: PoTStatusOK,
		Msg:    nil,
		Data:   nil,
	}
}

func TpLInitialized() *TpL {
	return &TpL {
		Status: PoTStatusOK,
		Msg:    nil,
		Response: &H{},
	}
}
