package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
}

type appLogger struct {
	level  string
	logger *zap.SugaredLogger
}

func NewAppLogger(level string) Logger {
	logger := getLogger(level)

	return &appLogger{
		level:  level,
		logger: logger,
	}
}

func (a *appLogger) Info(args ...interface{}) {
	a.logger.Info(args...)
}
func (a *appLogger) Error(args ...interface{}) {
	a.logger.Error(args...)
}
func (a *appLogger) Debug(args ...interface{}) {
	a.logger.Debug(args...)
}
func (a *appLogger) Warn(args ...interface{}) {
	a.logger.Warn(args...)
}
func WithComponent(level string, component string) Logger {
	if component == "" {
		return NewAppLogger(level)
	}

	logger := getLogger(level)

	return &appLogger{
		level:  level,
		logger: logger.Named(component),
	}
}

func getLogger(level string) *zap.SugaredLogger {
	var logger *zap.Logger
	var sugarLogger *zap.SugaredLogger

	switch level {
	case "production":
		logger, _ = zap.NewProduction()
		sugarLogger = logger.Sugar()
	default:
		logger, _ = zap.NewDevelopment()
		sugarLogger = logger.Sugar()
	}

	return sugarLogger
}
