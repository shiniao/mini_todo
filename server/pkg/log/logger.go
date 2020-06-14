package log

import (
	"errors"
)

// log：全局 log 变量
var log Logger

// Fields Type to pass when we want to call WithFields for structured logging
type Fields map[string]interface{}

const (
	// InstanceZapLogger zap logger
	InstanceZapLogger int = iota
	// here add other logger
)

var (
	errInvalidLoggerInstance = errors.New("invalid logger instance")
)

// Logger 接口
type Logger interface {
	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})
	WithFields(keyValues Fields) Logger
}

// Config 代表 log 的配置
type Config struct {
	Writers         string `yaml:"writers"` // 日志写入类型
	LoggerLevel     string `yaml:"logger_level"` // Logger 等级
	LoggerFile      string `yaml:"logger_file"` // Logger写入文件位置
	LoggerWarnFile  string `yaml:"logger_warn_file"`  // Warn 写入文件位置
	LoggerErrorFile string `yaml:"logger_error_file"` // Error写入文件位置
	LogFormatText   bool   `yaml:"log_format_text"`
	RollingPolicy   string `yaml:"rollingPolicy"`
	LogRotateDate   int    `yaml:"log_rotate_date"`
	LogRotateSize   int    `yaml:"log_rotate_size"`
	LogBackupCount  int    `yaml:"log_backup_count"`
}

// NewLogger 返回新的logger实例
func NewLogger(cfg *Config, loggerInstance int) error {
	// 可扩展
	switch loggerInstance {
	// 如果是zap logger
	case InstanceZapLogger:
		logger, err := NewZapLogger(cfg)
		if err != nil {
			return err
		}
		log = logger
		return nil
	default:
		return errInvalidLoggerInstance
	}
}

// Debug logger
func Debug(args ...interface{}) {
	log.Debug(args...)
}

// Info logger
func Info(args ...interface{}) {
	log.Info(args...)
}

// Warn logger
func Warn(args ...interface{}) {
	log.Warn(args...)
}

// Error logger
func Error(args ...interface{}) {
	log.Error(args...)
}

// Fatal logger
func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

// Debugf logger
func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

// Infof logger
func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

// Warnf logger
func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

// Errorf logger
func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

// Fatalf logger
func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

// Panicf logger
func Panicf(format string, args ...interface{}) {
	log.Panicf(format, args...)
}

// WithFields logger
func WithFields(keyValues Fields) Logger {
	return log.WithFields(keyValues)
}