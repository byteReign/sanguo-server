package logger

import (
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Config 日志初始化配置
type Config struct {
	Level      string `yaml:"level"`  // debug / info / warn / error
	Format     string `yaml:"format"` // json / text
	OutputPath string `yaml:"output"` // stdout / stderr / 文件路径
}

// Init 初始化全局日志实例，应在服务启动时调用
func Init(cfg Config) error {
	level, err := logrus.ParseLevel(strings.ToLower(cfg.Level))
	if err != nil {
		return fmt.Errorf("parse log level: %w", err)
	}
	log.SetLevel(level)

	switch strings.ToLower(cfg.Format) {
	case "", "text":
		log.SetFormatter(&logrus.TextFormatter{
			FullTimestamp:   true,
			TimestampFormat: "2006-01-02 15:04:05",
		})
	case "json":
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	default:
		return fmt.Errorf("unsupported log format: %s", cfg.Format)
	}

	output, err := resolveOutput(cfg.OutputPath)
	if err != nil {
		return err
	}
	log.SetOutput(output)

	return nil
}

func resolveOutput(path string) (io.Writer, error) {
	switch strings.ToLower(path) {
	case "", "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	default:
		file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
		if err != nil {
			return nil, fmt.Errorf("open log file: %w", err)
		}
		return file, nil
	}
}

// Get 返回底层 logrus 实例，便于需要自定义字段等高级场景
func Get() *logrus.Logger {
	return log
}

func Debug(args ...interface{}) {
	log.Debug(args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
}

func Info(args ...interface{}) {
	log.Info(args...)
}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
}

func Warn(args ...interface{}) {
	log.Warn(args...)
}

func Warnf(format string, args ...interface{}) {
	log.Warnf(format, args...)
}

func Error(args ...interface{}) {
	log.Error(args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
}

func Fatal(args ...interface{}) {
	log.Fatal(args...)
}

func Fatalf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}

func WithField(key string, value interface{}) *logrus.Entry {
	return log.WithField(key, value)
}

func WithFields(fields logrus.Fields) *logrus.Entry {
	return log.WithFields(fields)
}

func WithError(err error) *logrus.Entry {
	return log.WithError(err)
}
