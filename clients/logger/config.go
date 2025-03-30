package logger

import (
	"os"
	"strconv"
	"strings"
)

type LogConfig struct {
	FilePath      string
	MaxSize       int
	MaxBackups    int
	MaxAge        int
	Compress      bool
	ConsoleOutput bool
}

func loadLogConfig() LogConfig {
	return LogConfig{
		FilePath:      getEnv("LOG_FILE_PATH", "/var/log/agile-edge/app.log"),
		MaxSize:       getEnvAsInt("LOG_MAX_SIZE", 100),
		MaxBackups:    getEnvAsInt("LOG_MAX_BACKUPS", 3),
		MaxAge:        getEnvAsInt("LOG_MAX_AGE", 30),
		Compress:      getEnvAsBool("LOG_COMPRESS", true),
		ConsoleOutput: getEnvAsBool("LOG_CONSOLE", false),
	}
}

// 辅助函数：读取环境变量并转换类型
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}

func getEnvAsInt(key string, defaultValue int) int {
	if value, exists := os.LookupEnv(key); exists {
		if intValue, err := strconv.Atoi(value); err == nil {
			return intValue
		}
	}
	return defaultValue
}

func getEnvAsBool(key string, defaultValue bool) bool {
	if value, exists := os.LookupEnv(key); exists {
		return strings.ToLower(value) == "true"
	}
	return defaultValue
}
