package logger

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

// Config 日志初始化配置
type Config struct {
	Level    string `yaml:"level"`    // debug / info / warn / error
	Format   string `yaml:"format"`   // json / text
	Output   string `yaml:"output"`   // stdout / stderr / file
	FilePath string `yaml:"filePath"` // output=file 时生效，如 logs/app.log
	MaxAge   int    `yaml:"maxAge"`   // 日志保留天数，默认 7
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

	output, err := resolveOutput(cfg)
	if err != nil {
		return err
	}
	log.SetOutput(output)

	return nil
}

func resolveOutput(cfg Config) (io.Writer, error) {
	switch strings.ToLower(cfg.Output) {
	case "", "stdout":
		return os.Stdout, nil
	case "stderr":
		return os.Stderr, nil
	case "file":
		return newRotateWriter(cfg)
	default:
		// 兼容旧配置：output 直接写文件路径
		cfg.FilePath = cfg.Output
		cfg.Output = "file"
		return newRotateWriter(cfg)
	}
}

func newRotateWriter(cfg Config) (io.Writer, error) {
	path := cfg.FilePath
	if path == "" {
		path = "logs/app.log"
	}

	maxAge := cfg.MaxAge
	if maxAge <= 0 {
		maxAge = 7
	}

	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return nil, fmt.Errorf("create log dir: %w", err)
	}

	ext := filepath.Ext(path)
	name := strings.TrimSuffix(filepath.Base(path), ext)
	pattern := filepath.Join(dir, name+".%Y-%m-%d"+ext)

	opts := []rotatelogs.Option{
		rotatelogs.WithRotationTime(24 * time.Hour),
		rotatelogs.WithMaxAge(time.Duration(maxAge) * 24 * time.Hour),
	}
	// Windows 创建符号链接需要管理员权限，跳过 WithLinkName，直接写按日期分割的文件
	if runtime.GOOS != "windows" {
		opts = append(opts, rotatelogs.WithLinkName(path))
	}

	writer, err := rotatelogs.New(pattern, opts...)
	if err != nil {
		return nil, fmt.Errorf("init rotate log: %w", err)
	}

	return writer, nil
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
