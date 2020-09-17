// Copyright 2022 YuWenYu  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.

package zlog

import (
	"github.com/natefinch/lumberjack"
	"github.com/yuw-pot/pot/data"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

type (
	PoT struct {
		cfg *data.ZLogPoT
		zlc zapcore.Core
	}
)

func New(cfg *data.ZLogPoT) *PoT {
	return (&PoT{
		cfg: cfg,
		zlc: nil,
	}).initialized()
}

func (log *PoT) Made() *zap.Logger {
	return zap.New(log.zlc, zap.AddCaller(), zap.Development())
}

func (log *PoT) ZErr(err error) zap.Field {
	return zap.Error(err)
}

func (log *PoT) initialized() *PoT {
	//   - Check Zap Log PoT
	log.checkZapLogPoT()

	cfgEnCoder := &zapcore.EncoderConfig {
		MessageKey:     log.cfg.Message,
		LevelKey:       log.cfg.Level,
		TimeKey:        log.cfg.Time,
		NameKey:        log.cfg.Name,
		CallerKey:      log.cfg.Caller,
		StacktraceKey:  log.cfg.StackTrace,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	hook := lumberjack.Logger{
		Filename:   log.cfg.FileName,
		MaxSize:    log.cfg.MaxSize,
		MaxBackups: log.cfg.MaxBackups,
		MaxAge:     log.cfg.MaxAge,
		Compress:   true,
		LocalTime:	true,
	}

	log.zlc = zapcore.NewCore(
		log.selectZapEnCoder(cfgEnCoder),
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stderr), zapcore.AddSync(&hook)), // 打印到控制台和文件
		zap.NewAtomicLevelAt(log.selectZapCoreLevel()),
	)

	return log
}

func (log *PoT) checkZapLogPoT() {
	if log.cfg.Message == "" {
		log.cfg.Message = data.LogMessage
	}

	if log.cfg.Level == "" {
		log.cfg.Level = data.LogLevel
	}

	if log.cfg.Format == "" {
		log.cfg.Format = data.LogFormatJson
	}

	if log.cfg.Time == "" {
		log.cfg.Time = data.LogTime
	}

	if log.cfg.Name == "" {
		log.cfg.Name = data.LogName
	}

	if log.cfg.Caller == "" {
		log.cfg.Caller = data.LogCaller
	}

	if log.cfg.StackTrace == "" {
		log.cfg.StackTrace = data.LogStackTrace
	}

	if log.cfg.MaxSize == 0 {
		log.cfg.MaxSize = data.LogMaxSize
	}

	if log.cfg.MaxBackups == 0 {
		log.cfg.MaxBackups = data.LogMaxBackups
	}

	if log.cfg.MaxAge == 0 {
		log.cfg.MaxAge = data.LogMaxAge
	}
}

func (log *PoT) selectZapEnCoder(cfgEnCoder *zapcore.EncoderConfig) zapcore.Encoder {
	switch log.cfg.Format {
	case data.LogFormatJson:
		return zapcore.NewJSONEncoder(*cfgEnCoder)

	case data.LogFormatConsole:
		return zapcore.NewConsoleEncoder(*cfgEnCoder)

	default:
		return zapcore.NewConsoleEncoder(*cfgEnCoder)
	}
}

func (log *PoT) selectZapCoreLevel() zapcore.Level {
	switch strings.ToLower(log.cfg.ZapCoreLevel) {
	case "info":
		return zap.InfoLevel

	case "debug":
		return zap.DebugLevel

	case "warn":
		return zap.WarnLevel

	case "fatal":
		return zap.FatalLevel

	case "panic":
		return zap.PanicLevel

	case "dpanic":
		return zap.DPanicLevel

	default:
		return zap.InfoLevel
	}
}

