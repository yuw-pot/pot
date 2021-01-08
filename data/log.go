// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package data

const (
	// Zap Log Constant
	//   - LogFormat
	//   - LogFormatConsole
	//   - EncoderConfigure
	LogFormatJson 		string = "json"
	LogFormatConsole 	string = "console"

	LogFileName			string = "./logs/pot/request.log"
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
)

type (
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
