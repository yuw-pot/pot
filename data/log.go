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
