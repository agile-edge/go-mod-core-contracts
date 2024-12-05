/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

/*
Package logger provides a client for integration with the support-logging service. The client can also be configured
to write logs to a local file rather than sending them to a service.
*/
package logger

// Logging client for the Go implementation of agile-edge

import (
	"fmt"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/agile-edge/go-mod-core-contracts/v3/errors"
)

// LoggingClient defines the interface for logging operations.
type LoggingClient interface {
	// SetLogLevel sets minimum severity log level. If a logging method is called with a lower level of severity than
	// what is set, it will result in no output.
	SetLogLevel(logLevel string) errors.EdgeX
	// LogLevel returns the current log level setting
	LogLevel() string
	// Debug logs a message at the DEBUG severity level
	Debug(msg string, args ...interface{})
	// Error logs a message at the ERROR severity level
	Error(msg string, args ...interface{})
	// Info logs a message at the INFO severity level
	Info(msg string, args ...interface{})
	// Warn logs a message at the WARN severity level
	Warn(msg string, args ...interface{})
	// Debugf logs a formatted message at the DEBUG severity level
	Debugf(msg string, args ...interface{})
	// Errorf logs a formatted message at the ERROR severity level
	Errorf(msg string, args ...interface{})
	// Infof logs a formatted message at the INFO severity level
	Infof(msg string, args ...interface{})
	// Warnf logs a formatted message at the WARN severity level
	Warnf(msg string, args ...interface{})
}

func parseLevel(logLevel string) zapcore.Level {
	level, err := zapcore.ParseLevel(logLevel)
	if err != nil {
		level = zap.InfoLevel
	}
	return level
}

// NewClient creates an instance of LoggingClient
func NewClient(owningServiceName string, logLevel string) LoggingClient {
	level := zap.NewAtomicLevelAt(parseLevel(logLevel))
	//配置zap日志库的编码器
	encoder := zapcore.EncoderConfig{
		TimeKey:          "timestamp",
		LevelKey:         "level",
		NameKey:          "logger",
		CallerKey:        "caller",
		MessageKey:       "msg",
		StacktraceKey:    "stack",
		EncodeTime:       zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
		LineEnding:       zapcore.DefaultLineEnding,
		EncodeLevel:      CapitalLevelEncoder,
		EncodeDuration:   zapcore.SecondsDurationEncoder,
		EncodeCaller:     ShortCallerEncoder,
		ConsoleSeparator: " ",
	}

	logFilename := "app"
	if owningServiceName != "" {
		logFilename = strings.ReplaceAll(owningServiceName, "/", "-")
	}
	//日志切割
	writeSyncer := getLogWriter(logFilename)
	//设置日志级别
	var core zapcore.Core
	writers := []zapcore.WriteSyncer{zapcore.AddSync(writeSyncer)}
	if getEnv(LogStdout, "true") != "false" {
		writers = append(writers, zapcore.AddSync(os.Stdout))
	}
	core = zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoder),
		zapcore.NewMultiWriteSyncer(writers...),
		level,
	)
	logger := zap.New(core,
		zap.AddStacktrace(zap.NewAtomicLevelAt(zapcore.ErrorLevel)),
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.Development(),
	)
	return &zapLogger{log: logger, Sync: logger.Sync, level: level}
}

func (lc zapLogger) SetLogLevel(logLevel string) errors.EdgeX {
	if level, err := zapcore.ParseLevel(logLevel); err == nil {
		lc.level.SetLevel(level)
		return nil
	}

	return errors.NewCommonEdgeX(errors.KindContractInvalid, fmt.Sprintf("invalid log level `%s`", logLevel), nil)
}

func (lc zapLogger) LogLevel() string {
	return lc.level.String()
}

func (lc zapLogger) Info(msg string, args ...interface{}) {
	lc.log.Sugar().Infof(msg, args...)
}

func (lc zapLogger) Debug(msg string, args ...interface{}) {
	lc.log.Sugar().Debugf(msg, args...)
}

func (lc zapLogger) Warn(msg string, args ...interface{}) {
	lc.log.Sugar().Warnf(msg, args...)
}

func (lc zapLogger) Error(msg string, args ...interface{}) {
	lc.log.Sugar().Errorf(msg, args...)
}

func (lc zapLogger) Infof(msg string, args ...interface{}) {
	lc.log.Sugar().Infof(msg, args...)
}

func (lc zapLogger) Debugf(msg string, args ...interface{}) {
	lc.log.Sugar().Debugf(msg, args...)
}

func (lc zapLogger) Warnf(msg string, args ...interface{}) {
	lc.log.Sugar().Warnf(msg, args...)
}

func (lc zapLogger) Errorf(msg string, args ...interface{}) {
	lc.log.Sugar().Errorf(msg, args...)
}
