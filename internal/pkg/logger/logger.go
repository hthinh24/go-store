package logger

type Logger interface {
	Info(args ...interface{})
	Error(args ...interface{})
	Debug(args ...interface{})
	Warn(args ...interface{})
}

type appLogger struct {
	Logger *Logger
}

func NewAppLogger(logger *Logger) *appLogger {
	return &appLogger{
		Logger: logger,
	}
}
