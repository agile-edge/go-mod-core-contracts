package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"strconv"
)

var _ LoggingClient = (*zapLogger)(nil)

const (
	LogPath     = "LOG_PATH"
	LogFileSize = "LOG_FILE_SIZE"
	LogFileNum  = "LOG_FILE_NUM"
	LogFileAge  = "LOG_FILE_AGE"
	LogStdout   = "LOG_STDOUT"
)

type zapLogger struct {
	log   *zap.Logger
	level zap.AtomicLevel
	Sync  func() error
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

// 日志自动切割，采用 lumberjack 实现的
func getLogWriter(filename string) zapcore.WriteSyncer {
	maxSize, err := strconv.Atoi(getEnv(LogFileSize, "1024"))
	if err != nil {
		panic(err)
	}
	maxNum, err := strconv.Atoi(getEnv(LogFileNum, "3"))
	if err != nil {
		panic(err)
	}
	maxAge, err := strconv.Atoi(getEnv(LogFileAge, "30"))
	if err != nil {
		panic(err)
	}
	filename = fmt.Sprintf("%s/%s.log", getEnv(LogPath, "/var/log/agile-edge"), filename)
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename, //指定日志存储位置
		MaxSize:    maxSize,  //日志的最大大小（M）
		MaxBackups: maxNum,   //日志的最大保存数量
		MaxAge:     maxAge,   //日志文件存储最大天数
		Compress:   false,    //是否执行压缩
	}
	return zapcore.AddSync(lumberJackLogger)
}

// CapitalLevelEncoder serializes a Level to an all-caps string. For example,
// InfoLevel is serialized to "INFO".
func CapitalLevelEncoder(l zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + l.CapitalString() + "]")
}

// ShortCallerEncoder serializes a caller in package/file:line format, trimming
// all but the final directory from the full path.
func ShortCallerEncoder(caller zapcore.EntryCaller, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString("[" + caller.TrimmedPath() + "]")
}
