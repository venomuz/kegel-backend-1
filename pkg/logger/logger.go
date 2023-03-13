package logger

import (
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Field ...
type Field = zapcore.Field

var (
	// Int ..
	Int = zap.Int
	// String ...
	String = zap.String
	// Error ...
	Error = zap.Error
	// Bool ...
	Bool = zap.Bool

	// Any ...
	Any = zap.Any
)

// Logger ...
type Logger interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type Impl struct {
	zap *zap.Logger
}

var customTimeFormat string

// New ...
func New(level, namespace string) *Impl {
	if level == "" {
		level = LevelInfo
	}

	logger := Impl{
		zap: newZapLogger(level, time.RFC3339),
	}

	logger.zap = logger.zap.Named(namespace)

	zap.RedirectStdLog(logger.zap)

	return &logger
}

func (l *Impl) Debug(msg string, fields ...Field) {
	l.zap.Debug(msg, fields...)
}

func (l *Impl) Info(msg string, fields ...Field) {
	l.zap.Info(msg, fields...)
}

func (l *Impl) Warn(msg string, fields ...Field) {
	l.zap.Warn(msg, fields...)
}

func (l *Impl) Error(msg string, fields ...Field) {
	l.zap.Error(msg, fields...)
}

func (l *Impl) Fatal(msg string, fields ...Field) {
	l.zap.Fatal(msg, fields...)
}

// GetNamed ...
func GetNamed(l Logger, name string) Logger {
	switch v := l.(type) {
	case *Impl:
		v.zap = v.zap.Named(name)
		return v
	default:
		l.Info("logger.GetNamed: invalid logger type")
		return l
	}
}

// WithFields ...
func WithFields(l Logger, fields ...Field) Logger {
	switch v := l.(type) {
	case *Impl:
		return &Impl{
			zap: v.zap.With(fields...),
		}
	default:
		l.Info("logger.WithFields: invalid logger type")
		return l
	}
}

// Cleanup ...
func Cleanup(l Logger) error {
	switch v := l.(type) {
	case *Impl:
		return v.zap.Sync()
	default:
		l.Info("logger.Cleanup: invalid logger type")
		return nil
	}
}
