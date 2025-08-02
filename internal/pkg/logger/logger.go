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
	logger *zap.Logger
}

func NewAppLogger(level string) Logger {
	var logger *zap.Logger
	switch level {
	case "production":
		logger, _ = zap.NewProduction()
	default:
		logger, _ = zap.NewDevelopment()
	}

	return &appLogger{
		level:  level,
		logger: logger,
	}
}

func (a *appLogger) Info(args ...interface{}) {
	a.logger.Sugar().Info(args...)
}
func (a *appLogger) Error(args ...interface{}) {
	a.logger.Sugar().Error(args...)
}
func (a *appLogger) Debug(args ...interface{}) {
	a.logger.Sugar().Debug(args...)
}
func (a *appLogger) Warn(args ...interface{}) {
	a.logger.Sugar().Warn(args...)
}
